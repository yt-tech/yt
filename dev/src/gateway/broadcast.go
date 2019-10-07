package gateway

import (
	"sync"
	"yt/ytproto/msg"

	"github.com/lucas-clemente/quic-go"

	tp "github.com/henrylee2cn/teleport"
)

//Broadcast ..
type Broadcast struct {
	tp.PushCtx
}
type usersOfTopic struct {
	sync.RWMutex
	m      map[uint32]quic.SendStream
	holder uint32
}

var localTopicBroadcast sync.Map //tid-*usersOfTopic

//Push ..
func (b *Broadcast) Push(bmsg *msg.Msg) *tp.Rerror {
	mlog.Println(bmsg)
	return nil
}
func localBroadcastPush(uid, tid uint32, buff []byte) {
	userser, isExist := localTopicBroadcast.Load(tid)
	if !isExist {
		mlog.Println("broadcast topic is not exist")
	}
	users, ok := userser.(*usersOfTopic)
	if !ok {
		mlog.Println("no ok")
	}
	users.RLock()
	for id, sendStream := range users.m {
		if id != uid {
			sendStream.Write(buff)
		}
	}
	users.RUnlock()
}
