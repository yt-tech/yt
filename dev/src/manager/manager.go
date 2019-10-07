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
	users    map[uint32]bool
	gateways map[string]tp.Session
}

//Manager ..
type Manager struct {
	tp.CallCtx
}
