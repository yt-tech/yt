package gateway

import (
	"context"
	managerproto "yt/rpcproto"
)

func (c *gtInfo) joinGroup(cli managerproto.DataClient) {
	mlog.Println("gateway joinGroup")
	var i = &managerproto.GroupToInfo{
		Uid: 3,
		Gid: 10,
	}
	cli.JoinGroup(context.Background(), i)
}
