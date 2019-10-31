package gateway

import (
	"yt/ytproto/msg"
)

func (y *ytClientInfo) newUnsubscribeTopic(message *msg.Msg) error {
	buff, err := y.unsubscribeTopic(message)
	if err != nil {
		return err
	}
	_, err = y.commandStream.Write(buff)
	if err != nil {
		mlog.Println(err)
		return err
	}
	return nil
}

func (y *ytClientInfo) unsubscribeTopic(message *msg.Msg) ([]byte, error) {
	mlog.Println(message)
	var result int32
	rerr := y.tpSession.Call("/manager/unsubscribetopic", message, &result).Rerror()
	if rerr.ToError() != nil {
		mlog.Println(rerr.String())
		return nil, rerr.ToError()
	}
	tid := message.GetTid()
	y.currentTopic = 0
	y.topicPushServerAddr = nil

	buff, err := send2cliPack(message, msg.CMDID_UnsubscribeTopicAck, result)
	if err != nil {
		mlog.Println(err)
		return nil, err
	}
	localBroadcastPush(y.uid, tid, buff)
	y.deleteBroadcastStream(tid, result)
	return buff, nil
}

func (y *ytClientInfo) deleteBroadcastStream(tid uint32, r int32) int32 {
	topicer, isExist := localTopicBroadcast.Load(tid)
	if !isExist {
		return 103
	}
	topic, ok := topicer.(*usersOfTopic)
	if !ok {
		return 103
	}
	topic.Lock()
	topic.users[y.uid] = nil
	topic.Unlock()
	return r
}
