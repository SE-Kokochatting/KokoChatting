package internal

import (
	"KokoChatting/global"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Conn struct {
	*websocket.Conn
	uid               uint64
	expiredTime       time.Time
	mu                *sync.Mutex
	heartBeatDuration time.Duration
	msgChan           chan Message
}

// serve handle upstream message
func (conn *Conn) serve() {
	var msgType int
	var msg []byte
	var err error
	for {
		msgType, msg, err = conn.ReadMessage()
		if err != nil {
			global.Logger.Error("ws conn read message error", zap.Error(err))
			conn.Close()
			return
		}
		err = getHandler(msgType).handle(msg, conn)
		if err != nil {
			global.Logger.Error("msg handler handle msg error", zap.Error(err))
			conn.Close()
			return
		}
	}
}

// pushMsg send downstream message
func (conn *Conn) pushMsg() {
	for {
		select {
		case msg := <-conn.msgChan:
			conn.WriteMessage(msg.MessageType(), msg.Bytes())
			msgPool.Put(msg)
		}
	}
}


func (conn *Conn) Set(c *websocket.Conn,uid uint64,heartBeatDuration time.Duration){
	conn.Conn = c
	conn.uid = uid
	conn.expiredTime = time.Now().Add(heartBeatDuration)
	conn.mu = new(sync.Mutex)
	conn.msgChan = make(chan Message)
}

func NewConn(conn *websocket.Conn,uid uint64,heartBeatDuration time.Duration) *Conn {
	c := connPool.Get().(*Conn)
	c.Set(conn,uid,heartBeatDuration)
	return c
}

type WsConnManager struct {
	conns                     *sync.Map // map[uint64]*Conn
	maxConnNums               int
	curConnNums               int
	connNumMu                 *sync.Mutex
	HeartBeatDuration         time.Duration
	filterOfflineConnDuration time.Duration
	expiredConnsChan          chan *Conn
}

// AddConn add conn to map[uint64]*Conn
func (manager *WsConnManager) AddConn(conn *Conn) error {
	// if current conn nums == maxConnNum ,refuse this conn and put that conn object into pool
	global.Logger.Debug("user connected",zap.Uint64("uid",conn.uid))
	if manager.isUnAvailable() {
		conn.Close()
		connPool.Put(conn)
		return AddConnError
	}
	if value, ok := manager.conns.Load(conn.uid); ok {
		conn := value.(*Conn)
		conn.Close()
		manager.conns.Delete(conn.uid)
		connPool.Put(conn)
	}
	manager.conns.Store(conn.uid, conn)
	manager.connNumMu.Lock()
	manager.curConnNums++
	manager.connNumMu.Unlock()
	// start conn upstream and downstream goroutine service
	go conn.serve()
	go conn.pushMsg()
	return nil
}

// isAvailable judge whether or not current conn nums equal to maxConnNums
func (manager *WsConnManager) isUnAvailable() (res bool) {
	manager.connNumMu.Lock()
	res = manager.curConnNums == manager.maxConnNums
	manager.connNumMu.Unlock()
	return
}

// getConnByUid get conn object by uid from map
func (manager *WsConnManager) getConnByUid(uid uint64) (*Conn, error) {
	value, ok := manager.conns.Load(uid)
	if !ok {
		return nil, UserOfflineError
	}
	return value.(*Conn), nil
}

// SendTo send message to user whose id equals to uid
func (manager *WsConnManager) SendTo(uid uint64, msg Message) error {
	conn, err := manager.getConnByUid(uid)
	if err != nil {
		return err
	}
	conn.msgChan <- msg
	return nil
}

// forEach traverse conn map and execute 'do' function which will return
// a bool value represented the function result for every conn.
func (manager *WsConnManager) forEach(do func(*Conn) bool) {
	manager.conns.Range(func(key interface{}, value interface{}) bool {
		return do(value.(*Conn))
	})
}

// FilterOfflineConns put expired conns into expiredConnsChan when ticker ticks filterOfflineConnDuration
func (manager *WsConnManager) FilterOfflineConns() {
	ticker := time.NewTicker(manager.filterOfflineConnDuration)
	for {
		select {
		case <-ticker.C:
			manager.forEach(func(conn *Conn) bool {
				if conn.expiredTime.After(time.Now()) {
					manager.expiredConnsChan <- conn
				}
				return true
			})
		}
	}
}

// HandleOfflineConns handle conns from offlineChan continously
func (manager *WsConnManager) HandleOfflineConns() {
	for {
		select {
		case conn := <-manager.expiredConnsChan:
			manager.conns.Delete(conn.uid)
			if conn == nil{
				global.Logger.Error("conn == nil")
			}
			conn.Close()
			connPool.Put(conn)
		}
	}
}

func NewWsConnManager() *WsConnManager {
	connNum, err := global.GetGlobalConfig().GetConfigByPath("server.max-wsconn-num")
	if err != nil {
		global.Logger.Error("get server.MaxWSConnNum error", zap.Error(err))
		return nil
	}
	num, err := strconv.Atoi(connNum)
	if err != nil {
		global.Logger.Error("server.MaxWSConnNum format error")
		return nil
	}
	heartBeat, err := global.GetGlobalConfig().GetConfigByPath("server.websocket.health-check-duration")
	if err != nil {
		global.Logger.Error("server.websocket.health-check-duration get error", zap.Error(err))
		return nil
	}
	heartBeatDur, err := strconv.Atoi(heartBeat)
	if err != nil {
		global.Logger.Error("server.websocket.health-check-duration format error", zap.Error(err))
		return nil
	}
	filterDuration, err := global.GetGlobalConfig().GetConfigByPath("server.websocket.filter-offline-conn-duration")
	if err != nil {
		global.Logger.Error("server.websocket.filter-offline-conn-duration get error", zap.Error(err))
		return nil
	}
	filterDurationInt, err := strconv.Atoi(filterDuration)
	if err != nil {
		global.Logger.Error("server.websocket.filter-offline-conn-duration format error", zap.Error(err))
		return nil
	}
	return &WsConnManager{
		conns:                     new(sync.Map),
		maxConnNums:               num,
		curConnNums:               0,
		connNumMu:                 new(sync.Mutex),
		HeartBeatDuration:         time.Duration(heartBeatDur) * time.Second,
		filterOfflineConnDuration: time.Duration(filterDurationInt) * time.Second,
		expiredConnsChan:          make(chan *Conn),
	}
}