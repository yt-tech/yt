package manager

import (
	"context"
	command "yt/ytproto/cmd"
)

//SubscribeTopicRequest ..
func (m *Manager) SubscribeTopicRequest(ctx context.Context, request *command.SubscribeTopicRequestInfo) (*command.SubscribeTopicResponseInfo, error) {
	mlog.Println("Subscribe Topic Request")

	var result int32
	uid := request.GetUid()
	tid := request.GetTid()

	tlister, isEsixt := m.topics.Load(tid)
	if isEsixt {
		tlist, ok := tlister.(map[uint64]bool)
		if ok {
			tlist[uid] = true
			result = 2
		}
	} else {
		topic := make(map[uint64]bool, 50)
		m.topics.Store(tid, topic)
		result = 1
	}

	SubscribeAck := &command.SubscribeTopicResponseInfo{
		Result: result,
	}

	return SubscribeAck, nil
}
