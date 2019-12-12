package client

import (
	"yt/ytproto/msg"
)

//SubscribeTopic  ..
func SubscribeTopic() int32 {
	t := client.subscribeTopic()
	t.WaitTimeout(3e9)
	return t.ack
}
func (c *clientInfo) subscribeTopic() *baseToken {
	mid := getRangNumber()
	m := &msg.Msg{
		Mid:   mid,
		CmdID: msg.CMDID_SubscribeTopic,
		Uid:   c.uid,
		Tid:   c.tid,
	}
	btoken := newBaseToken()
	c.statusToken[mid] = btoken
	c.outChan <- m
	return btoken
}
