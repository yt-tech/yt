package client

import (
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
)

func (c *clientInfo) packAudioData() ([]byte, error) {
	cm := &msg.Msg{
		CmdID: msg.CMDID_Audio,
		Uid:   c.uid,
		Tid:   c.tid,
		AudioData: &msg.AudioData{
			Id:   24,
			Data: []byte("---11111---"),
		},
	}
	return ggproto.Marshal(cm)
}
