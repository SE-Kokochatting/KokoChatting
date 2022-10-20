package internal

import (
	"KokoChatting/global"

	"github.com/gorilla/websocket"
)

var handlers = make(map[int]msgHandler)

func init(){
	registerHandler(websocket.PingMessage,&pingHandler{})
	registerHandler(websocket.PongMessage,&pongHandler{})
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
	return conn.WriteMessage(websocket.PongMessage,msg)
}



type pongHandler struct {

}


func (handle *pongHandler) handle(msg []byte,conn *Conn)error{
	global.Logger.Debug("pong handler handle msg")
	return nil
}