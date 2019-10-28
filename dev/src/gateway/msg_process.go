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
	currentTopic        uint32
	topicPushServerAddr *net.UDPAddr
	quicSession         quic.Session
	quicStream          quic.Stream
	broadcastStream     quic.SendStream
	tpSession           tp.Session
}

func (y *ytClientInfo) process() {
	if y.tpSession == nil || y.quicSession == nil {
		mlog.Fatalln("rpcsess or sess is nil ")
	}
	var err error
	y.broadcastStream, err = y.quicSession.OpenUniStream()
	if err != nil {
		mlog.Println(err)
	}
	// go func() {
	// 	tike := time.NewTicker(10e9)
	// 	cm := &msg.Msg{
	// 		Mid:   88,
	// 		CmdID: msg.CMDID_Ping,
	// 		Uid:   8,
	// 	}
	// 	pingBytes, _ := ggproto.Marshal(cm)
	// 	for {
	// 		select {
	// 		case <-tike.C:
	// 			y.broadcastStream.Write(pingBytes)
	// 		}
	// 	}
	// }()
	for {
		y.quicStream, err = y.quicSession.AcceptStream(context.Background())
		if err != nil {
			y.closenet()
			mlog.Println(err)
			return
		}

		mlog.Printf("start streamID=%d\n", y.quicStream.StreamID())
		readBuff := make([]byte, 1024)
		var buff []byte
		var message = &msg.Msg{}
		for {
			n, err := y.quicStream.Read(readBuff)
			if err != nil {
				y.closenet()
				mlog.Println(err, y.quicSession.RemoteAddr().String(), y.quicSession.LocalAddr())
				return
			}
			mlog.Println(readBuff[:n])
			err = ggproto.Unmarshal(readBuff[:n], message)
			if err != nil {
				mlog.Println(err)
				break
			}

			switch message.GetCmdID() {
			case msg.CMDID_Connect:
				mlog.Println("connect")
				buff, err = y.connectRequest(message)
				if err != nil {
					mlog.Println("sess close")
					y.quicSession.Close()
					break
				}
				y.quicStream.Write(buff)
				// clientsMap.LoadOrStore(y.quicSession.RemoteAddr().String(), uid)
			case msg.CMDID_SubscribeTopic:
				mlog.Println("SubscribeTopic")
				if buff, err = y.subscribeTopic(message); err != nil {
					break
				}
				y.quicStream.Write(buff)
			case msg.CMDID_UnsubscribeTopic:
				mlog.Println("UnSubscribeTopic")
			case msg.CMDID_HoldMic:
				mlog.Println("holdMIc")
				if buff, err = y.holdMic(message); err != nil {
					break
				}
				y.quicStream.Write(buff)
			case msg.CMDID_ReleaseMic:
			case msg.CMDID_Disconnect:
			case msg.CMDID_Audio:
				mlog.Println("audio data")
				y.audioReceive(message)
			case msg.CMDID_Ping:
				mlog.Println(message)
			default:
				mlog.Println("--------")
				break
			}
		}
	}
}
