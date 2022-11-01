package req

type SendMsgReq struct{
	Receiver       uint64 `json:"receiver"`
	MessageContent string `json:"messageContent"`
	MessageType    int    `json:"messageType"`
}


type RevertMsgReq struct{
	MsgId uint64 `json:"msgid"`
}