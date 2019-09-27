package manager

import (
	"context"
	"yt/ytproto"
)

//Manager ..
type Manager struct{}

//UserConnect ..
func (m *Manager) UserConnect(ctx context.Context, li *ytproto.ActionRequest, reply *uint8) error {
	if loginQurey(li.GetUid()) {
		*reply = 1
	} else {
		*reply = 0
	}
	return nil
}
