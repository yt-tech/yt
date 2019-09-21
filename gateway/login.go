package gateway

import (
	"context"

	"github.com/smallnest/rpcx/client"
)

func login1(uid uint64, xcli client.XClient) {
	args := &LoginInfo{
		UID: uid,
	}
	reply := &GateWayReply{}
	err := xcli.Call(context.Background(), "Login", args, reply)
	if err != nil {
		mlog.Fatalf("failed to call: %v", err)
		return
	}
	switch reply.R {
	case 1:
		mlog.Printf("reply=%d", reply.R)
	case 2:
		mlog.Printf("reply=%d", reply.R)
	default:
		mlog.Printf("reply=%d", reply.R)
	}
}
