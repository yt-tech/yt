package manager

import (
	"database/sql"
	"fmt"
	"log"
	"yt/ytproto/msg"

	"github.com/go-redis/redis"

	tp "github.com/henrylee2cn/teleport"
)

//Subscribetopic ..
func (m *Manager) Subscribetopic(ytmsg *msg.Msg) (result int32, terr *tp.Rerror) {

	gwID := m.Session().ID()
	mlog.Println("Subscribe Topic Request ", gwID)
	gwSession, rerr := getGWSession(gwID)
	mlog.Println("||||||||||||++++++++++")
	if rerr != nil {
		return 100, rerr
	}
	uid := ytmsg.GetUid()
	tid := ytmsg.GetTid()

	var topic *topicInfo
	topicer, isEsixt := topics.Load(tid)

	//manager 无此topic
	if !isEsixt {
		mlog.Println("the topic had not cache")
		allow := topicAllow(tid)
		if !allow {
			mlog.Println("can not create the topic")
			return 101, nil
		}
		userAllow := queryUserOfTopicInfoFromRedis(uid, tid)
		if !userAllow {
			mlog.Println("the user has not auth for the topic")
			return 102, nil
		}
		newCreateTopic(tid, uid, gwID, gwSession)
		return 1, nil
	}
	topic, ok := topicer.(*topicInfo)
	if !ok {
		return 0, tp.NewRerror(11, "断言失败", "")
	}
	userAllow := queryUserOfTopicInfoFromRedis(uid, tid)
	if !userAllow {
		return 103, nil
	}
	topic.addrNewMember(tid, uid, gwID, gwSession, ytmsg)
	return 2, nil
}
func initUserInfo() *userInfo {
	return &userInfo{}
}
func newCreateTopic(tid, uid uint32, gwID string, gwSession tp.Session) {
	topic := &topicInfo{
		users:    make(map[uint32]*userInfo, 20),
		gateways: make(map[string]tp.Session, 5),
	}
	topic.gateways[gwID] = gwSession
	topic.users[uid] = initUserInfo()
	topics.Store(tid, topic)
}
func (t *topicInfo) addrNewMember(tid, uid uint32, gwID string, gwSession tp.Session, ytmsg *msg.Msg) {
	t.Lock()
	t.users[uid] = initUserInfo()
	t.gateways[gwID] = gwSession
	cmdBroadcast(t.gateways, ytmsg)
	t.Unlock()
	topics.Store(tid, t)
}
func getGWSession(gwid string) (tp.Session, *tp.Rerror) {
	gwSessioner, isExist := gatewayBroadcastAddrs.Load(gwid)
	if !isExist {
		return nil, tp.NewRerror(12, "网关不存在", "")
	}
	gwSession, ok := gwSessioner.(tp.Session)
	if !ok {
		return nil, tp.NewRerror(11, "断言失败", "")
	}
	return gwSession, nil
}

//command broadcast
func cmdBroadcast(tg map[string]tp.Session, ytmsg *msg.Msg) {
	for k, v := range tg {
		mlog.Println(k, v)
		broadcast(v, ytmsg)
	}
}

func queryDB() bool {
	connStr := "port=5433 dbname=yt user=postgres password=postgres "
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	var a, b, c int
	rows, err := db.Query("SELECT * FROM yt_user_and_topic WHERE id = $1 and uid=$2", 1, 2)
	fmt.Println(rows, err)
	for rows.Next() {
		rows.Scan(&a, &b, &c)
		fmt.Println(a, b, c)
	}
	return true
}
func topicAllow(tid uint32) bool {
	topicFromDB, err := queryTopicInfoFromRedis(tid)
	if err == redis.Nil {
		mlog.Println("redis has not the topic")
		topicFromDB, err = queryTopicInfoFromPQDB(tid)
		if err != nil {
			mlog.Println(err)
			return false
		}
		err = updateNewTopicInfoToRedis(tid, topicFromDB)
		if err != nil {
			mlog.Println(err)
			return false
		}
		err = updateUsersOfTopicInfoToRedis(tid)
		if err != nil {
			mlog.Println(err)
			return false
		}
	} else if err != nil {
		mlog.Println(err)
		return false
	}
	return true
}
