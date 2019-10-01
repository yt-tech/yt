package gateway

import (
	"context"

	"github.com/smallnest/rpcx/client"
)

func (c *gateway) holdMic(xcli client.XClient) {

	err := xcli.Call(context.Background(), "HoldMic", c.cmdMsg, &c.result)
	if err != nil {
		mlog.Fatalf("failed to call: %v", err)
	}
	switch c.result {
	case 1:
		mlog.Printf("holdMic result=%d", c.result)
	case 2:
		mlog.Printf("holdMic result=%d", c.result)
	default:
		mlog.Printf("holdMic result=%d", c.result)
	}
}
