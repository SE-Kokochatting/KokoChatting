package req

type SendMsgReq struct{
	Reciever uint64 `json:"reciever"`
	MessageContent string `json:"messageContent"`
	MessageType int `json:"messageType"`
}