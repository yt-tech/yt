package disp

import (
	"sync"
	"time"
)

var gwList = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string, 10)}

const etcdAddr = "http://127.0.0.1:2379"

func getGateway() string {
	var gwaddr string
	gwList.RLock()
	for _, v := range gwList.m {
		gwaddr = v
		gwList.RUnlock()
		return gwaddr
	}
	return ""
}

func gatewayWatcher() {
	m, err := NewMaster([]string{etcdAddr}, "gateway_list/")

	if err != nil {
		mlog.Fatal(err)
	}

	for {
		mlog.Printf("gateway number=%d\n", len(m.Nodes))
		gwList.Lock()
		for k, v := range m.Nodes {
			mlog.Printf("gateway name=%s, gateway address=%s\n", k, v.Info.IP)
			gwList.m[k] = v.Info.IP
		}
		gwList.Unlock()
		time.Sleep(time.Second * 5)
	}
}

func loginQurey(uid uint64) bool {
	// mlog.Printf("select data %d\n", uid)
	// row := dbConn.QueryRow("SELECT uid FROM yt_user WHERE uid = $1", uid)
	// var i uint64
	// err := row.Scan(&i)
	// if err != nil {
	// 	mlog.Println(err)
	// 	return false
	// }
	return true
}
