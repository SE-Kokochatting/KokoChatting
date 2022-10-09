package service

import "github.com/gorilla/websocket"

type WsService struct{

}

func (srv *WsService) AddConn(conn *websocket.Conn) error {
	return nil
}