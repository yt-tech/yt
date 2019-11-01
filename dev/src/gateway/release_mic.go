package gateway

import (
	"yt/ytproto/msg"
)

func (y *ytClientInfo) newReleaseMic(message *msg.Msg) error {
	buff, err := y.releaseMic(message)
	if err != nil {
		mlog.Println(err)
		return err
	}
	_, err = y.commandStream.Write(buff)
	if err != nil {
		mlog.Println(err)
		return err
	}
	return nil
}

func (y *ytClientInfo) releaseMic(message *msg.Msg) ([]byte, error) {
	mlog.Println("hold mic")
	var result int32
	uid := message.GetUid()
	tid := message.GetTid()
	processReleaseMic(uid, tid, &result)
	if result > 99 {
		return send2cliPack(message, msg.CMDID_ReleaseMicAck, result)
	}
	if rerr := y.tpSession.Call("/manager/releasemic", message, &result).Rerror(); rerr != nil {
		mlog.Println(rerr.String())
		result = 500
	}
	return send2cliPack(message, msg.CMDID_ReleaseMicAck, result)
}

func processReleaseMic(uid, tid uint32, r *int32) int32 {
	topicer, isExist := localTopicBroadcast.Load(tid)
	if !isExist {
		return 13
	}
	topic, ok := topicer.(*usersOfTopic)
	if !ok {
		return 103
	}
	topic.Lock()
	if topic.holder == 0 || topic.holder == uid {
		topic.holder = 0
		*r = 1
	}
	topic.Unlock()
	return *r
}
