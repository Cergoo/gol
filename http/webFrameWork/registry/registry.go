package registry

import (
	"sync"
)

var (
	// PoolResponseBody it's a pool of a buffers
	PoolResponseBody = sync.Pool{
		New: func() interface{} {
			return make([]byte, 0, 512)
		},
	}
)
