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
			panic(err)
		}
		mlog.Printf("streamID=%d\n", y.quicStream.StreamID())
		readBuff := make([]byte, 1024)
		var buff []byte
		for {
			n, err := y.quicStream.Read(readBuff)
			if err != nil {
				mlog.Println(err)
				break
			}
			message, ok := msgPool.Get().(*msg.Msg)
			if !ok {
				break
			}
			err = ggproto.Unmarshal(readBuff[:n], message)
			if err != nil {
				mlog.Println(err)
				msgPool.Put(message)
				break
			}
			switch message.GetMid() {
			case msg.MsgID_ConnectID:
				if err = y.connectRequest(message.Command.GetConnect()); err != nil {
					mlog.Println("sess close")
					y.quicSession.Close()
				}
			case msg.MsgID_SubscribeTopicID:
				y.subscribeTopic(message)
			case msg.MsgID_UnsubscribeTopicID:
			case msg.MsgID_HoldMicID:
				// if buff, err = gt.holdMic(rpcsess, message); err != nil {
				// 	msgPool.Put(message)
				// 	break
				// }
			case msg.MsgID_ReleaseMicID:
			case msg.MsgID_DisconnectID:
			case msg.MsgID_AudioDataID:
				mlog.Println("----audio----")
				y.audio(message)
			default:
				mlog.Println("--------")
				msgPool.Put(message)
				break
			}
			msgPool.Put(message)
			y.quicStream.Write(buff)
		}
	}
}
