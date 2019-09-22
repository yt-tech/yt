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

//GatewayRoot ..
type GatewayRoot struct{}

//GateWayRegisteInfo ..
type GateWayRegisteInfo struct {
	ID              uint32
	ReceiveListener string
}

//GateWayReply ..
type GateWayReply struct {
	R int8
}

func gwRegiste(xcli client.XClient) {

	args := &GateWayRegisteInfo{
		ID:              requestGatewayID,
		ReceiveListener: gatewayUDPListener,
	}
	reply := &GateWayReply{}
	err := xcli.Call(context.Background(), "GatewayRegiste", args, reply)
	if err != nil {
		mlog.Fatalf("failed to call: %v", err)
	}
	switch reply.R {
	case 1:
		mlog.Printf("gwRegiste reply=%d", reply.R)
	case 2:
		mlog.Printf("gwRegiste reply=%d", reply.R)
	default:
		mlog.Printf("gwRegiste reply=%d", reply.R)
	}
}
