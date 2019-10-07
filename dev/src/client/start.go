package client

import (
	"log"
	"os"
	"time"
)

var gatewayAddr string
var mlog = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
var inDataChannel = make(chan []byte, 100)
var outDataChannel = make(chan []byte, 100)

//Start ..
func Start() {
	// gatewayAddr = getDisp()
	mlog.Println(gatewayAddr)
	openQuic()
	// time.Sleep(1e9)
	// logger.Debug("sdfsdfsdf testset")
	time.Sleep(120e9)
	// userHoldMic()
	// select {}
}
