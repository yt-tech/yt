package manager

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

//Holdmic ..
func (m *Manager) Holdmic(ytmsg *msg.Msg) (result int32, terr *tp.Rerror) {
	uid := ytmsg.GetUid()
	tid := ytmsg.GetTid()
	topicer, isEsixt := topics.Load(tid)
	if !isEsixt {
		return 105, nil
	}
	topic, ok := topicer.(*topicInfo)
	if !ok {
		return 0, tp.NewRerror(11, "断言失败", "")
	}
	topic.Lock()
	if topic.uid != 0 && topic.uid != uid {
		mlog.Printf("uid=%d hold mic in tid=%d failed current holder=%d\n", uid, tid, topic.uid)
		return 100, nil
	}
	topic.uid = uid
	topic.Unlock()
	topics.Store(tid, topic)

	//给相关网关广播

	cmdBroadcast(topic.gateways, ytmsg)
	mlog.Printf("uid=%d hold mic in tid=%d success\n", uid, tid)
	return 1, nil
}
