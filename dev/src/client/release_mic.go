package client

import (
	"yt/ytproto/msg"
)

//ReleaseMic ..
func ReleaseMic() int32 {
	t := client.releaseMic()
	t.WaitTimeout(3e9)
	return t.ack
}

func (c *clientInfo) releaseMic() *baseToken {
	mid := getRangNumber()
	m := &msg.Msg{
		Mid:   mid,
		CmdID: msg.CMDID_ReleaseMic,
		Uid:   c.uid,
		Tid:   c.tid,
	}
	btoken := newBaseToken()
	c.statusToken[mid] = btoken
	c.outChan <- m
	return btoken
}
