package gateway

import (
	"context"
	"sync"
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
	tp "github.com/henrylee2cn/teleport"
	"github.com/lucas-clemente/quic-go"
)

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
			message := msgPool.Get().(*msg.Msg)
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
				gt.subscribeTopic(rpcsess, message)
			case msg.MsgID_HoldMicID:
				if buff, err = gt.holdMic(rpcsess, message); err != nil {
					msgPool.Put(message)
					break
				}
			// case command.CommandType_ReleaseMicRequest:
			// case command.CommandType_DisconnectRequest:
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
