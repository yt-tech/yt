package gateway

import (
	"sync"

	quic "github.com/lucas-clemente/quic-go"
)

var userlist = struct {
	sync.RWMutex
	ul map[uint64]quic.SendStream
}{ul: make(map[uint64]quic.SendStream, 20000)}
