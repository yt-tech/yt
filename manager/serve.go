package manager

import (
	"log"
	"net"
	managerproto "yt/rpcproto"

	"github.com/smallnest/rpcx/server"
	"google.golang.org/grpc"
)

var addr = "localhost:8973"

var s *server.Server

//StartManager ..
func StartManager() {
	connPG()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	managerproto.RegisterDataServer(s, &Manager{})
	s.Serve(lis)
}
