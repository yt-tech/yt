package gateway

import (
	"yt/sharelib/util"
	"yt/ytproto/msg"
)

func (y *ytClientInfo) newRemoveMic() error {
	var message = &msg.Msg{
		Mid:   util.GetRangNumber(),
		CmdID: msg.CMDID_RemoveMic,
		Uid:   y.uid,
		Tid:   y.currentTopic,
	}
	err := y.removeMic(message)
	if err != nil {
		mlog.Println(err)
		return err
	}
	return nil
}

func (y *ytClientInfo) removeMic(message *msg.Msg) error {
	tid := message.GetTid()
	result := processRemoveMic(tid)
	if result > 9 {
		return nil
	}

	rerr := y.tpSession.Call("/manager/removemic", message, &result).Rerror()
	if rerr != nil {
		mlog.Println(rerr.String())
		result = 500
	}
	return nil
}

func processRemoveMic(tid uint32) int32 {
	topicer, isExist := localTopicBroadcast.Load(tid)
	if !isExist {
		return 13
	}
	topic, ok := topicer.(*usersOfTopic)
	if !ok {
		return 14
	}
	topic.Lock()
	topic.micHolder = 0
	topic.Unlock()
	return 1
}
