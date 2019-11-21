package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
)

var gatewayAddr string
var mlog = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
var inDataChannel = make(chan []byte, 100)
var outDataChannel = make(chan []byte, 100)

//Start ..
func Start() {
	var setUID = uint32(4)
	cli := newClient(setUID, 1)
	// gatewayAddr = getDisp()
	mlog.Println(gatewayAddr)
	cli.openQuic()
	// time.Sleep(1e9)
	// logger.Debug("start zap log")
	go func() {
		tike := time.NewTicker(20e9)
		cm := &msg.Msg{
			CmdID: msg.CMDID_Ping,
			Uid:   setUID,
		}
		pingBytes, _ := ggproto.Marshal(cm)
		for {
			select {
			case <-tike.C:
				cli.quicStream.Write(pingBytes)
			}
		}
	}()

	for {
		f := bufio.NewReader(os.Stdin) //读取输入的内容
		fmt.Print("请输入命令->")
		Input, err := f.ReadString('\n') //定义一行输入的内容分隔符。
		if err != nil {
			panic(err)
		}
		if Input[:len(Input)-1] == "exit" {
			fmt.Println("exit")
			break
		}
		//cmd=1 uid=1 tid=1
		c1 := strings.Contains(Input, "cmd")
		a := Input[4:5]
		if c1 {
			switch {
			case a == "1":
				go func() {
					recv, _ := cli.session.AcceptUniStream(context.Background())
					ba := make([]byte, 1024)
					for {
						n, err := recv.Read(ba)
						if err != nil {
							os.Exit(10)
						}
						mlog.Println(ba[:n], err)
					}
				}()
				fmt.Println("connect")
				err = cli.connect()
				if err != nil {
					mlog.Println(err)
				}

			case a == "2":
				fmt.Println("sub")
				err = cli.subscribeTopic()
				if err != nil {
					mlog.Println(err)
				}
			case a == "3":
				fmt.Println("hold mic")
				err = cli.holdMic()
				if err != nil {
					mlog.Println(err)
				}

			case a == "4":
				fmt.Println("release mic")
				err = cli.releaseMic()
				if err != nil {
					mlog.Println(err)
				}
			case a == "5":
				fmt.Println("unsub")
				err = cli.unsubscribeTopic()
				if err != nil {
					mlog.Println(err)
				}
			case a == "6":
				fmt.Println("disconnect")
				err = cli.disconnect()
				if err != nil {
					mlog.Println(err)
				}
			}
			continue
		}

		fmt.Println("输入错误！ 请重新输入")
		continue
	}
}
