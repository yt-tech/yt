package client

import (
	"yt/ytproto/msg"
)

//HoldMic  ..
func HoldMic() int32 {
	t := client.holdMic()
	t.WaitTimeout(3e9)
	return t.ack
}

func (c *clientInfo) holdMic() *baseToken {
	mid := getRangNumber()
	m := &msg.Msg{
		Mid:   mid,
		CmdID: msg.CMDID_HoldMic,
		Uid:   c.uid,
		Tid:   c.tid,
	}
	btoken := newBaseToken()
	c.statusToken[mid] = btoken
	c.outChan <- m
	return btoken
}
