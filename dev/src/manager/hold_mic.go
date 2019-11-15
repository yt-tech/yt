package manager

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

//Holdmic ..
func (m *Manager) Holdmic(ytmsg *msg.Msg) (result int32, terr *tp.Rerror) {
	uid := ytmsg.GetUid()
	tid := ytmsg.GetTid()
	mlog.Printf("uid=%d hold mic in tid=%d\n", uid, tid)
	if topicer, isEsixt := topics.Load(tid); isEsixt {
		if topic, ok := topicer.(*topicInfo); ok {
			topic.Lock()
			if topic.holder == 0 || topic.holder == uid {
				topic.holder = uid
				result = 1
				//给相关网关广播
				cmdBroadcast(topic.gateways, ytmsg)
			} else {
				result = 100
			}
			topic.Unlock()
			topics.Store(tid, topic)
			return result, nil
		}
		return 110, tp.NewRerror(11, "断言失败", "")
	}
	return 105, nil
}
