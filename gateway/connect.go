package gateway

import (
	"context"
	"yt/proto"

	"github.com/smallnest/rpcx/client"
)

type gtInfo struct {
	action *proto.ActionRequest
	reply  uint8
}

func (c *gtInfo) connectAction(xcli client.XClient) uint8 {
	err := xcli.Call(context.Background(), "Login", c.action, &c.reply)
	if err != nil {
		mlog.Fatalf("failed to call: %v", err)
		return 1
	}
	switch c.reply {
	case 1:
		mlog.Printf("reply=%d", c.reply)
	case 2:
		mlog.Printf("reply=%d", c.reply)
	default:
		mlog.Printf("reply=%d", c.reply)
	}
	return c.reply
}
