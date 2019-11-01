package client

import (
	"time"
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
)

func (c *clientInfo) packHoldMic(mid uint32) ([]byte, error) {
	cm := &msg.Msg{
		Mid:   mid,
		CmdID: msg.CMDID_HoldMic,
		Uid:   c.uid,
		Tid:   c.tid,
	}

	return ggproto.Marshal(cm)
}

func (c *clientInfo) holdMic() error {
	mid := getRangNumber()
	data, err := c.packHoldMic(mid)
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
			switch cm.GetAckCode() {
			case 1:
				for it := 0; it < 2; it++ {
					data, err := c.packAudioData()
					if err != nil {
						mlog.Println(err)
						return err
					}
					c.quicStream.Write(data)
					time.Sleep(12e7)
				}
			default:
				mlog.Println("failed")
			}
		} else {
			mlog.Println("timeout data")
			ne, err := c.quicStream.Read(bf)
			mlog.Println(bf[:ne], err)
		}
	}
	mlog.Println(err)
	return err
}
