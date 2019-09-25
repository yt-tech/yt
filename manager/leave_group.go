package manager

import (
	"context"
	"yt/ytproto"
)

//LeaveGroup ..
func (m *Manager) LeaveGroup(ctx context.Context, li *ytproto.ActionRequest, reply *uint8) error {
	*reply = 1
	return nil
}
