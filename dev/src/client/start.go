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

var mlog = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

// var inDataChannel = make(chan []byte, 100)
// var outDataChannel = make(chan []byte, 100)

//Start ..
func Start() {
	var setUID = uint32(1)
	var userIn = requestDisp{
		user:   "user1",
		passwd: "abc",
	}
	gatewayAddr := userIn.getDisp()
	mlog.Println(gatewayAddr)
	NewClient(setUID, 1, gatewayAddr.AccessToken)
	cli := client
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
	go client.cmdRecieve()
	go client.handle()
	go client.cmdSend()
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
			switch a {
			case "1":
				createLocalRevieveBroadcast()
				mlog.Println("SignIn", SignIn())
			case "2":
				mlog.Println("SubscribeTopic", SubscribeTopic())
			case "3":
				mlog.Println("holdMic", HoldMic())
			case "4":
				mlog.Println("releaseMic", ReleaseMic())
			case "5":
				mlog.Println("unsubsricbeTopic", UnsubscribeTopic())
			case "6":
				mlog.Println("signOut", SignOut())
			}
			continue
		}

		fmt.Println("输入错误！ 请重新输入")
		continue
	}
}

func createLocalRevieveBroadcast() {
	var ms = msg.Msg{}
	go func() {
		recv, err := client.quicSession.AcceptUniStream(context.Background())
		if err != nil {
			mlog.Println(err)
		}
		ba := make([]byte, 1024)
		for {
			n, err := recv.Read(ba)
			if err != nil {
				mlog.Println(err)
				continue
			}

			err = ggproto.Unmarshal(ba[:n], &ms)
			if err != nil {
				mlog.Println(err)
				continue
			}
			switch ms.GetCmdID() {
			case msg.CMDID_SignInAck:
				mlog.Println(ms)
			case msg.CMDID_SignOutAck:
				mlog.Println(ms)
			case msg.CMDID_SubscribeTopicAck:
				mlog.Println(ms)
			case msg.CMDID_UnsubscribeTopicAck:
				mlog.Println(ms)
			case msg.CMDID_HoldMicAck:
				mlog.Println(ms)
			case msg.CMDID_ReleaseMicAck:
				mlog.Println(ms)
			case msg.CMDID_RemoveMicAck:
				mlog.Println(ms)
			case msg.CMDID_Audio:
				mlog.Println(ms)
			}
		}
	}()
}
