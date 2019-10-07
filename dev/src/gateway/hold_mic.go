package gateway

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

func (g *gateway) holdMic(rpcsess tp.Session, message *msg.Msg) ([]byte, error) {
	mlog.Println("hold mic")
	var result int32
	request := message.Command.GetHoldMic()
	uid := request.GetUid()
	tid := request.GetTid()
	result = newHoldMic(uid, tid, result)
	if result > 99 {
		return holdMicBytes(message, result)
	}
	if rerr := rpcsess.Call("/manager/holdmic", message, &result).Rerror(); rerr != nil {
		mlog.Println(rerr.String())
		result = 500
	}
	return holdMicBytes(message, result)
}
func holdMicBytes(message *msg.Msg, r int32) ([]byte, error) {
	message.Mid = msg.MsgID_HoldMIcAckID
	message.Command.HoldMicAck.Result = r
	bf, err := message.Marshal()
	if err != nil {
		return nil, err
	}
	return bf, nil
}

func newHoldMic(uid, tid uint32, r int32) int32 {
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
		topic.holder = uid
		return 1
	}
	topic.Unlock()
	return 100
}