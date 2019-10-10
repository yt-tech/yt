package gateway

import "time"

//StartGateway ..
func StartGateway() {
	// go gatewayRegister()
	go broadcastListen()
	time.Sleep(1e8)
	quicServer()
}
