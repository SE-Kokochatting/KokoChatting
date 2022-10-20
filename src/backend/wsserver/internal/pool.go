package internal

import (
	"sync"
)

var connPool = sync.Pool{
	New: func() interface{} {
		return new(Conn)
	},
}


var msgPool = sync.Pool{
	New:func() interface{}{
		return new(SingleMessage)
	},
}

func GetWsConn()*Conn{
	return connPool.Get().(*Conn)
}

func GetSingleMessage()*SingleMessage{
	return msgPool.Get().(*SingleMessage)
}
