package gateway

import "time"

var (
	pushServerAddr    = "localhost:8972"
	managerServerAddr = "localhost:8973"
	requestGatewayID  = uint32(2)
	requestUserID     = uint64(2)
)

func Start() {
	// go gatewayRegister()
	time.Sleep(1e9)
	QuicServer()
}
