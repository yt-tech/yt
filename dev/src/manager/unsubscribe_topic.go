package manager

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

//Unsubscribetopic ..
func (m *Manager) Unsubscribetopic(requestMsg *msg.Msg) (result int32, terr *tp.Rerror) {
	mlog.Println("Subscribe Topic Request")
	gwID := m.Session().ID()
	gwSessioner, isExist := gatewayBroadcastAddrs.Load(gwID)
	if !isExist {
		return 0, tp.NewRerror(12, "网关不存在", "")
	}
	gwSession, ok := gwSessioner.(tp.Session)
	if !ok {
		return 0, tp.NewRerror(11, "断言失败", "")
	}
	var topic *topicInfo
	uid := requestMsg.GetUid()
	tid := requestMsg.GetTid()
	topicer, isEsixt := topics.Load(tid)
	if isEsixt {
		if topic, ok = topicer.(*topicInfo); ok {
			topic.Lock()
			topic.users[uid] = true
			mlog.Println(topic.gateways)
			for k, v := range topic.gateways {
				mlog.Println(k, v)
				broadcast(v, requestMsg)
			}
			topic.gateways[gwID] = gwSession
			topic.Unlock()
			topics.Store(tid, topic)
			return 2, nil
		}
		return 0, tp.NewRerror(11, "断言失败", "")
	}
	topic = &topicInfo{
		users:    make(map[uint32]bool, 20),
		gateways: make(map[string]tp.Session, 5),
	}
	topic.gateways[gwID] = gwSession
	topic.users[uid] = true
	topics.Store(tid, topic)
	return 1, nil
}
