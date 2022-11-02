package req

type SendMsgReq struct {
	Receiver       uint64 `json:"receiver" binding:"required"`
	MessageContent string `json:"messageContent"`
	MessageType    int    `json:"messageType"`
}

type PullOutlineMsgReq struct {
	LastMessageId uint64 `json:"lastMessageId"`
}

type PullMsgReq struct {
	LastMessageId uint64 `json:"lastMessageId" binding:"required"`
	Id            uint64 `json:"id" binding:"required"`
	MsgType       int    `json:"msgType"`
}

type RevertMsgReq struct {
	MsgId uint64 `json:"msgid" binding:"required"`
}

type PullMsgHsyReq struct {
	Id             uint64 `json:"id" binding:"required"`
	FirstMessageId uint64 `json:"firstMessageId"`
	MsgType        int    `json:"msgType"`
	PageNum        int    `json:"pageNum"`
	PageSize       int    `json:"pageSize"`
}
