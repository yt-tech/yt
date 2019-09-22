package gateway

import (
	"context"
	"log"
	"yt/proto"

	"github.com/smallnest/rpcx/client"
)

type UserGroupRequestInfo struct {
	GatewayID uint32
	GroupID   uint64
	UserID    uint64
}

func joinGroupAction(a proto.ActionRequest, xcli client.XClient) int8 {
	return joinGroup(a, xcli)
}

func joinGroup(a proto.ActionRequest, xcli client.XClient) int8 {
	args := &UserGroupRequestInfo{
		GatewayID: requestGatewayID,
		GroupID:   1,
		UserID:    requestUserID,
	}
	reply := &GateWayReply{}
	err := xcli.Call(context.Background(), "JoinGroup", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
		return -1
	}
	switch reply.R {
	case 1:
		log.Printf("reply=%d", reply.R)
	case 2:
		log.Printf("reply=%d", reply.R)
	default:
		log.Printf("reply=%d", reply.R)
	}
	return reply.R
}
