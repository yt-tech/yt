package gateway

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"

	"github.com/lucas-clemente/quic-go"
)

func (y *ytClientInfo) subscribeTopic(rpcsess tp.Session, requestMsg *msg.Msg) {
	mlog.Println("gateway subscribe topic")

	var result int32
	rerr := rpcsess.Call("/manager/subscribetopic", requestMsg, &result).Rerror()
	if rerr.ToError() != nil {
		mlog.Println(rerr.String())
		return
	}
	request := requestMsg.Command.GetSubscribe()
	uid := request.GetUid()
	tid := request.GetTid()
	result = preStorageTopicBroadcastStream(uid, tid, result)
	streamer, ok := usersStream.Load(uid)
	if !ok {
		mlog.Println("streamer is not exist")
		return
	}
	stream, ok := streamer.(quic.Stream)
	if !ok {
		mlog.Println("get stream err")
		return
	}
	buff, err := subscriberTopicBytes(result)
	if err != nil {
		mlog.Println(err)
		return
	}
	y.tid = tid
	_, err = stream.Write(buff)
	if err != nil {
		mlog.Println(err)
		return
	}
	localBroadcastPush(uid, tid, buff)
}

func subscriberTopicBytes(r int32) ([]byte, error) {
	ack := msgPool.Get().(*msg.Msg)
	ack.Mid = msg.MsgID_SubscribeTopicID
	ack.Command.SubscribeAck.Result = r
	bf, err := ack.Marshal()
	msgPool.Put(ack)
	if err != nil {
		return nil, err
	}
	return bf, nil
}

// 预存topic广播流地址
func preStorageTopicBroadcastStream(uid, tid uint32, r int32) int32 {
	sendStreamer, isExist := usersBroadcastStream.Load(uid)
	if !isExist {
		return 100
	}
	sendStream, ok := sendStreamer.(quic.SendStream)
	if !ok {
		return 101
	}
	topicer, isExist := localTopicBroadcast.Load(tid)
	if !isExist {
		newTopic := &usersOfTopic{
			users: make(map[uint32]quic.SendStream, 50),
		}
		localTopicBroadcast.Store(tid, newTopic)
		return 12
	}
	topic, ok := topicer.(*usersOfTopic)
	if !ok {
		return 103
	}
	topic.Lock()
	topic.users[uid] = sendStream
	topic.Unlock()
	return r
}
