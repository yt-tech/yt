package gateway

import (
	"context"

	"github.com/smallnest/rpcx/client"
)

func (c *gtInfo) disconnect(xcli client.XClient) {
	err := xcli.Call(context.Background(), "Login", c.action, &c.reply)
	if err != nil {
		mlog.Fatalf("failed to call: %v", err)
	}
	switch c.reply {
	case 1:
		mlog.Printf("reply=%d", c.reply)
	case 2:
		mlog.Printf("reply=%d", c.reply)
	default:
		mlog.Printf("reply=%d", c.reply)
	}
}
