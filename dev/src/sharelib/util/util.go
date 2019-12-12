package util

import (
	"math/rand"
	"time"
)

//GetRangNumber 生成随机数
func GetRangNumber() uint32 {
	rand.Seed(time.Now().Unix())
	return rand.Uint32()
}
