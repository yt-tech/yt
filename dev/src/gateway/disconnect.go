package gateway

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
	"github.com/lucas-clemente/quic-go"
)

func (g *gateway) disconnectRequest(rpcsess tp.Session, sess quic.Session, stream quic.Stream, request *msg.ConnectInfo) error {
	mlog.Println("ConnectRequest----------------------->>>>>")
	var result int32
	rerr := rpcsess.Call("/manager/connect", request, &result, tp.WithAddMeta("author", "henrylee2cn")).Rerror()
	if rerr.ToError() != nil {
		mlog.Println(rerr)
		return rerr.ToError()
	}
	broadcastStream, err := sess.OpenUniStream()
	if err != nil {
		mlog.Println(err)
		return err
	}
	var uid = request.GetUid()
	mlog.Println("broadcastWriteStream ID", broadcastStream.StreamID())
	usersSession.LoadOrStore(uid, sess)
	usersStream.LoadOrStore(uid, stream)
	usersBroadcastStream.LoadOrStore(uid, broadcastStream)
	buff, err := connectAckBytes(result)
	if err != nil {
		mlog.Println(err)
		return err
	}
	_, err = stream.Write(buff)
	if err != nil {
		mlog.Println(err)
		return err
	}
	return nil
}

func disconnectAckBytes(r int32) ([]byte, error) {
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
