package res

import "time"

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

type PullOutlineMsgRes struct {
	Data struct {
		Message []megOutlineInfo `json:"message"`
	} `json:"data"`
}

type megOutlineInfo struct {
	SenderId uint64 `json:"senderId"`
	GroupId uint64 `json:"groupId"`
	MessageType int `json:"messageType"`
	MessageNum int `json:"messageNum"`
	LastMessageTime time.Time `json:"lastMessageTime"`
}