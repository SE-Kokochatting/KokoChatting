package wsserver

import "KokoChatting/wsserver/internal"



type WsServer struct{
	connManager *internal.WsConnManager
	msgChan chan internal.Message
}

func (s *WsServer) Run(){
	go s.connManager.FilterOfflineConns()
	go s.connManager.HandleOfflineConns()
}

func (s *WsServer) HandleMsg(){
	for{
		select{
		case msg := <-s.msgChan:
			s.connManager.SendTo()
		}
	}
}

func NewWsServer()*WsServer{
	return &WsServer{
		connManager: internal.NewWsConnManager(),
	}
}