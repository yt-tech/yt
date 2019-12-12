package manager

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

//Removemic ..
func (m *Manager) Removemic(ytmsg *msg.Msg) (result int32, terr *tp.Rerror) {
	mlog.Println("sys remove mic")
	tid := ytmsg.GetTid()
	topicer, isExist := topics.Load(tid)
	if !isExist {
		return 111, nil
	}
	topic, ok := topicer.(*topicInfo)
	if !ok {
		return 110, tp.NewRerror(11, "断言失败", "")
	}
	topic.Lock()
	topic.uid = 0
	for k, v := range topic.gateways {
		mlog.Println(k, v)
		systembroadcast(v, ytmsg)
	}
	topic.Unlock()
	return 1, nil
}
