package manager

import (
	"context"
	command "yt/ytproto/cmd"
)

//ConnectRequest ..
func (m *Manager) ConnectRequest(ctx context.Context, cri *command.ConnectRequestInfo) (*command.ConnectResponseInfo, error) {
	uid := cri.GetUid()
	mlog.Println("connect----->", uid)
	var result int32 = 1
	_, hadExist := m.users.LoadOrStore(uid, true)
	if hadExist {
		result = 2
	}
	return &command.ConnectResponseInfo{Result: result}, nil
}
