package manager

import (
	managerproto "yt/rpcproto"
)

//Broadcast ..
func (m *Manager) Broadcast(r *managerproto.BroadcastRegiste, stream managerproto.Data_BroadcastServer) error {
	mlog.Println(r.GetId())
	var sn = &managerproto.BroadcastInfo{
		Id: 1,
	}
	stream.Send(sn)
	return nil
}
