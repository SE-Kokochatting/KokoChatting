package provider

import (
	"KokoChatting/model/dataobject"
)

type MessageProvider struct{
	mysqlProvider
}


func (prd *MessageProvider) StoreMessage(msg *dataobject.Message) error {
	prd.mysqlDb.Model(&dataobject.Message{}).Create(msg)
	return prd.mysqlDb.Error
}

func NewMessageProvider()*MessageProvider{
	return &MessageProvider{
		mysqlProvider:*NewMysqlProvider(),
	}
}