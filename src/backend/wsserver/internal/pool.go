package internal

import (
	"sync"
)

var connPool = sync.Pool{
	New: func() interface{} {
		return new(Conn)
	},
}
