package main

import (
	"flag"
	"fmt"
	"yt/client"
	"yt/gateway"
	"yt/manager"
	"yt/push"
)

func main() {
	var args = flag.String("m", "", "--")
	flag.Parse()
	switch *args {
	case "c":
		fmt.Println("start client")
		client.Start()
	case "g":
		fmt.Println("start gateway")
		gateway.QuicServer()
	case "m":
		fmt.Println("start manager")
		manager.StartManager()
	case "p":
		fmt.Println("start push")
		push.StartPush()
	default:
		panic("-----")
	}
}
