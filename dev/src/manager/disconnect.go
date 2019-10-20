package manager

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

//Disconnect ..
func (m *Manager) Disconnect(ytmsg *msg.Msg) (result int32, err *tp.Rerror) {
	uid := ytmsg.GetUid()
	mlog.Printf("uid=%d disconnect\n", uid)
	users.Delete(uid)
	return 1, nil
}
