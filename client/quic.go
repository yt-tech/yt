package client

import (
	"crypto/tls"
	"io"
	"yt/proto"

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
		inBuffer := make([]byte, 10)
		for {
			_, err = io.ReadFull(stream, inBuffer)
			if err != nil {
				mlog.Println(err)
			}
			mlog.Println("Client: Got", inBuffer, stream.StreamID())
			mlog.Println(inBuffer)
			var rl proto.ActionRequest
			ggproto.Unmarshal(inBuffer, &rl)
			mlog.Println(rl)
			// inDataChannel <- inBuffer
			switch rl.GetActionID() {
			case 1:
			case 2:
				mlog.Println(rl)
			case 3:
			case 4:
				mlog.Println(rl)
			case 5:
			case 6:
				mlog.Println(rl)
			case 7:
				mlog.Println(rl)
			case 8:
			case 9:
				mlog.Println(rl)
			case 10:
			case 11:
			case 12:
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
