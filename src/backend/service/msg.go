package service

import (
	"KokoChatting/global"
	"KokoChatting/model/dataobject"
	"KokoChatting/model/res"
	"KokoChatting/provider"
	"encoding/json"
	"time"

	"go.uber.org/zap"
)

func init(){
	register(global.SingleMessage,WrapSingleMessage)
	// register(global.GroupMessage,)
	register(global.FriendRequestNotify,WrapSingleMessage)
}

// MessageHandleFunc
type MessageWrapFunc func(from,to uint64,contents []byte)*dataobject.CommonMessage

var MsgWrapMap = make(map[int]MessageWrapFunc)

func register(msgt int,f MessageWrapFunc){
	MsgWrapMap[msgt] = f
}


func GetMessageWrapFunc(msgt int)(MessageWrapFunc,error){
	if f,ok := MsgWrapMap[msgt];ok{
		return f,nil
	}
	return nil,global.MessageTypeError
}


func WrapSingleMessage(from,to uint64,contents []byte)*dataobject.CommonMessage{
	wsmsg := &res.WsMessage{
		From: from,
		MsgType: global.SingleMessage,
		Contents: contents,
	}
	c,err := json.Marshal(wsmsg)
	if err != nil{
		return nil
	}
	return &dataobject.CommonMessage{
		From: from,
		Tos: []uint64{to},
		Contents: c,
	}
}



type MessageService struct{
	msgPrd *provider.MessageProvider
}


func (srv *MessageService) StoreMessage(from,to uint64,contents string,msgType int) (uint64,error) {
	entity := &dataobject.Message{
		FromId: from,
		ToId: to,
		Contents: contents,
		SendTime: time.Now(),
		IsRevert: false,
		Type: msgType,
	}
	err := srv.msgPrd.StoreMessage(entity)
	if err != nil{
		global.Logger.Error("store msg error",zap.Error(err))
		return 0,global.StoreMessageError
	}
	return entity.Id,nil
}


func (srv *MessageService) WrapCommonMessage(from,to uint64,contents string,msgType int)(*dataobject.CommonMessage,error){
	f,err := GetMessageWrapFunc(msgType)
	if err != nil{
		return nil,err
	}
	return f(from,to,[]byte(contents)),nil
}


func NewMessageService()*MessageService{
	return &MessageService{

	}
}



