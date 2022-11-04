package service

import (
	"KokoChatting/global"
	"KokoChatting/model/dataobject"
	"KokoChatting/model/res"
	"KokoChatting/provider"
	"encoding/json"
	"go.uber.org/zap"
	"time"
)

type MessageWrapFunc func(from,to uint64,contents string)(*dataobject.CommonMessage,error)

type MessageService struct{
	*WsService
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
		MsgType: global.GroupMessage,
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
	}else if !ok{
		return nil,global.MessageSenderError
	}
	uids,err := srv.mngPrd.GetUserIdOfGroup(to)
	if err != nil{
		global.Logger.Error("query user id error", zap.Error(err))
		return nil, global.GetMemberOfGroupError
	}
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

// PushStoredSystemMessage push system message
// @return msgid,error
func (srv *MessageService) PushStoredSystemMessage(from,to uint64,contents string,msgType int) (uint64,error) {
	// wrap msg
	wrapMsg,err := srv.WrapCommonMessage(from, to, contents, msgType)
	if err != nil{
		global.Logger.Error("wrap msg error",zap.Error(err))
		return 0,err
	}

	// store msg
	msgid,err := srv.StoreMessage(from, to, contents, msgType)
	if err != nil{
		global.Logger.Error("store msg err",zap.Error(err))
		return 0,err
	}

	// push msg
	err = srv.SendMessage(wrapMsg)
	if err != nil{
		global.Logger.Error("send msg error",zap.Error(err))
		return 0,err
	}
	return msgid,nil
}


func (srv *MessageService) PushUnStoredSystemMessage(from,to uint64,contents string,msgType int) error{
	// wrap msg
	wrapMsg,err := srv.WrapCommonMessage(from, to, contents, msgType)
	if err != nil{
		global.Logger.Error("wrap msg error",zap.Error(err))
		return err
	}
	//fmt.Println(string(wrapMsg.Bytes()))

	// push msg
	err = srv.SendMessage(wrapMsg)
	if err != nil{
		global.Logger.Error("send msg error",zap.Error(err))
		return err
	}
	return nil
}

func (srv *MessageService) RevertMessage(uid,msgid uint64) error {
	if ok,err := srv.msgPrd.CheckMessageOfUser(uid,msgid);err != nil{
		global.Logger.Error("database query error",zap.Error(err))
		return global.DatabaseQueryError
	}else if !ok{
		global.Logger.Error("msg whose id equals to 'msgid' is not sent by current user")
		return global.RevertMessageError
	}
	if is,err := srv.msgPrd.CheckIsReverted(msgid);err != nil{
		global.Logger.Error("database error")
		return global.DatabaseQueryError
	}else if is{
		global.Logger.Error("msg has been reverted")
		return global.MsgHasBeenRevertedError
	}
	to,msgType,err := srv.msgPrd.GetToIdAndTypeByMsgid(msgid)
	if err != nil{
		global.Logger.Error("database query error",zap.Error(err))
		return global.DatabaseQueryError
	}
	if msgType != global.GroupMessage && msgType != global.SingleMessage{
		global.Logger.Error("only group msg and single msg can be reverted")
		return global.RevertedMessageTypeError
	}
	err = srv.msgPrd.MarkMessageAsReverted(msgid)
	if err != nil{
		global.Logger.Error("database query error",zap.Error(err))
		return global.DatabaseQueryError
	}
	bs,err := json.Marshal(map[string]interface{}{
		"revertedMsgId":msgid,
	})
	if err != nil{
		global.Logger.Error("json marshal error",zap.Error(err))
		return global.WsJsonMarshalError
	}
	switch msgType{
	case global.SingleMessage:
		_,err = srv.PushStoredSystemMessage(uid,to,string(bs),global.RevertSingleMessageNotify)
	case global.GroupMessage:
		_,err = srv.PushStoredSystemMessage(uid,to,string(bs),global.RevertGroupMessageNotify)
	}
	if err != nil{
		global.Logger.Error("push revert notify error",zap.Error(err))
		return err
	}
	return nil
}

func (srv *MessageService) DeleteMessage (msgId uint64) error {
	err := srv.msgPrd.DeleteMessage(msgId)
	if err != nil{
		global.Logger.Error("delete message err", zap.Error(err))
		return err
	}
	return err
}

func NewMessageService()*MessageService{
	srv := &MessageService{
		WsService:new(WsService),
		msgPrd: provider.NewMessageProvider(),
		mngPrd:provider.NewManageProvider(),
		msgWrapMap: make(map[int]MessageWrapFunc),
	}
	srv.register(global.SingleMessage,srv.wrapSingleMessage)
	srv.register(global.FriendRequestNotify,srv.wrapSingleMessage)
	srv.register(global.GroupMessage,srv.wrapGroupMessage)
	srv.register(global.RevertSingleMessageNotify,srv.wrapSingleMessage)
	srv.register(global.RevertGroupMessageNotify,srv.wrapGroupMessage)
	srv.register(global.QuitGroupNotify,srv.wrapGroupMessage)
	srv.register(global.JoinGroupNotify,srv.wrapGroupMessage)
	srv.register(global.AddFriendResponseNotify,srv.wrapSingleMessage)
	srv.register(global.DeleteFriendNotify,srv.wrapSingleMessage)
	return srv
}
