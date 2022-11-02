package provider

import (
	"KokoChatting/global"
	"KokoChatting/model/dataobject"
	"KokoChatting/util"
	"encoding/json"
	"go.uber.org/zap"
)

type MessageProvider struct{
	mysqlProvider
}


func (prd *MessageProvider) StoreMessage(msg *dataobject.Message) error {
	prd.mysqlDb.Model(&dataobject.Message{}).Create(msg)
	return prd.mysqlDb.Error
}

// CheckMessageOfUser 检查消息是否是当前用户发出
func (prd *MessageProvider) CheckMessageOfUser(uid,msgid uint64)(bool,error){
	var count int
	err := prd.mysqlDb.Model(&dataobject.Message{}).Where("from_id = ? and id = ?",uid,msgid).Count(&count).Error
	return count == 1,err
}

func (prd *MessageProvider) CheckIsReverted(msgid uint64) (bool,error){
	res := new(dataobject.Message)
	err := prd.mysqlDb.Model(&dataobject.Message{}).Where("id = ?",msgid).Find(res).Error
	return res.IsRevert,err
}

// MarkMessageAsReverted 将消息标记为被撤回
func (prd *MessageProvider) MarkMessageAsReverted(msgid uint64) error {
	return prd.mysqlDb.Model(&dataobject.Message{}).Where("id = ?",msgid).Updates(dataobject.Message{
		IsRevert: true,
	}).Error
}


func (prd *MessageProvider) MarkMessageAsRead(uid,msgId uint64) (*dataobject.Message,error) {
	msg := new(dataobject.Message)
	err := prd.mysqlDb.Model(&dataobject.Message{}).Where("id = ?",msgId).Find(msg).Error
	if err != nil{
		global.Logger.Error("msg record not find")
		return nil,err
	}
	uids := make([]uint64,0)
	if msg.ReadUids != ""{
		err = json.Unmarshal([]byte(msg.ReadUids),&uids)
		if err != nil{
			return nil,err
		}
	}
	_,isExist := util.BinarySearch(uids,uid)
	if isExist{
		global.Logger.Debug("this msg has been mark as read")
		return msg,nil
	}
	uids = append(uids,uid)
	util.QuickSort(uids,0,len(uids))
	data,err := json.Marshal(uids)
	if err != nil{
		return nil,err
	}
	return msg,prd.mysqlDb.Model(&dataobject.Message{}).Where("id = ?",msgId).Updates(dataobject.Message{
		ReadUids: string(data),
	}).Error
}


func (prd *MessageProvider) GetToIdAndTypeByMsgid(msgid uint64)(uint64,int,error){
	res := &dataobject.Message{}
	err := prd.mysqlDb.Model(&dataobject.Message{}).Where("id = ?",msgid).Find(res).Error
	if err != nil{
		global.Logger.Error("query database error",zap.Error(err))
		return 0,0,err
	}
	return res.ToId,res.Type,nil
}

func NewMessageProvider()*MessageProvider{
	return &MessageProvider{
		mysqlProvider:*NewMysqlProvider(),
	}
}