package manager

import (
	command "yt/ytproto/cmd"
)

//Broadcast ..
func (m *Manager) Broadcast(r *command.Msg, stream command.Manager_BroadcastServer) error {

	mlog.Println(r)
	stream.Send(r)
	return nil
}

//BroadcastRegiste ..
func (m *Manager) BroadcastRegiste(r *command.BroadcastRegiste, stream command.Manager_BroadcastRegisteServer) error {
	return nil
}
