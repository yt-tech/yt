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
	users    map[uint32]*userInfo
	micHolder
}
type userInfo struct {
	name     string
	role     uint8
	stauts   bool
	priority uint8
}
type micHolder struct {
	uid      uint32
	priority uint8
}

//Manager ..
type Manager struct {
	tp.CallCtx
}
