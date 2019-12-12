package gateway

import (
	"sync"
	"yt/ytproto/msg"

	"github.com/lucas-clemente/quic-go"
)

const gatewayID = uint32(1)
const gatewayUDPListener = "127.0.0.1:9003"

//RegisteInfo ..
var localTopicBroadcast sync.Map //tid-*usersOfTopic
var clientsMap sync.Map

type gateway struct{}

type usersOfTopic struct {
	sync.RWMutex
	usersBroadcastStream map[uint32]quic.SendStream
	micHolder            uint32
}

func send2cliPack(message *msg.Msg, cid msg.CMDID, result int32) ([]byte, error) {
	message.CmdID = cid
	message.AckCode = result
	bf, err := message.Marshal()
	if err != nil {
		return nil, err
	}
	return bf, nil
}
