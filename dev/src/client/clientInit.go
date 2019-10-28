package client

import (
	"github.com/lucas-clemente/quic-go"
)

type clientInfo struct {
	uid        uint32
	tid        uint32
	session    quic.Session
	quicStream quic.Stream
}

func newClient(uid, tid uint32) *clientInfo {
	return &clientInfo{
		uid: uid,
		tid: tid,
	}
}
