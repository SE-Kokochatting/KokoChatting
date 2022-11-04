package service

import (
	"KokoChatting/global"
	"KokoChatting/model/dataobject"
	"KokoChatting/model/req"
	"KokoChatting/model/res"
	"KokoChatting/provider"
	"time"

	"go.uber.org/zap"
)

type MsgPullService struct {
	msgPullPro *provider.MsgPullProvider
	*provider.ManageProvider
	*provider.UserProvider
}

type senderMsgIdx struct {
	senderId      uint64
	senderMsgType int
}

type senderMsgInfo struct {
	msgTotalNum int
	lastMsgTime time.Time
}

func (msgPullSrv *MsgPullService) PullOutlineMsg(uid uint64, pullReq *req.PullOutlineMsgReq, pullOutlineRes *res.PullOutlineMsgRes) error {
	uidMsgs := make([]dataobject.Message, 0)

	err := msgPullSrv.msgPullPro.GetUidMessage(uid, pullReq.LastMessageId, &uidMsgs)
	if err != nil {
		return err
	}

	// store each FromId's msg outline
	senderInfo := make(map[senderMsgIdx]*senderMsgInfo, 0)

	for _, message := range uidMsgs {
		sender_key := senderMsgIdx{
			senderId:      message.FromId,
			senderMsgType: message.Type,
		}
		msgInfo, ok := senderInfo[sender_key]
		if !ok {
			senderInfo[sender_key] = &senderMsgInfo{
				msgTotalNum: 1,
				lastMsgTime: message.SendTime,
			}
			continue
		}
		senderInfo[sender_key].msgTotalNum++
		// update the last send time
		if msgInfo.lastMsgTime.Before(message.SendTime) {
			senderInfo[sender_key].lastMsgTime = message.SendTime
		}
	}

	// check whether the user is in the group
	for senderKey, _ := range senderInfo {
		if senderKey.senderMsgType == global.GroupMessage {
			isInGroup, err := msgPullSrv.ManageProvider.IsInGroup(uid, senderKey.senderId)
			if !isInGroup || err != nil {
				return err
			}
		}
	}

	messageList := &pullOutlineRes.Data.Message
	for senderKey, msgInfo := range senderInfo {
		senderId := senderKey.senderId
		var groupId uint64 = 0
		if senderKey.senderMsgType == global.GroupMessage {
			groupId = senderId
			senderId = 0
		}
		name, avatarUrl, err := msgPullSrv.GetProfileById(senderId, groupId, senderKey.senderMsgType)
		if err != nil {
			return err
		}
		*messageList = append(*messageList, res.MegOutlineInfo{
			SenderId:        senderId,
			GroupId:         groupId,
			Name:            name,
			AvatarUrl:       avatarUrl,
			MessageType:     senderKey.senderMsgType,
			MessageNum:      msgInfo.msgTotalNum,
			LastMessageTime: msgInfo.lastMsgTime,
		})
	}

	return nil
}

// get user(group) name and avatarurl by senderId(groupId)
// user or group depends on the message type
func (msgPullSrv *MsgPullService) GetProfileById(senderId, groupId uint64, messageType int) (string, string, error) {
	var name, avatarUrl string
	if messageType == global.SingleMessage {
		// get user profile
		userProfile := &dataobject.UserProfile{
			Uid: senderId,
		}
		if err := msgPullSrv.UserProvider.CheckExist(userProfile); err != nil {
			global.Logger.Error("message pull outline: get user profile err", zap.Error(err))
			return "", "", err
		}
		return userProfile.Name, userProfile.AvatarUrl, nil
	} else if messageType == global.GroupMessage {
		// get group profile
		groupProfile := &dataobject.GroupProfile{
			Gid: groupId,
		}
		if err := msgPullSrv.ManageProvider.GetGroupInfo(groupProfile); err != nil {
			global.Logger.Error("message pull outline: get group profile err", zap.Error(err))
			return "", "", err
		}
		return groupProfile.Name, groupProfile.AvatarUrl, nil
	}
	return name, avatarUrl, nil
}

func (msgPullSrv *MsgPullService) PullMsg(uid, lastMesId, fromId uint64, msgType int) (res.PullMsgRes, error) {
	var pullMsgRes res.PullMsgRes

	messages, err := msgPullSrv.msgPullPro.GetMessage(uid, lastMesId, fromId, msgType)
	if err != nil {
		global.Logger.Error("pull message error", zap.Error(err))
		return pullMsgRes, err
	}

	if msgType == global.GroupMessage {
		if isInGroup, err := msgPullSrv.ManageProvider.IsInGroup(uid, fromId); !isInGroup || err != nil {
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
			MessageId:      v.Id,
			SenderId:       senderId,
			GroupId:        groupId,
			MessageContent: v.Contents,
			MessageType:    v.Type,
			ReadUids:       v.ReadUids,
		})
	}

	return pullMsgRes, nil
}

func (msgPullSrv *MsgPullService) PullMsgHistory(uid uint64, pullMsgHsyReq *req.PullMsgHsyReq) (res.PullMsgRes, error) {
	var pullMsgRes res.PullMsgRes

	messages, err := msgPullSrv.msgPullPro.GetMessageHistory(uid, pullMsgHsyReq.
		FirstMessageId, pullMsgHsyReq.Id, pullMsgHsyReq.MsgType, pullMsgHsyReq.PageNum, pullMsgHsyReq.PageSize)
	if err != nil {
		global.Logger.Error("pull message error", zap.Error(err))
		return pullMsgRes, err
	}

	if pullMsgHsyReq.MsgType == global.GroupMessage {
		if isInGroup, err := msgPullSrv.ManageProvider.IsInGroup(uid, pullMsgHsyReq.Id); !isInGroup || err != nil {
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
			MessageId:      v.Id,
			SenderId:       senderId,
			GroupId:        groupId,
			MessageContent: v.Contents,
			MessageType:    v.Type,
			ReadUids:       v.ReadUids,
		})
	}

	return pullMsgRes, nil
}

func NewMsgPullService() *MsgPullService {
	return &MsgPullService{
		msgPullPro:     provider.NewMsgPullProvider(),
		ManageProvider: provider.NewManageProvider(),
		UserProvider:   provider.NewUserProvider(),
	}
}
