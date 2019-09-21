package gateway

import (
	"context"
	"log"

	"github.com/smallnest/rpcx/client"
)

type UserGroupRequestInfo struct {
	GatewayID uint32
	GroupID   uint64
	UserID    uint64
}

func joinGroup(xcli client.XClient) {
	args := &UserGroupRequestInfo{
		GatewayID: requestGatewayID,
		GroupID:   1,
		UserID:    requestUserID,
	}
	reply := &GateWayReply{}
	err := xcli.Call(context.Background(), "JoinGroup", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
		return
	}
	switch reply.R {
	case 1:
		log.Printf("reply=%d", reply.R)
	case 2:
		log.Printf("reply=%d", reply.R)
	default:
		log.Printf("reply=%d", reply.R)
	}
}
