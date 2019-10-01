package client

import (
	command "yt/ytproto/cmd"

	ggproto "github.com/gogo/protobuf/proto"
)

func packSubscribeTopic() ([]byte, error) {
	cm := &command.Msg{
		Ctype: command.CommandType_SubscribeTopicRequest,
		Request: &command.Request{
			Subscribe: &command.SubscribeTopicRequestInfo{
				Uid: 1,
			},
		},
	}

	return ggproto.Marshal(cm)
}
