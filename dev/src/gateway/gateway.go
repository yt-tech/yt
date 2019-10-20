package gateway

import (
	"sync"
	"yt/ytproto/msg"
)

const gatewayID = uint32(1)
const gatewayUDPListener = "127.0.0.1:9003"

//RegisteInfo ..
var localTopicBroadcast sync.Map //tid-*usersOfTopic
var clientsMap sync.Map

type gateway struct{}

func send2cliPack(message *msg.Msg, mid msg.MsgID, result int32) ([]byte, error) {
	message.Mid = mid
	message.AckCode = result
	bf, err := message.Marshal()
	if err != nil {
		return nil, err
	}
	return bf, nil
}
