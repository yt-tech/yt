package manager

import (
	"errors"
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

var errorOK = errors.New("error")

//Connect ..
func (m *Manager) Connect(request *msg.ConnectInfo) (result int32, err *tp.Rerror) {
	uid := request.GetUid()
	mlog.Printf("uid=%d connect", uid)
	if _, hadExist := users.LoadOrStore(uid, true); hadExist {
		result = 2
		return result, nil
	}
	return 1, nil
}
