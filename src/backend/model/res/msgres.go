package res

import "time"

type SendMessageRes struct {
	Data struct {
		Msgid uint64 `json:"msgid"`
	} `json:"data"`
}

type WsMessage struct {
	From     uint64
	MsgType  int
	Contents string
	To       uint64
}

type PullOutlineMsgRes struct {
	Data struct {
		Message []MegOutlineInfo `json:"message"`
	} `json:"data"`
}

type MegOutlineInfo struct {
	SenderId        uint64    `json:"senderId"`
	GroupId         uint64    `json:"groupId"`
	Name            string    `json:"name"`
	AvatarUrl       string    `json:"avatarUrl"`
	MessageType     int       `json:"messageType"`
	MessageNum      int       `json:"messageNum"`
	LastMessageTime time.Time `json:"lastMessageTime"`
}

type PullMsgRes struct {
	Data struct {
		Message []MessageInfo `json:"message"`
	} `json:"data"`
}

type MessageInfo struct {
	MessageId      uint64    `json:"messageId"`
	SenderId       uint64    `json:"senderId"`
	GroupId        uint64    `json:"groupId"`
	Name           string    `json:"name"`
	AvatarUrl      string    `json:"avatarUrl"`
	SendTime       time.Time `json:"sendTime"`
	MessageContent string    `json:"messageContent"`
	MessageType    int       `json:"messageType"`
	ReadUids       string    `json:"readUids"`
}


type PongMessage struct {
	From     int    `json:"From"`
	MsgType  int    `json:"MsgType"`
	Contents string `json:"Contents"`
	To       int    `json:"To"`
}