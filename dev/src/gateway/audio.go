package gateway

import (
	"yt/ytproto/msg"
)

func (y *ytClientInfo) audioReceive(data *msg.Msg) {
	buff, err := data.Marshal()
	if err != nil {
		mlog.Println(err)
	}
	_, err = ublconn.WriteToUDP(buff, y.topicPushServerAddr)
	if err != nil {
		mlog.Println(err)
	}
}

func broadcastAudio(m *msg.Msg) {
	tid := m.GetTid()
	tstreamer, isExist := localTopicBroadcast.Load(tid)
	if !isExist {
		mlog.Printf("tid=%d is not exist", tid)
		return
	}
	buff, _ := m.Marshal()
	ts := tstreamer.(*usersOfTopic)
	ts.RLock()
	for _, v := range ts.users {
		_, err := v.Write(buff)
		if err != nil {
			mlog.Println(err)
		}
	}
	ts.RUnlock()
}
