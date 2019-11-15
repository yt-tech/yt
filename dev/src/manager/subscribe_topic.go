package manager

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

//Subscribetopic ..
func (m *Manager) Subscribetopic(ytmsg *msg.Msg) (result int32, terr *tp.Rerror) {
	gwID := m.Session().ID()
	mlog.Println("Subscribe Topic Request ", gwID)
	gwSession, rerr := getGWSession(gwID)
	if rerr != nil {
		return 100, rerr
	}
	uid := ytmsg.GetUid()
	tid := ytmsg.GetTid()

	var topic *topicInfo
	topicer, isEsixt := topics.Load(tid)

	if !isEsixt {
		newCreateTopic(tid, uid, gwID, gwSession)
		return 1, nil
	}
	topic, ok := topicer.(*topicInfo)
	if ok {
		topic.addrNewMember(tid, uid, gwID, gwSession, ytmsg)
		return 2, nil
	}
	return 0, tp.NewRerror(11, "断言失败", "")
}
func newCreateTopic(tid, uid uint32, gwID string, gwSession tp.Session) {
	topic := &topicInfo{
		users:    make(map[uint32]bool, 20),
		gateways: make(map[string]tp.Session, 5),
	}
	topic.gateways[gwID] = gwSession
	topic.users[uid] = true
	topics.Store(tid, topic)
}
func (t *topicInfo) addrNewMember(tid, uid uint32, gwID string, gwSession tp.Session, ytmsg *msg.Msg) {
	t.Lock()
	t.users[uid] = true
	t.gateways[gwID] = gwSession
	cmdBroadcast(t.gateways, ytmsg)
	t.Unlock()
	topics.Store(tid, t)
}
func getGWSession(gwid string) (tp.Session, *tp.Rerror) {
	gwSessioner, isExist := gatewayBroadcastAddrs.Load(gwid)
	if !isExist {
		return nil, tp.NewRerror(12, "网关不存在", "")
	}
	gwSession, ok := gwSessioner.(tp.Session)
	if !ok {
		return nil, tp.NewRerror(11, "断言失败", "")
	}
	return gwSession, nil
}

//command broadcast
func cmdBroadcast(tg map[string]tp.Session, ytmsg *msg.Msg) {
	for k, v := range tg {
		mlog.Println(k, v)
		broadcast(v, ytmsg)
	}
}
