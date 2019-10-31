package gateway

import (
	"fmt"
	"net"
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
)

var ublconn *net.UDPConn
var pushAddr *net.UDPAddr

func audioListen() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9998")
	if err != nil {
		fmt.Println(err)
	}
	push, err := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
	}
	pushAddr = push
	ublconn, err = net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
	}
	buff := make([]byte, 1024)
	for {
		n, _, err := ublconn.ReadFromUDP(buff)
		if err != nil {
			mlog.Println(err, n)
		}
		mlog.Println(buff[:n])
		var mm = new(msg.Msg)
		err = ggproto.Unmarshal(buff[:n], mm)
		if err != nil {
			mlog.Println(err)
		}
		broadcastAudio(mm)
	}
}
