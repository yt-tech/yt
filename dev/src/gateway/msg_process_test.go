package gateway

import (
	"sync"
	"testing"
)

// func Benchmark_pool(b *testing.B) {

// 	for i := 0; i < b.N; i++ { //use b.N for looping
// 		withPool()
// 	}
// }

// // func init() {
// // 	for k := 0; k < 200; k++ {
// // 		var _ = msgPool.Get().(*msg.Msg)
// // 	}
// // }
// func withPool() {
// 	var p = msgPool.Get().(*msg.Msg)
// 	msgPool.Put(p)
// }

var smap sync.Map

func Benchmark_syncMap(b *testing.B) {
	for k := uint32(0); k < 15000; k++ {
		smap.Store(k, "121231231")
	}
	for i := 0; i < b.N; i++ { //use b.N for looping
		s, _ := smap.Load(500)
		_ = s.(string)
	}
}
