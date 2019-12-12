package gateway

import (
	"yt/ytproto/msg"

	"github.com/lucas-clemente/quic-go"
)

func (y *ytClientInfo) newSubscribeTopic(message *msg.Msg) {
	buff, err := y.subscribeTopic(message)
	if err != nil {
		mlog.Println(err)
		return
	}
	_, err = y.commandStream.Write(buff)
	if err != nil {
		mlog.Println(err)
		return
	}
	mlog.Printf("uid=%d tid=%d subscribeTopic success\n", y.uid, message.GetTid())
}
func (y *ytClientInfo) subscribeTopic(message *msg.Msg) ([]byte, error) {
	mlog.Println(message)

	var result int32
	rerr := y.tpSession.Call("/manager/subscribetopic", message, &result).Rerror()
	if rerr.ToError() != nil {
		mlog.Println(rerr.String())
		return nil, rerr.ToError()
	}
	tid := message.GetTid()
	y.preStorageTopicBroadcastStream(tid, &result)
	buff, err := send2cliPack(message, msg.CMDID_SubscribeTopicAck, result)
	if err != nil {
		mlog.Println(err)
		return nil, err
	}
	y.currentTopic = tid
	y.topicPushServerAddr = pushAddr
	// clientDistribute(y.uid, tid, buff)
	return buff, nil
}

// 预存topic广播流地址
func (y *ytClientInfo) preStorageTopicBroadcastStream(tid uint32, r *int32) {

	topicer, isExist := localTopicBroadcast.Load(tid)
	if !isExist {
		newTopic := &usersOfTopic{
			usersBroadcastStream: make(map[uint32]quic.SendStream, 10),
		}
		newTopic.usersBroadcastStream[y.uid] = y.broadcastStream
		localTopicBroadcast.Store(tid, newTopic)
		*r = 1
		return
	}
	topic, ok := topicer.(*usersOfTopic)
	if !ok {
		*r = 103
		return
	}
	topic.Lock()
	topic.usersBroadcastStream[y.uid] = y.broadcastStream
	topic.Unlock()
	// localTopicBroadcast.Store(tid, topic)
	// tmp
	localTopicBroadcast.Range(func(k, v interface{}) bool {
		mlog.Println(k.(uint32), ":", v.(*usersOfTopic))
		return true
	})
}
