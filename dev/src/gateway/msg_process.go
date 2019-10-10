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
}

var msgPool = sync.Pool{
	New: func() interface{} {
		return &msg.Msg{
			Mid: 0,
			Command: &msg.Command{
				Connect:      new(msg.ConnectInfo),
				Subscribe:    new(msg.SubscribeTopicInfo),
				ConnectAck:   new(msg.ConnectAckInfo),
				SubscribeAck: new(msg.SubscribeTopicAckInfo),
			},
			AudioData: &msg.AudioData{},
		}
	},
}

func process(quicsess quic.Session, rpcsess tp.Session) {
	if rpcsess == nil || quicsess == nil {
		mlog.Fatalln("rpcsess or sess is nil ")
	}
	var ytCli = new(ytClientInfo)
	for {
		stream, err := quicsess.AcceptStream(context.Background())
		if err != nil {
			panic(err)
		}
		mlog.Printf("streamID=%d\n", stream.StreamID())
		readBuff := make([]byte, 1024)
		var buff []byte
		for {
			n, err := stream.Read(readBuff)
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
				if err = gt.connectRequest(rpcsess, quicsess, stream, message.Command.GetConnect()); err != nil {
					mlog.Println("sess close")
					quicsess.Close()
				}
			case msg.MsgID_SubscribeTopicID:
				ytCli.subscribeTopic(rpcsess, message)
			case msg.MsgID_UnsubscribeTopicID:
			case msg.MsgID_HoldMicID:
				if buff, err = gt.holdMic(rpcsess, message); err != nil {
					msgPool.Put(message)
					break
				}
			case msg.MsgID_ReleaseMicID:
			case msg.MsgID_DisconnectID:
			case msg.MsgID_AudioDataID:
				mlog.Println("----audio----")
				ytCli.audio(message)
			default:
				mlog.Println("--------")
				msgPool.Put(message)
				break
			}
			msgPool.Put(message)
			stream.Write(buff)
		}
	}
}
