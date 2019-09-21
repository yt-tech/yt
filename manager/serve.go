package manager

import "github.com/smallnest/rpcx/server"

var addr = "localhost:8973"

var s *server.Server

//StartManager ..
func StartManager() {
	connPG()
	s = server.NewServer()
	s.Register(new(Manager), "")
	s.Serve("tcp", addr)
}
