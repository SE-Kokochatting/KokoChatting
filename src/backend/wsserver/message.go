package wsserver

import "KokoChatting/wsserver/internal"

type Message interface {
	Bytes() []byte
	MessageType() int
	GetUids() []int64
	FromUid() int64
}


func convertToInternalMsgs(msg Message) []internal.Message {
	var res []internal.Message
	for _,uid := range msg.GetUids(){
		singleMsg := internal.GetSingleMessage()
		singleMsg.Set(msg.FromUid(),uid,msg.Bytes(),msg.MessageType())
		res = append(res,singleMsg)
	}
	return res
}