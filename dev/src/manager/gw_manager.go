package manager

import tp "github.com/henrylee2cn/teleport"

func (m *Manager) Broadcastregister(request *string) (uint8, *tp.Rerror) {
	mlog.Println("----->", m.Session().ID())
	gatewayBroadcastAddrs.Store(m.Session().ID(), m.Session())
	return 1, nil
}
