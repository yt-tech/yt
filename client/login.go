package client

import (
	"yt/proto"

	ggproto "github.com/gogo/protobuf/proto"
)

func getLoginData() ([]byte, error) {
	lr := &proto.ActionRequest{
		Id:  1,
		Uid: 22,
		Gid: 10,
	}
	return ggproto.Marshal(lr)
}
