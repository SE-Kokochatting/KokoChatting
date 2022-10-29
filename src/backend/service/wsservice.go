package service

import (
	"KokoChatting/global"
	"KokoChatting/wsserver"
	"time"

	"github.com/gorilla/websocket"
)

type WsService struct{

}

func (srv *WsService) AddConn(conn *websocket.Conn,uid uint64) error {
	return wsserver.AddConn(conn,uid)
}


func (srv *WsService) SendMessage(msg wsserver.Message) error {
	t := time.NewTimer(3 * time.Second)
	select{
	case <- t.C:
		return global.MessageServerBusy
	case <- wsserver.PushMessage(msg):
		return nil
	}
}