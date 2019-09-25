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
	gatewayAddr = getDisp()
	mlog.Println(gatewayAddr)
	go openQuic()
	userConnect()
	time.Sleep(2e9)
	userJoinGroup()
	time.Sleep(2e9)
	userHoldMic()
	select {}
}
