package gateway

import (
	"testing"
	"yt/ytproto/msg"
)

func Benchmark_pool(b *testing.B) {

	for i := 0; i < b.N; i++ { //use b.N for looping
		withPool()
	}
}
func init() {
	for k := 0; k < 200; k++ {
		var _ = msgPool.Get().(*msg.Msg)
	}
}
func Benchmark_no_pool(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		withoutPool()
	}
}

func withPool() {
	var p = msgPool.Get().(*msg.Msg)
	msgPool.Put(p)
}

func withoutPool() {
	var op = &msg.Msg{
		Mid: 0,
		Request: &msg.Request{
			Connect:   new(msg.ConnectRequestInfo),
			Subscribe: new(msg.SubscribeTopicRequestInfo),
		},
		Response: &msg.Response{
			ConnectAck:   new(msg.ConnectResponseInfo),
			SubscribeAck: new(msg.SubscribeTopicResponseInfo),
		},
		AudioData: &msg.AudioData{},
	}
	var _ = op
}
