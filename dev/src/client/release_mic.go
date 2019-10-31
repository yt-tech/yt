package client

import (
	"time"
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
)

func (c *clientInfo) packReleaseMic(mid uint32) ([]byte, error) {
	cm := &msg.Msg{
		Mid:   mid,
		CmdID: msg.CMDID_ReleaseMic,
		Uid:   c.uid,
		Tid:   c.tid,
	}

	return ggproto.Marshal(cm)
}

func (c *clientInfo) releaseMic() error {
	mid := getRangNumber()
	data, err := c.packReleaseMic(mid)
	if err != nil {
		mlog.Println(err)
		return err
	}
	c.quicStream.Write(data)
	c.quicStream.SetReadDeadline(time.Now().Add(3e9))
	bf := make([]byte, 1024)
	n, err := c.quicStream.Read(bf)
	if err == nil {
		cm := &msg.Msg{}
		err = ggproto.Unmarshal(bf[:n], cm)
		if err != nil {
			mlog.Println(err)
		}
		if cm.Mid == mid {
			mlog.Println(cm)
		} else {
			mlog.Println("timeout data")
			ne, err := c.quicStream.Read(bf)
			mlog.Println(bf[:ne], err)
		}
	}
	mlog.Println(err)
	return err
}
