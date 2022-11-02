package provider

import (
	"KokoChatting/global"
	"KokoChatting/model/dataobject"
	"fmt"

	"go.uber.org/zap"
)

type MsgPullProvider struct {
	mysqlProvider
	*ManageProvider
}

func (m *MsgPullProvider) GetUidMessage(uid, lastId uint64, uidMsgList *[]dataobject.Message) error {
	// 首先从message中查找到ToId = uid的消息
	// 然后截取id > LastMessageId 并且没有被撤回的消息
	// 然后将消息封装到返回体中
	dbClient := m.mysqlProvider.mysqlDb
	if dbClient == nil {
		return fmt.Errorf("the db client is nil")
	}
	err := dbClient.Where("to_id=? and id > ? and is_revert = ?", uid, lastId, 0).Find(uidMsgList).Error
	if err != nil {
		global.Logger.Error("find message error", zap.Error(err))
		return err
	}
	return nil
}

// GetMessage find messages whose id > lastMsgId and toId = uid and fromId = fromId and msgType = msgType
func (m *MsgPullProvider) GetMessage(uid, lastMsgId, fromId uint64, msgType int) ([]dataobject.Message, error) {
	var messages []dataobject.Message
	dbClient := m.mysqlProvider.mysqlDb
	if dbClient == nil {
		return nil, fmt.Errorf("the db client is nil")
	}
	// 如果消息已经被撤回，不用拉取给用户
	err := dbClient.Where("id > ? and to_id = ? and from_id = ? and type = ? and is_revert = ?", lastMsgId, uid, fromId, msgType, 0).Find(&messages).Error
	if err != nil {
		global.Logger.Error("find message error", zap.Error(err))
		return nil, nil
	}
	return messages, nil
}

// GetMessage find messages whose id < firstMsgId and toId = uid and fromId = fromId and msgType = msgType
func (m *MsgPullProvider) GetMessageHistory(uid, firstMsgId, fromId uint64, msgType, pageNum, pageSize int) ([]dataobject.Message, error) {
	var messages []dataobject.Message
	dbClient := m.mysqlProvider.mysqlDb
	if dbClient == nil {
		return nil, fmt.Errorf("the db client is nil")
	}
	// paging return
	err := dbClient.Limit(pageSize).Offset((pageNum-1)*pageSize).
		Where("id < ? and to_id = ? and from_id = ? and type = ?", firstMsgId, uid, fromId, msgType).Find(&messages).Error
	if err != nil {
		global.Logger.Error("find message error", zap.Error(err))
		return nil, nil
	}
	return messages, nil
}

func NewMsgPullProvider() *MsgPullProvider {
	return &MsgPullProvider{
		mysqlProvider:  *NewMysqlProvider(),
		ManageProvider: NewManageProvider(),
	}
}
