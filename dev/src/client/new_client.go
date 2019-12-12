package client

import (
	"yt/ytproto/msg"

	"github.com/lucas-clemente/quic-go"
)

var client *clientInfo

type userInfo struct {
	uid uint32
	tid uint32
}
type clientInfo struct {
	uid         uint32
	tid         uint32
	state       uint8
	accessToken string
	statusToken map[uint32]*baseToken
	quicSession quic.Session
	quicStream  quic.Stream
	inChan      chan msg.Msg
	outChan     chan *msg.Msg
}

//NewClient ..
func NewClient(uid, tid uint32, atkn string) {
	client = &clientInfo{
		uid:         uid,
		tid:         tid,
		accessToken: atkn,
	}
	client.statusToken = make(map[uint32]*baseToken, 5)
	client.inChan = make(chan msg.Msg, 5)
	client.outChan = make(chan *msg.Msg, 5)
}
