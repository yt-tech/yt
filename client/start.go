package client

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"os"

	quic "github.com/lucas-clemente/quic-go"
)

const addr = "localhost:4242"

const message = "foobar"

var mlog = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
var inDataChannel = make(chan []byte, 100)
var outDataChannel = make(chan []byte, 100)

//Start ..
func Start() {
	go openQuic()
	userLogin()
	select {}
}
func openQuic() {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	session, err := quic.DialAddr(addr, tlsConf, nil)
	if err != nil {
		mlog.Println(err)
	}
	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		mlog.Println(err)
	}
	go func() {
		inBuffer := make([]byte, len(message))
		for {
			_, err = io.ReadFull(stream, inBuffer)
			if err != nil {
				mlog.Println(err)
			}
			mlog.Println("Client: Got", inBuffer, stream.StreamID())
			inDataChannel <- inBuffer
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
