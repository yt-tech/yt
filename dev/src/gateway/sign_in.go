package gateway

import (
	"context"
	"yt/ytproto/msg"
)

func (y *ytClientInfo) newConnect(message *msg.Msg) {
	result := y.accountSignup(message)
	buff, err := send2cliPack(message, msg.CMDID_SignInAck, result)
	if err != nil {
		mlog.Println(err)
	}
	_, err = y.commandStream.Write(buff)
	if err != nil {
		mlog.Println(err)
	}
	if result > 19 {
		mlog.Printf("result=%d uid=%d connect remoteAddr=%s commandstreamID=%d ,now close session\n", result, y.uid, y.quicSession.RemoteAddr().String(), y.commandStream.StreamID())
		if err := y.quicSession.Close(); err != nil {
			mlog.Println(err)
		}
		return
	}
	y.createClientBroadcastStream()
	mlog.Printf("result=%d uid=%d connect remoteAddr=%s commandstreamID=%d success\n", result, y.uid, y.quicSession.RemoteAddr().String(), y.commandStream.StreamID())
}
func (y *ytClientInfo) accountSignup(message *msg.Msg) int32 {
	//accoun auth
	var result int32
	//账号未授权
	mytoken := message.Token
	allowAccount := verifyAccount(mytoken)
	if allowAccount == false {
		return 20
	}
	//注册mangager
	rerr := y.tpSession.Call("/manager/connect", message, &result).Rerror()
	if rerr.ToError() != nil {
		mlog.Println(rerr)
		return 21
	}
	return result
}

//为新连接创建下行广播流
func (y *ytClientInfo) createClientBroadcastStream() {
	var err error
	y.broadcastStream, err = y.quicSession.OpenUniStreamSync(context.Background())
	if err != nil {
		mlog.Println(err)
	}
	// _, err = y.broadcastStream.Write([]byte("create broadcast success"))
	// if err != nil {
	// 	mlog.Println(err)
	// }
	mlog.Printf("uid=%d broadcastStream streamID=%d\n", y.uid, y.broadcastStream.StreamID())
}
