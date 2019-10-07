package push

import "sync"

//GroupList ..
var GroupList sync.Map

type GroupInfo struct {
	ContainGateWays map[uint64]struct{}
}
