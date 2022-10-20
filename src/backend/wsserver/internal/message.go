package internal

type Message interface {
	Bytes() []byte
	MessageType() int
	GetUid() int64
}

type SingleMessage struct {
	content []byte
	msgType int
	fromuid int64
	touid   int64
}

func (sm *SingleMessage) Bytes() []byte {
	return sm.content
}

func (sm *SingleMessage) MessageType() int {
	return sm.msgType
}

func (sm *SingleMessage) GetUid() int64 {
	return sm.touid
}

func (sm *SingleMessage) Set(from, to int64, content []byte, msgType int) {
	sm.content = content
	sm.fromuid = from
	sm.touid = to
	sm.msgType = msgType
}

type MessageManager struct {
	msgChan chan Message
}

func (manager *MessageManager) AddMessage(msg Message) {
	manager.msgChan <- msg
}

func (manager *MessageManager) PushMsg(connManager *WsConnManager) {
	for {
		select {
		case msg := <-manager.msgChan:
			connManager.SendTo(msg.GetUid(), msg)
		}
	}
}

func NewMessageManager() *MessageManager {
	return &MessageManager{
		msgChan: make(chan Message),
	}
}