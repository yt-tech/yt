package disp

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type ServiceInfo struct {
	IP string
}
type Master struct {
	Path   string
	Nodes  map[string]*Node
	Client *clientv3.Client
}

//node is a client
type Node struct {
	State bool
	Key   string
	Info  ServiceInfo
}

func NewMaster(endpoints []string, watchPath string) (*Master, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second,
	})

	if err != nil {
		mlog.Fatal(err)
		return nil, err
	}

	master := &Master{
		Path:   watchPath,
		Nodes:  make(map[string]*Node),
		Client: cli,
	}

	go master.WatchGateway()
	return master, err
}

func (m *Master) AddNode(key string, info *ServiceInfo) {
	node := &Node{
		State: true,
		Key:   key,
		Info:  *info,
	}

	m.Nodes[node.Key] = node
}

func GetServiceInfo(ev *clientv3.Event) *ServiceInfo {
	info := &ServiceInfo{}
	err := json.Unmarshal([]byte(ev.Kv.Value), info)
	if err != nil {
		log.Println(err)
	}
	return info
}

//WatchGateway ..
func (m *Master) WatchGateway() {
	rch := m.Client.Watch(context.Background(), m.Path, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				mlog.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				info := GetServiceInfo(ev)
				m.AddNode(string(ev.Kv.Key), info)
			case clientv3.EventTypeDelete:
				mlog.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				delete(m.Nodes, string(ev.Kv.Key))
			}
		}
	}
}
