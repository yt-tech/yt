package gateway

import (
	"context"
	command "yt/ytproto/cmd"
)

func (c *gateway) subscribeTopic(cli command.ManagerClient, request *command.SubscribeTopicRequestInfo) (*command.SubscribeTopicResponseInfo, error) {
	mlog.Println("gateway subscribe topic")

	return cli.SubscribeTopicRequest(context.Background(), request)
}
