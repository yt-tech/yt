package gateway

import (
	"context"

	"github.com/smallnest/rpcx/client"
)

func (c *gtInfo) connect(xcli client.XClient) (uint8, error) {
	err := xcli.Call(context.Background(), "UserConnect", c.action, &c.reply)
	if err != nil {
		mlog.Fatalf("failed to call: %v", err)
		return 0, err
	}
	return c.reply, nil
}
