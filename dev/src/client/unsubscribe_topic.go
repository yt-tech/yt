package client

import (
	"yt/ytproto/msg"
)

//UnsubscribeTopic ..
func UnsubscribeTopic() int32 {
	t := client.unsubscribeTopic()
	t.WaitTimeout(3e9)
	return t.ack
}
func (c *clientInfo) unsubscribeTopic() *baseToken {
	mid := getRangNumber()
	m := &msg.Msg{
		Mid:   mid,
		CmdID: msg.CMDID_UnsubscribeTopic,
		Uid:   c.uid,
		Tid:   c.tid,
	}
	btoken := newBaseToken()
	c.statusToken[mid] = btoken
	c.outChan <- m
	return btoken
}
