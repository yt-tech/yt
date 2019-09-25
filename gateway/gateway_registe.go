package gateway

import (
	"log"
)

var (
	serviceName = "gateway-01"
	serviceInfo = ServiceInfo{IP: quicaddr}
)

func gatewayRegister() {
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
