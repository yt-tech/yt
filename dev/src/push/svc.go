package push

import (
	"net"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = "localhost:8972"
)

type pushServerInfo struct {
	ID         uint32
	serverAddr string
}

const pushServerID = uint32(1)

var s *server.Server

//StartPush ..
func StartPush() {
	addr, err := net.ResolveUDPAddr("udp", ":9999")
	if err != nil {
		mlog.Println(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		mlog.Println(err)
	}
	buff := make([]byte, 1024)
	for {
		n, raddr, _ := conn.ReadFromUDP(buff)
		_, err = conn.WriteToUDP(buff[:n], raddr)
		mlog.Println(buff[:n], err)
	}
}
