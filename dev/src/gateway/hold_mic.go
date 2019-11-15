package gateway

import (
	"yt/ytproto/msg"
)

func (y *ytClientInfo) newHoldMic(message *msg.Msg) error {
	buff, err := y.holdMic(message)
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
func (y *ytClientInfo) holdMic(message *msg.Msg) ([]byte, error) {
	mlog.Println("hold mic")
	var result int32
	mid := message.GetMid()
	uid := message.GetUid()
	tid := message.GetTid()
	processHoldMic(uid, tid, &result)
	if result > 20 {
		mlog.Printf("mid=%d uid=%d hold mic in tid=%d result=%d\n", mid, uid, tid, result)
		return send2cliPack(message, msg.CMDID_HoldMicAck, result)
	}
	if rerr := y.tpSession.Call("/manager/holdmic", message, &result).Rerror(); rerr != nil {
		mlog.Println("call manager failed", rerr.String())
		result = 100
	}

	buff, err := send2cliPack(message, msg.CMDID_HoldMicAck, result)
	if err != nil {
		mlog.Printf("pack error:%v", err)
		return nil, err
	}
	// 抢麦成功
	if result == 1 {
		mlog.Println("broadcast command to other holdmic")
		clientDistribute(uid, tid, buff) //广播给当前网关的其他客户端端
		mlog.Printf("mid=%d uid=%d hold mic in tid=%d success result=%d\n", mid, uid, tid, result)
		return buff, nil
	}
	mlog.Printf("mid=%d uid=%d hold mic in tid=%d failed result=%d\n", mid, uid, tid, result)
	return buff, nil
}

func processHoldMic(uid, tid uint32, result *int32) {
	topicer, isExist := localTopicBroadcast.Load(tid)
	if !isExist {
		*result = 21
		return
	}
	topic, ok := topicer.(*usersOfTopic)
	if !ok {
		*result = 22
		return
	}
	topic.Lock()
	if topic.holder == 0 || topic.holder == uid {
		topic.holder = uid
		*result = 1
	} else {
		*result = 21
	}
	topic.Unlock()
}
