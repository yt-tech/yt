package client

import (
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
)

func packAudioData() ([]byte, error) {
	cm := &msg.Msg{
		Mid: msg.MsgID_AudioDataID,
		AudioData: &msg.AudioData{
			Id:   124,
			Uid:  2,
			Tid:  1,
			Data: []byte("audio"),
		},
	}
	return ggproto.Marshal(cm)
}
