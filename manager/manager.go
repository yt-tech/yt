package manager

import "sync"

//Manager ..
type Manager struct {
	users  sync.Map
	topics sync.Map
}

var gatewayList sync.Map
