package gateway

import (
	"sync"
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
	"github.com/lucas-clemente/quic-go"
)

//Broadcast ..
type Broadcast struct {
	tp.PushCtx
}

type usersOfTopic struct {
	sync.RWMutex
	users  map[uint32]quic.SendStream
	holder uint32
}

//Push ..
func (b *Broadcast) Push(bmsg *msg.Msg) *tp.Rerror {
	mlog.Println(bmsg)
	uid := bmsg.GetUid()
	tid := bmsg.GetTid()
	buff, err := bmsg.Marshal()
	if err != nil {
		localBroadcastPush(uid, tid, buff)
	}
	mlog.Println(err)
	return nil
}
func localBroadcastPush(uid, tid uint32, buff []byte) {
	userser, isExist := localTopicBroadcast.Load(tid)
	if !isExist {
		mlog.Println("broadcast topic is not exist")
	}
	topic, ok := userser.(*usersOfTopic)
	mlog.Println(topic)
	if !ok {
		mlog.Println("no ok")
	}
	topic.RLock()
	for id, sendStream := range topic.users {
		mlog.Printf("uid=%d broadcast streamID=%d\n", id, sendStream.StreamID())
		if id != uid {
			_, err := sendStream.Write(buff)
			mlog.Println(err)
		}
	}
	topic.RUnlock()
}
