package client

import (
	"yt/proto"

	ggproto "github.com/gogo/protobuf/proto"
)

func packConnectData() ([]byte, error) {
	lr := &proto.ActionRequest{
		ActionID: 1,
		Uid:      3,
		Gid:      10,
	}
	return ggproto.Marshal(lr)
}

func packJoinGroupData() ([]byte, error) {
	lr := &proto.ActionRequest{
		ActionID: 3,
		Uid:      3,
		Gid:      10,
	}
	return ggproto.Marshal(lr)
}
func packLeaveGroupData() ([]byte, error) {
	lr := &proto.ActionRequest{
		ActionID: 5,
		Uid:      3,
		Gid:      10,
	}
	return ggproto.Marshal(lr)
}
func packHoldMicData() ([]byte, error) {
	lr := &proto.ActionRequest{
		ActionID: 7,
		Uid:      3,
		Gid:      10,
	}
	return ggproto.Marshal(lr)
}
func packReleaseMicData() ([]byte, error) {
	lr := &proto.ActionRequest{
		ActionID: 9,
		Uid:      3,
		Gid:      10,
	}
	return ggproto.Marshal(lr)
}

func packDisconnectData() ([]byte, error) {
	lr := &proto.ActionRequest{
		ActionID: 11,
		Uid:      3,
		Gid:      10,
	}
	return ggproto.Marshal(lr)
}
