package internal

import (
	"KokoChatting/global"
	req "KokoChatting/model/req"
	"encoding/json"
	"go.uber.org/zap"

	"github.com/gorilla/websocket"
)

var handlers = make(map[int]msgHandler)

func init(){
	registerHandler(websocket.PingMessage,&pingHandler{})
	registerHandler(websocket.PongMessage,&pongHandler{})
	registerHandler(websocket.TextMessage,&textHandle{})
}

func registerHandler(msgType int,handle msgHandler){
	handlers[msgType] = handle
}

func getHandler(msgType int)msgHandler{
	return handlers[msgType]
}

type msgHandler interface{
	handle(msg []byte,conn *Conn) error
}


type pingHandler struct{

}


func (handle *pingHandler) handle(msg []byte,conn *Conn)error{
	conn.mu.Lock()
	conn.expiredTime.Add(conn.heartBeatDuration)
	conn.mu.Unlock()
	global.Logger.Debug("ping handle",zap.String("msg",string(msg)))
	return conn.WriteMessage(websocket.TextMessage,msg)
}



type pongHandler struct {

}


func (handle *pongHandler) handle(msg []byte,conn *Conn)error{
	global.Logger.Debug("pong handler handle msg")
	return nil
}



type textHandle struct{

}


func (handle *textHandle) handle(msg []byte,conn *Conn)error{
	pingReq := &req.PingReq{}
	err := json.Unmarshal(msg,pingReq)
	global.Logger.Debug("text handle",zap.String("msg",string(msg)))
	if err != nil{
		//
		global.Logger.Debug(string(msg))
		return nil
	}
	s := &struct{
		MsgType int `json:"messageType"`
		Time string
	}{
		MsgType: global.PongMessage,
		Time: pingReq.Time,
	}
	bs,err := json.Marshal(s)
	if err != nil{
		//
		global.Logger.Error(string(msg))
		return nil
	}
	return getHandler(websocket.PingMessage).handle(bs,conn)
}