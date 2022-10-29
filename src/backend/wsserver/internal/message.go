package internal

type Message interface {
	Bytes() []byte
	MessageType() int
	GetUid() uint64
}

type SingleMessage struct {
	content []byte
	msgType int
	fromuid uint64
	touid   uint64
}

func (sm *SingleMessage) Bytes() []byte {
	return sm.content
}

func (sm *SingleMessage) MessageType() int {
	return sm.msgType
}

func (sm *SingleMessage) GetUid() uint64 {
	return sm.touid
}

func (sm *SingleMessage) Set(from, to uint64, content []byte, msgType int) {
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