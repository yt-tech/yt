package gateway

import (
	"context"
	"yt/ytproto/msg"
)

func (y *ytClientInfo) newConnect(message *msg.Msg) error {
	buff, err := y.connectRequest(message)
	if err != nil {
		mlog.Println("sess close")
		y.quicSession.Close()
		return err
	}
	y.createClientBroadcastStream()
	_, err = y.commandStream.Write(buff)
	if err != nil {
		mlog.Println(err)
		return err
	}
	return nil
}
func (y *ytClientInfo) connectRequest(message *msg.Msg) ([]byte, error) {
	mlog.Println("ConnectRequest----------------------->>>>>")
	var result int32
	rerr := y.tpSession.Call("/manager/connect", message, &result).Rerror()
	if rerr.ToError() != nil {
		mlog.Println(rerr)
		return nil, rerr.ToError()
	}
	buff, err := send2cliPack(message, msg.CMDID_ConnectAck, result)
	if err != nil {
		mlog.Println(err)
		return nil, err
	}
	_, err = y.commandStream.Write(buff)
	if err != nil {
		mlog.Println(err)
		return nil, err
	}
	return buff, nil
}
func (y *ytClientInfo) createClientBroadcastStream() {
	var err error
	y.broadcastStream, err = y.quicSession.OpenUniStreamSync(context.Background())
	if err != nil {
		mlog.Println(err)
	}
	_, err = y.broadcastStream.Write([]byte("create broadcast success"))
	if err != nil {
		mlog.Println(err)
	}
	mlog.Printf("uid=%d broadcastStream streamID=%d\n", y.uid, y.broadcastStream.StreamID())
}
