package gateway

import (
	"context"
	"log"

	"github.com/smallnest/rpcx/client"
)

func (c *gtInfo) leaveGroup(xcli client.XClient) {

	err := xcli.Call(context.Background(), "leaveGroup", c.action, &c.reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	switch c.reply {
	case 1:
		log.Printf("reply=%d", c.reply)
	case 2:
		log.Printf("reply=%d", c.reply)
	default:
		log.Printf("reply=%d", c.reply)
	}
}
