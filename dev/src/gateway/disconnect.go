package gateway

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
	"github.com/lucas-clemente/quic-go"
)

func (g *gateway) disconnectRequest(rpcsess tp.Session, sess quic.Session, stream quic.Stream, message *msg.Msg) error {
	mlog.Println("disconnectRequest----------------------->>>>>")
	var result int32
	rerr := rpcsess.Call("/manager/disconnect", message, &result).Rerror()
	if rerr.ToError() != nil {
		mlog.Println(rerr)
		return rerr.ToError()
	}
	broadcastStream, err := sess.OpenUniStream()
	if err != nil {
		mlog.Println(err)
		return err
	}
	mlog.Println("broadcastWriteStream ID", broadcastStream.StreamID())
	buff, err := send2cliPack(message, msg.MsgID_DisConnectAckID, result)
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

func (y *ytClientInfo) closenet() {
	u, isExist := clientsMap.Load(y.quicSession.RemoteAddr().String())
	if u != nil {
		uu := u.(uint32)
		mlog.Printf("user=%d(%s) not recent network activity", uu, y.quicSession.RemoteAddr().String())
		if isExist {
			t, _ := localTopicBroadcast.Load(y.currentTopic)
			tt := t.(*usersOfTopic)
			tt.Lock()
			delete(tt.users, uu)
			tt.Unlock()
		}
	}
	clientsMap.Delete(y.quicSession.RemoteAddr().String())
}
