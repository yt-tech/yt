package gateway

import (
	"fmt"
	"net"
)

var ublconn *net.UDPConn

func init() {
	addr, err := net.ResolveUDPAddr("udp", "192.168.1.113:9999")
	if err != nil {
		fmt.Println(err)
	}
	ublconn, err := net.ListenUDP("udp", addr)
	net.DialUDP()
	if err != nil {
		fmt.Println(err)
	}
}
func audioListen() {

	buff := make([]byte, 1024)
	for {
		n, _, err := ublconn.ReadFromUDP(buff)
		if err != nil {
			fmt.Println(err)
		}
	}
}
