package client

import (
	"context"

	"github.com/lucas-clemente/quic-go"
)

func (c *clientInfo) openQuic() {
	tlsConf := settingQuic()
	session, err := quic.DialAddr(quicServeAddr(), tlsConf, nil)
	if err != nil {
		mlog.Println(err)
	}
	quicStream, _ := session.OpenStreamSync(context.Background())
	c.quicSession = session
	c.quicStream = quicStream
	if err != nil {
		mlog.Println(err)
	}

}
