package wsserver

import (
	"KokoChatting/wsserver/internal"

	"github.com/gorilla/websocket"
)

var server *WsServer

func init(){
	server = newWsServer()
}

type WsServer struct{
	connManager *internal.WsConnManager
	msgManager *internal.MessageManager
}

func (s *WsServer) Run(){
	go s.connManager.FilterOfflineConns()
	go s.connManager.HandleOfflineConns()
	go s.msgManager.PushMsg(s.connManager)
}


func (s *WsServer) AddConn(conn *websocket.Conn,uid int64)error{
	return s.connManager.AddConn(internal.NewConn(conn,uid,s.connManager.HeartBeatDuration))
}

func (s *WsServer) SendMessage(msg Message){
	msgs := convertToInternalMsgs(msg)
	for _,m := range msgs{
		s.msgManager.AddMessage(m)
	}
}

func newWsServer()*WsServer{
	return &WsServer{
		connManager: internal.NewWsConnManager(),
		msgManager: internal.NewMessageManager(),
	}
}


func Server()*WsServer{
	return server
}


func Run(){
	server.Run()
}


func AddConn(conn *websocket.Conn,uid int64)error{
	return server.AddConn(conn,uid)
}

func PushMessage(msg Message){
	server.SendMessage(msg)
}