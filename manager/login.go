package manager

import "context"

type Manager struct{}
type ManagerReply struct {
	R int8
}
type LoginInfo struct {
	UID uint64 `json:"uid"`
}

func (m *Manager) Login(ctx context.Context, li *LoginInfo, reply *ManagerReply) error {
	if loginQurey(li.UID) {
		reply.R = 1
	} else {
		reply.R = 0
	}
	return nil
}
