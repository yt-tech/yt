package client

import (
	"yt/ytproto/msg"
)

//SignOut ..
func SignOut() int32 {
	t := client.signIn()
	t.WaitTimeout(3e9)
	return t.ack
}

func (c *clientInfo) signOut() *baseToken {
	mid := getRangNumber()
	m := &msg.Msg{
		Mid:   mid,
		CmdID: msg.CMDID_SignOutAck,
		Uid:   c.uid,
	}
	btoken := newBaseToken()
	c.statusToken[mid] = btoken
	c.outChan <- m
	return btoken
}
