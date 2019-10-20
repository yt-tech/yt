package client

import (
	"crypto/tls"
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
	"github.com/lucas-clemente/quic-go"
)

var quicStream quic.Stream

type clientInfo struct {
	uid     uint32
	tid     uint32
	session quic.Session
}

func (c *clientInfo) openQuic() {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	session, err := quic.DialAddr(quicServeAddr(), tlsConf, nil)
	if err != nil {
		mlog.Println(err)
	}
	quicStream, err = session.OpenStream()
	if err != nil {
		mlog.Println(err)
	}
	mlog.Printf("streamID=%d\n", quicStream.StreamID())
	go func() {
		inBuffer := make([]byte, 1024)
		var message = &msg.Msg{}
		for {
			nn, _ := quicStream.Read(inBuffer)
			mlog.Println(inBuffer[:nn])
			err = ggproto.Unmarshal(inBuffer[:nn], message)
			if err != nil {
				mlog.Println(err)
				break
			}

			mlog.Println(message)
		}
	}()
	// 	}
	// }()

	// if rl.Response.ConnectAck.GetResult() == 1 || rl.Response.ConnectAck.GetResult() == 2 {

	// }

	// var broadcastReadStream quic.ReceiveStream
	// mlog.Println("========================")
	// broadcastReadStream, err = session.AcceptUniStream(context.Background())
	// mlog.Println("========================")
	// if err != nil {
	// 	mlog.Println(err)
	// }
	// mlog.Println("open read stream", broadcastReadStream.StreamID())
	// mlog.Println("========================")
	// mlog.Println("open read stream")

	// go func() {
	// 	for {
	// 		n, err := broadcastReadStream.Read(inBuffer)
	// 		if err != nil {
	// 			mlog.Println(err)
	// 		}
	// 		mlog.Println("streamID", stream.StreamID())
	// 		var rl msg.Msg
	// 		if err := ggproto.Unmarshal(inBuffer[:n], &rl); err != nil {
	// 			mlog.Println(rl, err)
	// 			break
	// 		}
	// 		// inDataChannel <- inBuffer
	// 		switch rl.GetMid() {
	// 		// case msg.MsgID_SubscribeTopicRequestID:
	// 		// 	mlog.Println(rl)
	// 		default:
	// 			mlog.Println("unkown----------", broadcastReadStream.StreamID(), "------", rl)
	// 		}
	// 	}
	// }()

}
