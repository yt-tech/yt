package client

import (
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
)

func (c *clientInfo) cmdRecieve() {
	mlog.Println("cmdRecieve")
	var bf = make([]byte, 512)
	for {
		n, err := c.quicStream.Read(bf)
		cm := msg.Msg{}
		if err != nil {
			if err.Error() == "deadline exceeded" {
				ne, _ := c.quicStream.Read(bf)
				err := ggproto.Unmarshal(bf[:ne], &cm)
				if err != nil {
					mlog.Println(err)
				}
				mlog.Println("timeout data", cm)
			}
		}
		err = ggproto.Unmarshal(bf[:n], &cm)
		if err != nil {
			mlog.Println(err)
			continue
		}
		c.inChan <- cm
	}
}
func (c *clientInfo) handle() {
	mlog.Println("handle")
	for {
		im := <-c.inChan
		t, ok := c.statusToken[im.GetMid()]
		if ok && t != nil {
			t.flowComplete()
			t.ack = im.GetAckCode()
		} else {
			if !ok {
				mlog.Println("not ok")
			} else {
				mlog.Println(t)
			}
		}
	}
}
