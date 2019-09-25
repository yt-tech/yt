package manager

import (
	"context"
	"yt/ytproto"
)

//JoinGroup ..
func (m *Manager) JoinGroup(ctx context.Context, li *ytproto.ActionRequest, reply *uint8) error {
	*reply = 1
	return nil
}
