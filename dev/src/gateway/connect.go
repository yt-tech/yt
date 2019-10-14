package gateway

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

func (y *ytClientInfo) connectRequest(message *msg.Msg) ([]byte, error) {
	mlog.Println("ConnectRequest----------------------->>>>>")
	var result int32
	rerr := y.tpSession.Call("/manager/connect", message, &result, tp.WithAddMeta("author", "henrylee2cn")).Rerror()
	if rerr.ToError() != nil {
		mlog.Println(rerr)
		return nil, rerr.ToError()
	}
	buff, err := send2cliPack(message, msg.MsgID_ConnectAckID, result)
	if err != nil {
		mlog.Println(err)
		return nil, err
	}
	_, err = y.quicStream.Write(buff)
	if err != nil {
		mlog.Println(err)
		return nil, err
	}
	return buff, nil
}

// func connectAckBytes(message *msg.Msg, r int32) ([]byte, error) {
// 	message.Mid = msg.MsgID_ConnectAckID
// 	message.Command.ConnectAck.Result = r
// 	bf, err := message.Marshal()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bf, nil
// }
