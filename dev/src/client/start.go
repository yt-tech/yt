package client

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
		ss := strings.Split(Input, " ")
		c1 := strings.Contains(ss[0], "cmd")
		c2 := strings.Contains(ss[1], "uid")
		c3 := strings.Contains(ss[2], "tid")
		a := ss[0][4:]
		if c1 && c2 && c3 {
			switch {
			case a == "1":
				fmt.Println("connect")
				data, err := packConnectData()
				if err != nil {
					mlog.Println(err)
					return
				}
				mlog.Println(err)
				quicStream.SetDeadline(time.Now().Add(3e9))
				quicStream.Write(data)
			case a == "2":
				fmt.Println("sub")
				data, err := packSubscribeTopic()
				if err != nil {
					mlog.Println(err)
					return
				}
				quicStream.Write(data)
			case a == "3":
				fmt.Println("hold mic")
				go func() {
					for {
						data, err := packAudioData()
						if err != nil {
							mlog.Println(err)
							return
						}
						mlog.Println(data)
						quicStream.Write(data)
						time.Sleep(12e7)
					}
				}()
			case a == "4":
				fmt.Println("release")
			case a == "5":
				fmt.Println("unsub")
			case a == "6":
				fmt.Println("disconnect")
			}
			continue
		}

		fmt.Println("输入错误！ 请重新输入")
		continue
	}
}
