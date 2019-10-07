package gateway

import (
	"log"

	tp "github.com/henrylee2cn/teleport"
)

func gatewayRegister() {
	var (
		serviceName = "gateway-01"
		serviceInfo = ServiceInfo{IP: quicServeAddr()}
	)

	s, err := NewService(serviceName, serviceInfo,
		[]string{
			"http://127.0.0.1:2379",
		})

	if err != nil {
		log.Fatal(err)
	}

	mlog.Printf("name:%s, ip:%s\n", s.Name, s.Info.IP)
	s.Start()
	// go func() {
	// 	time.Sleep(time.Second * 20)
	// 	s.Stop()
	// }()
}

func broadcastRegister(rpcsess tp.Session) error {
	var result int
	rerr := rpcsess.Call("/manager/broadcastregister",
		"1",
		&result,
	).Rerror()
	return rerr.ToError()
}
