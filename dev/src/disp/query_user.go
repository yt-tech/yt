package disp

import (
	"github.com/go-redis/redis"

	// pq
	"database/sql"

	jsoniter "github.com/json-iterator/go"
	_ "github.com/lib/pq"
)

type userSignUpInfo struct {
	userAccount string
	password    string
}
type userInfo struct {
	Passwd   string `json:"passwd"`
	UserName string `json:"user_name"`
}

func queryUser(cliRequest userSignUpInfo) (uint8, string) {
	var uInfo userInfo
	uid, err := redisdb.Get(cliRequest.userAccount).Result()
	if err == redis.Nil {
		mlog.Printf("%s is not found\n", uid)
		connStr := "postgres://postgres:postgres@localhost:5433/yt"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			mlog.Fatal(err)
			return 0, "0"
		}

		row := db.QueryRow("SELECT uid,pwd,user_name FROM yt_user WHERE account = $1", cliRequest.userAccount)
		if err = row.Scan(&uid, &uInfo.Passwd, &uInfo.UserName); err != nil {
			mlog.Println(err)
			return 0, "0"
		}
		if cliRequest.password != uInfo.Passwd {
			mlog.Println("pwd is error")
			return 6, "0"
		}
		mlog.Println("add new user to redis", cliRequest.userAccount)
		err = redisdb.Set(cliRequest.userAccount, uid, 0).Err()
		if err != nil {
			mlog.Println(err)
			return 0, "0"
		}
		userstr, err := jsoniter.MarshalToString(&uInfo)
		if err != nil {
			mlog.Println(err)
		}
		mlog.Println(uid, userstr, uInfo)
		err = redisdb.Set(uid, userstr, 0).Err()
		if err != nil {
			mlog.Println(err)
			return 0, "0"
		}
	} else if err != nil {
		mlog.Println(err)
		return 0, "0"
	}
	ubytes, err := redisdb.Get(uid).Bytes()
	if err != nil {
		mlog.Println(err)
	}
	err = jsoniter.Unmarshal(ubytes, &uInfo)
	if err != nil {
		mlog.Println(err)
		return 0, "0"
	}

	return 1, uid
}
