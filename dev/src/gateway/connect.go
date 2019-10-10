package gateway

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

func (y *ytClientInfo) connectRequest(request *msg.ConnectInfo) error {
	mlog.Println("ConnectRequest----------------------->>>>>")
	var result int32
	rerr := y.tpSession.Call("/manager/connect", request, &result, tp.WithAddMeta("author", "henrylee2cn")).Rerror()
	if rerr.ToError() != nil {
		mlog.Println(rerr)
		return rerr.ToError()
	}
	// broadcastStream, err := sess.OpenUniStream()
	// if err != nil {
	// 	mlog.Println(err)
	// 	return err
	// }
	// var uid = request.GetUid()
	// mlog.Println("broadcastWriteStream ID", broadcastStream.StreamID())
	// usersSession.LoadOrStore(uid, sess)
	// usersStream.LoadOrStore(uid, stream)
	// usersBroadcastStream.LoadOrStore(uid, broadcastStream)
	buff, err := connectAckBytes(result)
	if err != nil {
		mlog.Println(err)
		return err
	}
	_, err = y.quicStream.Write(buff)
	if err != nil {
		mlog.Println(err)
		return err
	}
	return nil
}

func connectAckBytes(r int32) ([]byte, error) {
	ack := msgPool.Get().(*msg.Msg)
	ack.Mid = msg.MsgID_ConnectID
	ack.Command.ConnectAck.Result = r
	bf, err := ack.Marshal()
	msgPool.Put(ack)
	if err != nil {
		return nil, err
	}
	return bf, nil
}
