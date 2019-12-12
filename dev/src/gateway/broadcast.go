package gateway

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

//Broadcast ..
type Broadcast struct {
	tp.PushCtx
}

//Push ..
func (b *Broadcast) Push(bmsg *msg.Msg) *tp.Rerror {
	mlog.Println(bmsg)
	uid := bmsg.GetUid()
	tid := bmsg.GetTid()
	buff, err := bmsg.Marshal()
	if err != nil {
		mlog.Println(err)
		return tp.NewRerror(11, "打包异常", "")
	}
	clientDistribute(uid, tid, buff)
	return nil
}

//分发数据到客户端，语音数据，其他控制指令
func clientDistribute(uid, tid uint32, buff []byte) {
	userser, isExist := localTopicBroadcast.Load(tid)
	if !isExist {
		mlog.Println("broadcast topic is not exist")
		return
	}
	topic, ok := userser.(*usersOfTopic)
	if !ok {
		mlog.Println("no ok")
		return
	}
	mlog.Println(topic)
	topic.RLock()
	for id, sendStream := range topic.usersBroadcastStream {
		if id != uid {
			mlog.Printf("broadcast uid=%d streamID=%d\n", id, sendStream.StreamID())
			_, err := sendStream.Write(buff)
			if err != nil {
				mlog.Println(err)
				break
			}
		}
	}
	topic.RUnlock()
}

func getTopic(tid uint32) *usersOfTopic {
	userser, isExist := localTopicBroadcast.Load(tid)
	if !isExist {
		mlog.Println("broadcast topic is not exist")
		return nil
	}
	topic, ok := userser.(*usersOfTopic)
	if !ok {
		mlog.Println("no ok")
		return nil
	}
	return topic
}
func (b *Broadcast) Systembroadcast(bmsg *msg.Msg) *tp.Rerror {
	mlog.Println(bmsg)
	buff, err := bmsg.Marshal()
	if err != nil {
		mlog.Println(err)
		return tp.NewRerror(11, "打包异常", "")
	}
	tid := bmsg.GetTid()
	topic := getTopic(tid)
	mlog.Println(topic)
	if topic == nil {
		return tp.NewRerror(12, "打包异常", "")
	}
	topic.RLock()
	for id, broadcastStream := range topic.usersBroadcastStream {
		mlog.Printf("systembroadcast uid=%d streamID=%d\n", id, broadcastStream.StreamID())
		_, err := broadcastStream.Write(buff)
		if err != nil {
			mlog.Println(err)
			break
		}
	}
	topic.RUnlock()
	return nil
}

func releasenotice() {

}
