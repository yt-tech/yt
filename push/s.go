package push

import (
	"fmt"
	"log"
	"time"
)

func main() {

	serviceName := "s-test"
	serviceInfo := ServiceInfo{IP: "192.168.1.26"}

	s, err := NewService(serviceName, serviceInfo, []string{
		"http://127.0.0.1:2379",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name:%s, ip:%s\n", s.Name, s.Info.IP)

	go func() {
		time.Sleep(time.Second * 20)
		s.Stop()
	}()

	s.Start()
}
