package manager

import (
	"context"
	"yt/ytproto"
)

func (m *Manager) GatewayRegiste(ctx context.Context, li *ytproto.ActionRequest, reply *uint8) error {
	*reply = 1
	return nil
}