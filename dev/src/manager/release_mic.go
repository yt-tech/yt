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
	if !isEsixt {
		return 105, nil
	}
	topic, ok := topicer.(*topicInfo)
	if !ok {
		return 110, tp.NewRerror(11, "断言失败", "")
	}
	topic.Lock()
	newUser := topic.users[uid]
	if topic.uid != 0 && topic.uid != uid {
		return 111, nil
	}
	if topic.priority > newUser.priority {
		return 112, nil
	}
	topic.uid = 0
	result = 1
	for k, v := range topic.gateways {
		mlog.Println(k, v)
		broadcast(v, ytmsg)
	}
	topic.Unlock()
	return result, nil
}
