package client

import (
	"math/rand"
	"time"
)

//生成随机数
func getRangNumber() uint32 {
	rand.Seed(time.Now().Unix())
	return rand.Uint32()
}
