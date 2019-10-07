package push

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"
)

//Reply ..
type Reply struct {
	R int8
}

//GateWayListener ..
type GateWayListener struct {
	GateWayName string
	Listener    *net.UDPAddr
}
type UserGroupRequestInfo struct {
	GatewayID uint64
	GroupID   uint64
	UserID    uint64
}

var mlog = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

//JoinGroup ..
func (g *PushServer) JoinGroup(ctx context.Context, ur *UserGroupRequestInfo, reply *Reply) error {
	_, isExist := GateWayList.Load(ur.GatewayID) //检查网关注册
	if !isExist {                                //网关没有注册
		reply.R = -1
		return nil
	}
	//网关已经注册
	egws, isExist := GroupList.Load(ur.GroupID)
	if !isExist { //首个进入群组
		gatewaysNew := make(map[uint64]struct{}, 10)
		gatewaysNew[ur.GatewayID] = struct{}{}
		a := &GroupInfo{
			ContainGateWays: gatewaysNew,
		}
		GroupList.Store(ur.GroupID, a)
		reply.R = 0 //当前群组第一个登录
		return nil
	}
	if gw, ok := egws.(*GroupInfo); ok {
		for k := range gw.ContainGateWays {
			if k == ur.GatewayID {
				reply.R = 2
				break
			}
			if conner, isExist := GateWayList.Load(k); isExist {
				conn, ok := conner.(*GateWayInfo)
				if ok {
					s.SendMessage(conn.TCP, "", "", nil, []byte("uid="+strconv.FormatUint(ur.UserID, 10)+" groupid="+strconv.FormatUint(ur.GroupID, 10)+" gatewayid="+strconv.FormatUint(uint64(k), 10)))
				}
			}
		}
		if _, ok := gw.ContainGateWays[ur.GatewayID]; !ok {
			reply.R = 1
			gw.ContainGateWays[ur.GatewayID] = struct{}{}
		}
	}
	return nil
}

// //LeaveGroup ..
// func (u *GateWay) LeaveGroup(ctx context.Context, reply *Reply) error {
// 	return nil
// }

// //RemoveGroup ..
// func (u *GateWay) RemoveGroup(ctx context.Context, reply *string) error {
// 	return nil
// }
