package main

import (
	"flag"
	"fmt"
	"yt/client"
	"yt/disp"
	"yt/gateway"
	"yt/manager"
	"yt/push"
)

func main() {
	var args = flag.String("m", "", "--")
	flag.Parse()
	switch *args {
	case "c": //client
		fmt.Println("start client")
		client.Start()
	case "g": //gateway
		fmt.Println("start gateway")
		gateway.QuicServer()
	case "m": //
		fmt.Println("start manager")
		manager.StartManager()
	case "p": //push
		fmt.Println("start push")
		push.StartPush()
	case "d": //disp
		fmt.Println("start disp")
		disp.StartDisp()
	default:
		panic("-----")
	}
}
