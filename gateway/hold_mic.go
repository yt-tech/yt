package gateway

import (
	"context"

	"github.com/smallnest/rpcx/client"
)

func (c *gtInfo) holdMic(xcli client.XClient) {

	err := xcli.Call(context.Background(), "HoldMic", c.action, &c.reply)
	if err != nil {
		mlog.Fatalf("failed to call: %v", err)
	}
	switch c.reply {
	case 1:
		mlog.Printf("holdMic reply=%d", c.reply)
	case 2:
		mlog.Printf("holdMic reply=%d", c.reply)
	default:
		mlog.Printf("holdMic reply=%d", c.reply)
	}
}
