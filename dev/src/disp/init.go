package disp

import (
	"github.com/go-redis/redis"
)

var redisdb = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "", // no password set
	DB:       15, // use default DB
})
