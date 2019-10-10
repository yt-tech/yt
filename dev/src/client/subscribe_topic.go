package client

import (
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
)

func packSubscribeTopic() ([]byte, error) {
	cm := &msg.Msg{
		Mid: msg.MsgID_SubscribeTopicID,
		Command: &msg.Command{
			Subscribe: &msg.SubscribeTopicInfo{
				Uid: 2,
				Tid: 1,
			},
		},
	}

	return ggproto.Marshal(cm)
}
