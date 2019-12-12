package gateway

import (
	"github.com/go-redis/redis"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "", // no password set
	DB:       15, // use default DB
})
