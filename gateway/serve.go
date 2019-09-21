package gateway

import (
	"fmt"
	"net"

	jsoniter "github.com/json-iterator/go"
)

type LoginInfo struct {
	UID uint64 `json:"uid"`
}

func Clisten() {
	addr, err := net.ResolveUDPAddr("udp", "192.168.1.113:9999")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
	}
	buff := make([]byte, 1024)
	for {
		var l LoginInfo
		n, cAddr, err := conn.ReadFromUDP(buff)
		if err != nil {
			fmt.Println(err)
		}
		jsoniter.Unmarshal(buff[:n], &l)
		fmt.Println(l)
		conn.WriteToUDP(buff[:n], cAddr)
	}
}
