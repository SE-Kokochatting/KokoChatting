package service

import (
	"KokoChatting/wsserver"

	"github.com/gorilla/websocket"
)

type WsService struct{

}

func (srv *WsService) AddConn(conn *websocket.Conn,uid int64) error {
	return wsserver.AddConn(conn,uid)
}


func (srv *WsService) SendMessage(msg wsserver.Message){
	wsserver.PushMessage(msg)
}