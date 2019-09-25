package manager

import (
	"context"
	"yt/ytproto"
)

//HoldMic ..
func (m *Manager) HoldMic(ctx context.Context, li *ytproto.ActionRequest, reply *uint8) error {
	*reply = 1
	return nil
}
