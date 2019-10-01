package gateway

import (
	"context"
	command "yt/ytproto/cmd"

	"github.com/smallnest/rpcx/client"
)

const gatewayID = uint32(1)
const gatewayUDPListener = "127.0.0.1:9003"
const quicaddr = "127.0.0.1:9002"

//RegisteInfo ..
type RegisteInfo struct {
	args  *argsInfo
	reply uint8
}
type argsInfo struct {
	ID       uint32
	Listener string
}
type gateway struct {
	cmdMsg *command.Msg
	result uint8
}

func (g *RegisteInfo) registe2manager(xcli client.XClient) {
	err := xcli.Call(context.Background(), "GatewayRegiste", g.args, &g.reply)
	if err != nil {
		mlog.Fatalf("failed to call: %v", err)
	}
	switch g.reply {
	case 1:
		mlog.Printf("gwRegiste reply=%d", g.reply)
	case 2:
		mlog.Printf("gwRegiste reply=%d", g.reply)
	default:
		mlog.Printf("gwRegiste reply=%d", g.reply)
	}
}
