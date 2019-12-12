package gateway

import (
	"github.com/go-redis/redis"
)

func verifyAccount(accounToken string) bool {
	val, err := rdb.Get(accounToken).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		mlog.Println(err)
		return false
	}
	mlog.Println("key", val)
	return true
}
