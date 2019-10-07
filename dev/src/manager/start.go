package manager

import (
	tp "github.com/henrylee2cn/teleport"
)

var srv tp.Peer

//StartManager ..
func StartManager() {
	defer tp.FlushLogger()
	// graceful
	go tp.GraceSignal()

	// server peer
	srv = tp.NewPeer(tp.PeerConfig{
		CountTime:   true,
		ListenPort:  9090,
		PrintDetail: true,
	})

	srv.RouteCall(new(Manager))
	srv.ListenAndServe()
	select {}
}
