package client

import (
	"crypto/tls"
	"time"
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
	"github.com/lucas-clemente/quic-go"
)

func openQuic() {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	session, err := quic.DialAddr(quicServeAddr(), tlsConf, nil)
	if err != nil {
		mlog.Println(err)
	}
	stream, err := session.OpenStream()
	if err != nil {
		mlog.Println(err)
	}
	mlog.Printf("streamID=%d\n", stream.StreamID())
	go func() {
		inBuffer := make([]byte, 1024)
		for {
			nn, _ := stream.Read(inBuffer)
			mlog.Println(inBuffer[:nn])
		}
	}()
	data, err := packConnectData()
	if err != nil {
		mlog.Println(err)
		return
	}
	var rl msg.Msg

	mlog.Println(err)
	stream.Write(data)
	mbf := make([]byte, 1024)
	// stream.Read(mbf)
	ggproto.Unmarshal(mbf, &rl)
	var r2 msg.Msg
	mbf2 := make([]byte, 1024)
	// go func() {
	// 	for {
	data, err = packSubscribeTopic()
	if err != nil {
		mlog.Println(err)
		return
	}
	stream.Write(data)
	time.Sleep(2000e6)
	// _, err = stream.Read(mbf2)
	ggproto.Unmarshal(mbf2, &r2)
	mlog.Println(r2, err)
	// 	}
	// }()
	go func() {
		for {
			data, err = packAudioData()
			if err != nil {
				mlog.Println(err)
				return
			}
			stream.Write(data)
			time.Sleep(12e7)
		}
	}()
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
