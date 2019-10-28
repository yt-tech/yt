package client

import (
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
)

func packAudioData() ([]byte, error) {
	cm := &msg.Msg{
		CmdID: msg.CMDID_Audio,
		AudioData: &msg.AudioData{
			Id:   24,
			Data: []byte("---11111---"),
		},
	}
	return ggproto.Marshal(cm)
}
