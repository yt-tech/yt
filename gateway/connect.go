package gateway

import (
	"context"
	command "yt/ytproto/cmd"
)

func (c *gateway) connect(cli command.ManagerClient, uid uint64) (*command.ConnectResponseInfo, error) {
	var cri = &command.ConnectRequestInfo{Uid: uid}
	return cli.ConnectRequest(context.Background(), cri)
}
