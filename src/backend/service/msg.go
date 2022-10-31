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

type MessageWrapFunc func(from,to uint64,contents string)(*dataobject.CommonMessage,error)

type MessageService struct{
	msgPrd *provider.MessageProvider
	mngPrd *provider.ManageProvider
	msgWrapMap map[int]MessageWrapFunc
}

// private:

func (srv *MessageService) register(msgt int,f MessageWrapFunc){
	srv.msgWrapMap[msgt] = f
}


func (srv *MessageService) getMessageWrapFunc(msgt int)(MessageWrapFunc,error){
	if f,ok := srv.msgWrapMap[msgt];ok{
		return f,nil
	}
	return nil,global.MessageTypeError
}


func (srv *MessageService) wrapSingleMessage(from,to uint64,contents string)(*dataobject.CommonMessage,error){
	wsmsg := &res.WsMessage{
		From: from,
		MsgType: global.SingleMessage,
		Contents: contents,
		To: to,
	}
	c,err := json.Marshal(wsmsg)
	if err != nil{
		return nil,global.WsJsonMarshalError
	}
	if ok,err := srv.mngPrd.IsInBlock(to,from);err != nil{
		global.Logger.Error("is in block judge error",zap.Error(err))
		return nil,global.QueryBlockRelationError
	}else if ok{
		return nil,global.HasBeenBlocked
	}
	return &dataobject.CommonMessage{
		From: from,
		Tos: []uint64{to},
		Contents: c,
	},nil
}

func (srv *MessageService) wrapGroupMessage(from,to uint64,contents string)(*dataobject.CommonMessage,error){
	wsmsg := &res.WsMessage{
		From: from,
		MsgType: global.SingleMessage,
		Contents: contents,
		To: to,
	}
	c,err := json.Marshal(wsmsg)
	if err != nil{
		return nil,global.WsJsonMarshalError
	}
	if ok,err := srv.mngPrd.IsInGroup(from,to);err != nil{
		global.Logger.Error("query whether or not user is in group error",zap.Error(err))
		return nil,global.QueryIsInGroup
	}else if ok{
		return nil,global.MessageSenderError
	}
	uids,err := srv.mngPrd.GetUserIdOfGroup(to)
	return &dataobject.CommonMessage{
		From: from,
		Tos: uids,
		Contents: c,
	},nil
}


// public:

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
	f,err := srv.getMessageWrapFunc(msgType)
	if err != nil{
		return nil,err
	}
	return f(from,to,contents)
}


func NewMessageService()*MessageService{
	srv := &MessageService{
		msgPrd: provider.NewMessageProvider(),
		mngPrd:provider.NewManageProvider(),
		msgWrapMap: make(map[int]MessageWrapFunc),
	}
	srv.register(global.SingleMessage,srv.wrapSingleMessage)
	srv.register(global.FriendRequestNotify,srv.wrapSingleMessage)
	srv.register(global.GroupMessage,srv.wrapGroupMessage)
	return srv
}
