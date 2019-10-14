package gateway

import (
	"yt/ytproto/msg"

	"github.com/lucas-clemente/quic-go"
)

func (y *ytClientInfo) subscribeTopic(message *msg.Msg) ([]byte, error) {
	mlog.Println("gateway subscribe topic")

	var result int32
	rerr := y.tpSession.Call("/manager/subscribetopic", message, &result).Rerror()
	if rerr.ToError() != nil {
		mlog.Println(rerr.String())
		return nil, rerr.ToError()
	}
	request := message.Command.GetSubscribe()
	uid := request.GetUid()
	tid := request.GetTid()
	result = y.preStorageTopicBroadcastStream(uid, tid, result)
	buff, err := send2cliPack(message, msg.MsgID_SubscribeTopicAckID, result)
	if err != nil {
		mlog.Println(err)
		return nil, err
	}
	y.currentTopic = tid
	y.topicPushServerAddr = pushAddr
	localBroadcastPush(uid, tid, buff)
	return buff, nil
}

// 预存topic广播流地址
func (y *ytClientInfo) preStorageTopicBroadcastStream(uid, tid uint32, r int32) int32 {

	topicer, isExist := localTopicBroadcast.Load(tid)
	mlog.Println(tid, isExist, "|||||||||||||")
	if !isExist {
		newTopic := &usersOfTopic{
			users: make(map[uint32]quic.Stream, 10),
		}
		newTopic.users[uid] = y.quicStream
		localTopicBroadcast.Store(tid, newTopic)
		mlog.Println(tid, isExist, "|||||||||---------||||", newTopic)
		return 12
	}
	topic, ok := topicer.(*usersOfTopic)
	if !ok {
		return 103
	}
	topic.Lock()
	topic.users[uid] = y.quicStream
	topic.Unlock()
	// localTopicBroadcast.Store(tid, topic)
	return r
}
