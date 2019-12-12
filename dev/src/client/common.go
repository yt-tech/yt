package client

import (
	"errors"
	"math/rand"
	"time"
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
)

var errLen error = errors.New("buff too long")

//生成随机数
func getRangNumber() uint32 {
	rand.Seed(time.Now().Unix())
	return rand.Uint32()
}
func newBaseToken() *baseToken {
	return &baseToken{complete: make(chan struct{})}
}

func cmdMarsal(cmdMsg *msg.Msg) ([]byte, error) {
	buff, err := ggproto.Marshal(cmdMsg)
	if err != nil {
		return nil, err
	}
	if len(buff) > 512 {
		mlog.Println("len")
		return nil, errLen
	}
	return buff, nil
}
