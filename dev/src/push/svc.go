package push

import (
	"net"
	"net/http"

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
	ln, _ := net.Listen("tcp", ":9981")
	go http.Serve(ln, nil)
	s = server.NewServer()
	s.Register(new(PushServer), "")
	go s.Serve("tcp", addr)
	select {}
}
