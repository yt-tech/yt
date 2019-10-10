package gateway

import "time"

//StartGateway ..
func StartGateway() {
	// go gatewayRegister()
	go audioListen()
	time.Sleep(1e8)
	quicServer()
}
