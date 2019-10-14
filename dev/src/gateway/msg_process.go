package gateway

import (
	"context"
	"net"
	"sync"
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
	tpSession           tp.Session
}

var msgPool = sync.Pool{
	New: func() interface{} {
		return &msg.Msg{
			Mid: 0,
			Command: &msg.Command{
				Connect:      &msg.ConnectInfo{},
				Subscribe:    &msg.SubscribeTopicInfo{},
				ConnectAck:   &msg.ConnectAckInfo{},
				SubscribeAck: &msg.SubscribeTopicAckInfo{},
			},
			AudioData: &msg.AudioData{},
		}
	},
}

func (y *ytClientInfo) process() {
	if y.tpSession == nil || y.quicSession == nil {
		mlog.Fatalln("rpcsess or sess is nil ")
	}
	var err error
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
			mlog.Println(message)

			switch message.GetMid() {
			case msg.MsgID_ConnectID:
				buff, err = y.connectRequest(message)
				if err != nil {
					mlog.Println("sess close")
					y.quicSession.Close()
					break
				}
				// clientsMap.LoadOrStore(y.quicSession.RemoteAddr().String(), uid)
			case msg.MsgID_SubscribeTopicID:
				buff, err = y.subscribeTopic(message)
				if err != nil {
					break
				}
			case msg.MsgID_UnsubscribeTopicID:
			case msg.MsgID_HoldMicID:
				if buff, err = y.holdMic(message); err != nil {
					msgPool.Put(message)
					break
				}
			case msg.MsgID_ReleaseMicID:
			case msg.MsgID_DisconnectID:
			case msg.MsgID_AudioDataID:
				mlog.Println("----audio----")
				y.audio(message)
			default:
				mlog.Println("--------")
				break
			}
			msgPool.Put(message)
			y.quicStream.Write(buff)
		}
	}
}
