package client

import "yt/ytproto/msg"

//SignIn ..
func SignIn() int32 {
	t := client.signIn()
	t.WaitTimeout(3e9)
	return t.ack
}
func (c *clientInfo) signIn() *baseToken {
	mid := getRangNumber()
	m := &msg.Msg{
		Mid:   mid,
		CmdID: msg.CMDID_SignIn,
		Uid:   c.uid,
		Token: c.accessToken,
	}
	btoken := newBaseToken()
	c.statusToken[mid] = btoken
	c.outChan <- m
	return btoken
}
