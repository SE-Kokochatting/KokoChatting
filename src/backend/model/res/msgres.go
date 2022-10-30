package res

type SendMessageRes struct{
	Data struct{
		Msgid uint64 `json:"msgid"`
	} `json:"data"`
}


type WsMessage struct{
	From uint64
	MsgType int
	Contents []byte
}