package req

type SendMsgReq struct {
	Receiver       uint64 `json:"receiver"`
	MessageContent string `json:"messageContent"`
	MessageType    int    `json:"messageType"`
}

type PullOutlineMsgReq struct {
	LastMessageId uint64 `json:"lastMessageId"`
}

type PullMsgReq struct {
	LastMessageId uint64 `json:"lastMessageId"`
	Id            uint64 `json:"id"`
	MsgType       int    `json:"msgType"`
}

type RevertMsgReq struct {
	MsgId uint64 `json:"msgid"`
}
