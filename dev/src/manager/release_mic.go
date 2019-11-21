package manager

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

//Releasemic ..
func (m *Manager) Releasemic(ytmsg *msg.Msg) (result int32, terr *tp.Rerror) {
	mlog.Println("release mic")
	uid := ytmsg.GetUid()
	tid := ytmsg.GetTid()
	topicer, isEsixt := topics.Load(tid)
	if isEsixt {
		if topic, ok := topicer.(*topicInfo); ok {
			topic.Lock()
			if topic.holder == 0 || topic.holder == uid {
				topic.holder = 0
				result = 1
			}
			for k, v := range topic.gateways {
				mlog.Println(k, v)
				broadcast(v, ytmsg)
			}
			topic.Unlock()
			topics.Store(tid, topic)
			return result, nil
		}
		return 110, tp.NewRerror(11, "断言失败", "")
	}
	return 5, nil
}
