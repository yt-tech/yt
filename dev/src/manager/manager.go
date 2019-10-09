package manager

import (
	"sync"

	tp "github.com/henrylee2cn/teleport"
)

var (
	gatewayBroadcastAddrs sync.Map
	users                 sync.Map //在线用户
	topics                sync.Map //组成员及成员网关地址
)

type topicInfo struct {
	sync.RWMutex
	gateways map[string]tp.Session
	users    map[uint32]bool
	holder   uint32
}

//Manager ..
type Manager struct {
	tp.CallCtx
}
