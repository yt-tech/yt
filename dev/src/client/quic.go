package client

import (
	"context"
	"time"

	"github.com/lucas-clemente/quic-go"
)

func (c *clientInfo) openQuic() {
	tlsConf := settingQuic()
	session, err := quic.DialAddr(quicServeAddr(), tlsConf, nil)
	if err != nil {
		mlog.Println(err)
	}
	c.session = session

	go func() {
		rcev, err := session.AcceptUniStream(context.Background())
		if err != nil {
			mlog.Println(err)
		}
		ba := make([]byte, 1024)
		for {
			n, err := rcev.Read(ba)
			mlog.Println(string(ba[:n]), err)
		}
	}()
	time.Sleep(1e8)
	quicStream, _ := session.OpenStream()
	c.quicStream = quicStream
	// quicStream, err = session.OpenStream()
	// if err != nil {
	// 	mlog.Println(err)
	// }
	// mlog.Printf("streamID=%d\n", quicStream.StreamID())
	// go func() {
	// 	inBuffer := make([]byte, 1024)
	// 	var message = &msg.Msg{}
	// 	for {
	// 		nn, _ := quicStream.Read(inBuffer)
	// 		mlog.Println(inBuffer[:nn])
	// 		err = ggproto.Unmarshal(inBuffer[:nn], message)
	// 		if err != nil {
	// 			mlog.Println(err)
	// 			break
	// 		}

	// 		mlog.Println(message)
	// 	}
	// }()
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
