package manager

import (
	"context"
	"yt/proto"
)

//Manager ..
type Manager struct{}

func (m *Manager) userConnect(ctx context.Context, li *proto.ActionRequest, reply uint8) error {
	if loginQurey(li.GetUid()) {
		reply = 1
	} else {
		reply = 0
	}
	return nil
}
