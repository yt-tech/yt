package manager

import (
	"database/sql"
	"strconv"
	"yt/sharelib/redisinfo"

	jsoniter "github.com/json-iterator/go"

	// pq
	_ "github.com/lib/pq"
)

func queryUserInfoFromPQDB(uid uint32) (*redisinfo.UserRedisInfo, error) {
	var user redisinfo.UserRedisInfo
	connStr := "postgres://postgres:postgres@localhost:5433/yt"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		mlog.Fatal(err)
	}

	row := db.QueryRow("SELECT pwd,user_name FROM yt_user WHERE uid = $1", uid)
	if err = row.Scan(&user.Passwd, &user.UserName); err != nil {
		mlog.Println(err)
		return nil, err
	}
	return &user, nil
}

func queryUserInfoFromRedis(uid uint32) (*redisinfo.UserRedisInfo, error) {
	uidstr := strconv.FormatInt(int64(uid), 10)
	var user redisinfo.UserRedisInfo
	userBytes, err := redisdb.Get(uidstr).Bytes()
	if err != nil {
		mlog.Println(err)
		return nil, err
	}
	if err = jsoniter.Unmarshal(userBytes, &user); err != nil {
		mlog.Println(err)
		return nil, err
	}
	return &user, nil
}
func queryTopicInfoFromRedis(tid uint32) (*redisinfo.TopicRedisInfo, error) {
	tidstr := strconv.FormatInt(int64(tid), 10)
	var topic redisinfo.TopicRedisInfo
	topicBytes, err := redisdb.Get("topic_" + tidstr).Bytes()
	if err != nil {
		return nil, err
	}
	if err = jsoniter.Unmarshal(topicBytes, &topic); err != nil {
		return nil, err
	}
	return &topic, nil
}

func updateNewTopicInfoToRedis(tid uint32, topic *redisinfo.TopicRedisInfo) error {
	tidstr := strconv.FormatInt(int64(tid), 10)
	topicBytes, err := jsoniter.Marshal(topic)
	if err != nil {
		mlog.Println(err)
		return err
	}
	if err := redisdb.Set("topic_"+tidstr, topicBytes, 0).Err(); err != nil {
		mlog.Println(err)
		return err
	}
	return nil
}

func queryUserOfTopicInfoFromRedis(uid, tid uint32) bool {
	tidstr := strconv.FormatInt(int64(tid), 10)
	uidstr := strconv.FormatInt(int64(uid), 10)
	isExist := redisdb.HExists("tid_"+tidstr, uidstr).Val()
	return isExist
}

func queryTopicInfoFromPQDB(tid uint32) (*redisinfo.TopicRedisInfo, error) {
	var topic redisinfo.TopicRedisInfo
	connStr := "postgres://postgres:postgres@localhost:5433/yt"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		mlog.Fatal(err)
	}

	row := db.QueryRow("SELECT topic_name FROM yt_topic WHERE tid = $1", tid)
	err = row.Scan(&topic.TopicName)
	if err != nil {
		mlog.Println(err)
		return nil, err
	}
	return &topic, nil
}
func queryUsersOfTopicInfoFromPQDB(tid uint32) map[uint32]uint32 {
	var usersOfTopic = make(map[uint32]uint32, 20)
	connStr := "postgres://postgres:postgres@localhost:5433/yt"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		mlog.Fatal(err)
	}

	row := db.QueryRow("SELECT * FROM yt_user_and_topic WHERE tid = $1", tid)
	if err = row.Scan(); err != nil {
		mlog.Println(err)
		return usersOfTopic
	}
	return usersOfTopic
}

func updateUsersOfTopicInfoToRedis(tid uint32) error {
	tidstr := strconv.FormatInt(int64(tid), 10)
	connStr := "postgres://postgres:postgres@localhost:5433/yt"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		mlog.Fatal(err)
		return err
	}
	var uid string
	var userName string
	rows, err := db.Query("select tu.uid,u.user_name from yt_user_and_topic tu left join yt_user u on u.uid=tu.uid  where tid=$1;", tid)
	if err != nil {
		mlog.Println(err)
		return err
	}
	for rows.Next() {
		if err = rows.Scan(&uid, &userName); err != nil {
			mlog.Println(err)
		}
		redisdb.HSet("tid_"+tidstr, uid, userName)
	}
	return nil
}
