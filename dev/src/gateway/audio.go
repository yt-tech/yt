package gateway

import (
	"yt/ytproto/msg"
)

func (y *ytClientInfo) audio(data *msg.Msg) {
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
	tid := m.AudioData.GetTid()
	tstreamer, isExist := localTopicBroadcast.Load(tid)
	if !isExist {
		mlog.Println(tid, "is not exist")
		return
	}
	buff, _ := m.Marshal()
	ts := tstreamer.(*usersOfTopic)
	ts.RLock()
	for _, v := range ts.users {
		v.Write(buff)
	}
	ts.RUnlock()
}
