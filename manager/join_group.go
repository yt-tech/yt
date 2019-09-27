package manager

import (
	"context"
	managerproto "yt/rpcproto"
)

//Manager ..
type Manager struct{}

//JoinGroup ..
func (m *Manager) JoinGroup(ctx context.Context, request *managerproto.GroupToInfo) (*managerproto.GroupToAckInfo, error) {
	mlog.Println("manager JoinGroup")
	var nl = &managerproto.GroupToAckInfo{
		Id: 22,
	}
	return nl, nil
}
