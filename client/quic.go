package client

import (
	"crypto/tls"
	command "yt/ytproto/cmd"

	ggproto "github.com/gogo/protobuf/proto"
	"github.com/lucas-clemente/quic-go"
)

func openQuic() {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	session, err := quic.DialAddr(gatewayAddr, tlsConf, nil)
	if err != nil {
		mlog.Println(err)
	}
	stream, err := session.OpenStream()
	// sendStream, err := session.OpenUniStream()
	if err != nil {
		mlog.Println(err)
	}
	go func() {
		mlog.Println("open read stream")
		inBuffer := make([]byte, 1024)
		for {
			n, err := stream.Read(inBuffer)
			if err != nil {
				mlog.Println(err)
			}
			mlog.Println("Client: Got", inBuffer[:n], stream.StreamID())
			var rl command.Msg
			ggproto.Unmarshal(inBuffer[:n], &rl)
			mlog.Println(rl)
			// inDataChannel <- inBuffer
			switch rl.GetCtype() {
			case command.CommandType_ConnectResponse:
				mlog.Println(rl.Response.ConnectAck)
			case command.CommandType_SubscribeTopicResponse:
				mlog.Println(rl.Response.SubscribeAck)
			// case 3:
			// case 4:
			// 	mlog.Println(rl)
			// case 5:
			// case 6:
			// 	mlog.Println(rl)
			// case 7:
			// 	mlog.Println(rl)
			// case 8:
			// case 9:
			// 	mlog.Println(rl)
			// case 10:
			// case 11:
			// case 12:
			default:
				mlog.Panicln("unkown")
			}
		}
	}()
	for {
		outBuffer := <-outDataChannel
		_, err = stream.Write(outBuffer)
		if err != nil {
			mlog.Println(err)
		}
	}
}
