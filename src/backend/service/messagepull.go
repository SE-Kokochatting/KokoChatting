package service

import (
	"KokoChatting/global"
	"KokoChatting/model/dataobject"
	"KokoChatting/model/req"
	"KokoChatting/model/res"
	"KokoChatting/provider"
	"go.uber.org/zap"
	"time"
)

type MsgPullService struct {
	msgPullPro *provider.MsgPullProvider
	*provider.ManageProvider
}

func (msgPullSrv *MsgPullService) PullOutlineMsg(uid uint64, pullReq *req.PullOutlineMsgReq, pullOutlineRes *res.PullOutlineMsgRes) error {
	uidMsgs := make([]dataobject.Message, 0)

	err := msgPullSrv.msgPullPro.GetUidMessage(uid, pullReq.LastMessageId, &uidMsgs)
	if err != nil {
		return err
	}

	// store each FromId's msg outline
	senderInfo := make(map[uint64]*struct{
		msgType int
		msgTotalNum int
		lastMsgTime time.Time
	}, 0)

	for _, message := range uidMsgs {
		msgInfo, ok := senderInfo[message.FromId]
		if !ok {
			senderInfo[message.FromId] = &struct {
				msgType int
				msgTotalNum int
				lastMsgTime time.Time
			}{
				msgType:     message.Type,
				msgTotalNum: 1,
				lastMsgTime: message.SendTime,
			}
			continue
		}
		senderInfo[message.FromId].msgTotalNum++
		// update the last send time
		if msgInfo.lastMsgTime.Before(message.SendTime) {
			senderInfo[message.FromId].lastMsgTime = message.SendTime
		}
	}

	// check whether the user is in the group
	for senderId, msgInfo := range senderInfo {
		if msgInfo.msgType == global.GroupMessage {
			isInGroup, err := msgPullSrv.ManageProvider.IsInGroup(uid, senderId)
			if !isInGroup || err != nil {
				return err
			}
		}
	}

	messageList := &pullOutlineRes.Data.Message
	for senderIdKey, msgInfo := range senderInfo {
		senderId := senderIdKey
		var groupId uint64 = 0
		if msgInfo.msgType == global.GroupMessage {
			groupId = senderId
			senderId = 0
		}
		*messageList = append(*messageList, res.MegOutlineInfo{
			SenderId:        senderId,
			GroupId:         groupId,
			MessageType:     msgInfo.msgType,
			MessageNum:      msgInfo.msgTotalNum,
			LastMessageTime: msgInfo.lastMsgTime,
		})
	}

	return nil
}

func (msgPullSrv *MsgPullService) PullMsg(uid, lastMesId, fromId uint64, msgType int) (res.PullMsgRes, error) {
	var pullMsgRes res.PullMsgRes

	messages, err := msgPullSrv.msgPullPro.GetMessage(uid, lastMesId, fromId, msgType)
	if err != nil {
		global.Logger.Error("pull message error", zap.Error(err))
		return pullMsgRes, err
	}

	if msgType == global.GroupMessage {
		if isInGroup, err := msgPullSrv.ManageProvider.IsInGroup(uid, fromId); !isInGroup||err!=nil{
			return pullMsgRes, err
		}
	}

	messageList := &pullMsgRes.Data.Message
	for _, v := range messages {
		senderId := v.FromId
		var groupId uint64 = 0
		if v.Type == global.GroupMessage {
			groupId = v.FromId
			senderId = 0
		}
		*messageList = append(*messageList, res.MessageInfo{
			MessageId: v.Id,
			SenderId: senderId,
			GroupId: groupId,
			MessageContent: v.Contents,
			MessageType: v.Type,
		})
	}

	return pullMsgRes, nil
}

func NewMsgPullService() *MsgPullService{
	return &MsgPullService{
		msgPullPro: provider.NewMsgPullProvider(),
		ManageProvider: provider.NewManageProvider(),
	}
}