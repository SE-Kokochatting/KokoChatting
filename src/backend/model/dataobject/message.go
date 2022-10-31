package dataobject

import (
	"time"

	"github.com/gorilla/websocket"
)

type Message struct{
	Id uint64 `gorm:"primaryKey"`
	FromId uint64
	ToId uint64
	Contents string
	SendTime time.Time
	IsRevert bool
	Type int
}



type CommonMessage struct{
	From uint64
	Tos []uint64
	Contents []byte
}

func (msg *CommonMessage) FromUid() uint64 {
	return msg.From
}


func (msg *CommonMessage) MessageType()int{
	return websocket.TextMessage
}

func (msg *CommonMessage) Bytes() []byte {
	return msg.Contents
}

func (msg *CommonMessage) GetUids() []uint64{
	return msg.Tos
}