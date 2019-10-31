package client

import (
	"time"
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
)

func (c *clientInfo) packDisconnectData(mid uint32) ([]byte, error) {
	cm := &msg.Msg{
		Mid:   mid,
		CmdID: msg.CMDID_Disconnect,
		Uid:   c.uid,
	}
	return ggproto.Marshal(cm)
}

func (c *clientInfo) disconnect() error {
	mid := getRangNumber()
	data, err := c.packDisconnectData(mid)
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
			mlog.Println(string(bf[:ne]), err)
			cm := &msg.Msg{}
			err = ggproto.Unmarshal(bf[:n], cm)
			if err != nil {
				mlog.Println(err)
			}
			mlog.Println(cm)
		}
	}
	mlog.Println(err)
	return err
}
