package client

import (
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
)

func packConnectData() ([]byte, error) {
	cm := &msg.Msg{
		Mid: msg.MsgID_ConnectID,
		Uid: 222,
	}
	return ggproto.Marshal(cm)
}

// func packLeaveGroupData() ([]byte, error) {
// 	lr := &ytproto.ActionRequest{
// 		ActionID: 5,
// 		Uid:      3,
// 		Gid:      10,
// 	}
// 	return ggproto.Marshal(lr)
// }
// func packHoldMicData() ([]byte, error) {
// 	lr := &ytproto.ActionRequest{
// 		ActionID: 7,
// 		Uid:      3,
// 		Gid:      10,
// 	}
// 	return ggproto.Marshal(lr)
// }
// func packReleaseMicData() ([]byte, error) {
// 	lr := &ytproto.ActionRequest{
// 		ActionID: 9,
// 		Uid:      3,
// 		Gid:      10,
// 	}
// 	return ggproto.Marshal(lr)
// }

// func packDisconnectData() ([]byte, error) {
// 	lr := &ytproto.ActionRequest{
// 		ActionID: 11,
// 		Uid:      3,
// 		Gid:      10,
// 	}
// 	return ggproto.Marshal(lr)
// }
