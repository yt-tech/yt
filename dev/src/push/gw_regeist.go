package push

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"

	"github.com/smallnest/rpcx/server"
)

var GateWayList sync.Map //key:GatewayID  value:*GateWayInfo

type PushServer struct{}

type GateWayInfo struct {
	TCP net.Conn
	UDP *net.UDPAddr
}

type GateWayRegisteInfo struct {
	ID              uint64
	ReceiveListener string
}
type GateWayReply struct {
	R int8
}

//GatewayRegiste ..
func (g *PushServer) GatewayRegiste(ctx context.Context, gwri *GateWayRegisteInfo, reply *GateWayReply) error {
	rlAddr, err := createGateWayReceiveListener(gwri.ReceiveListener)
	if gwri.ID < 1 {
		return errors.New("GatewayID")
	} else if err != nil {
		return err
	}
	if gwConn, ok := ctx.Value(server.RemoteConnContextKey).(net.Conn); ok {
		gwi := &GateWayInfo{
			TCP: gwConn,
			UDP: rlAddr,
		}
		mlog.Println("request gw id=", gwri.ID)
		if _, IsExist := GateWayList.LoadOrStore(gwri.ID, gwi); IsExist {
			reply.R = 2
			return nil
		}
		reply.R = 1
		return nil
	}
	return errors.New("conn")
}

func createGateWayReceiveListener(remoteAddr string) (*net.UDPAddr, error) {
	addr, err := net.ResolveUDPAddr("udp4", remoteAddr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return addr, nil
}
