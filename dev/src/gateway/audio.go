package gateway

import (
	"yt/ytproto/msg"
)

func (y *ytClientInfo) audio(data *msg.Msg) {
	buff, _ := data.Marshal()
	ublconn.WriteToUDP(buff, y.topicPushServerAddr)
}
