package gateway

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

//Broadcast ..
type Broadcast struct {
	tp.PushCtx
}

//Push ..
func (b *Broadcast) Push(bmsg *msg.Msg) *tp.Rerror {
	mlog.Println(bmsg)
	uid := bmsg.GetUid()
	tid := bmsg.GetTid()
	buff, err := bmsg.Marshal()
	if err != nil {
		clientDistribute(uid, tid, buff)
	}
	mlog.Println(err)
	return nil
}

//分发数据到客户端，语音数据，其他控制指令
func clientDistribute(uid, tid uint32, buff []byte) {
	userser, isExist := localTopicBroadcast.Load(tid)
	if !isExist {
		mlog.Println("broadcast topic is not exist")
	}
	topic, ok := userser.(*usersOfTopic)
	if !ok {
		mlog.Println("no ok")
	}
	mlog.Println(topic)
	topic.RLock()
	for id, sendStream := range topic.users {
		if id != uid {
			mlog.Printf("broadcast uid=%d streamID=%d\n", id, sendStream.StreamID())
			if _, err := sendStream.Write(buff); err != nil {
				mlog.Println(err)
			}
		}
	}
	topic.RUnlock()
}
