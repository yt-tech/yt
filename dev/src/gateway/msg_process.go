package gateway

import (
	"context"
	"net"
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
	tp "github.com/henrylee2cn/teleport"
	"github.com/lucas-clemente/quic-go"
)

type ytClientInfo struct {
	uid                 uint32
	currentTopic        uint32
	topicPushServerAddr *net.UDPAddr
	quicSession         quic.Session
	commandStream       quic.Stream
	broadcastStream     quic.SendStream
	tpSession           tp.Session
}

func process(quicSession quic.Session, tpSession tp.Session) {
	if tpSession == nil || quicSession == nil {
		mlog.Fatalln("rpcsess or sess is nil ")
	}
	var y = new(ytClientInfo)
	y.quicSession = quicSession
	y.tpSession = tpSession

	var err error
	y.commandStream, err = y.quicSession.AcceptStream(context.Background())
	if err != nil {
		y.closenet()
		mlog.Println(err)
		return
	}
	readBuff := make([]byte, 1024)
	var buff []byte
	var message = &msg.Msg{}
	for {
		n, err := y.commandStream.Read(readBuff)
		if err != nil {
			y.closenet()
			mlog.Println(err, y.quicSession.RemoteAddr().String(), y.quicSession.LocalAddr())
			return
		}
		// mlog.Println(readBuff[:n])
		err = ggproto.Unmarshal(readBuff[:n], message)
		if err != nil {
			mlog.Println(err)
			break
		}
		y.uid = message.GetUid()
		switch message.GetCmdID() {
		case msg.CMDID_Connect:
			if err = y.newConnect(message); err == nil {
				mlog.Printf("uid=%d connect remoteAddr=%s commandstreamID=%d success\n", y.uid, quicSession.RemoteAddr().String(), y.commandStream.StreamID())
			}
		case msg.CMDID_SubscribeTopic:
			if err = y.newSubscribeTopic(message); err == nil {
				mlog.Printf("uid=%d tid=%d subscribeTopic success\n", y.uid, message.GetTid())
			}
		case msg.CMDID_UnsubscribeTopic:
			if err = y.newUnsubscribeTopic(message); err == nil {
				mlog.Printf("uid=%d tid=%d unSubscribeTopic success\n", y.uid, message.GetTid())
			}
		case msg.CMDID_HoldMic:
			mlog.Println(message)
			if buff, err = y.holdMic(message); err != nil {
				break
			}
			y.commandStream.Write(buff)
		case msg.CMDID_ReleaseMic:
		case msg.CMDID_Disconnect:
		case msg.CMDID_Audio:
			mlog.Println("audio data")
			y.audioReceive(message)
		case msg.CMDID_Ping:
			// mlog.Println(message)
		default:
			mlog.Println("--------")
			break
		}
	}

}
