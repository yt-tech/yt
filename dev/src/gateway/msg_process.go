package gateway

import (
	"context"
	"net"
	"time"
	"yt/sharelib/timewheel"
	"yt/ytproto/msg"

	ggproto "github.com/gogo/protobuf/proto"
	tp "github.com/henrylee2cn/teleport"
	"github.com/lucas-clemente/quic-go"
)

type ytClientInfo struct {
	uid                 uint32
	currentTopic        uint32
	topicPushServerAddr *net.UDPAddr
	quicSession         quic.Session
	commandStream       quic.Stream
	broadcastStream     quic.SendStream
	tpSession           tp.Session
}

func process(quicSession quic.Session, tpSession tp.Session) {
	if tpSession == nil || quicSession == nil {
		mlog.Fatalln("rpcsess or sess is nil ")
	}
	var y = new(ytClientInfo)
	y.quicSession = quicSession
	y.tpSession = tpSession

	var err error
	y.commandStream, err = y.quicSession.AcceptStream(context.Background())
	if err != nil {
		y.closenet()
		mlog.Println(err)
		return
	}
	readBuff := make([]byte, 1024)
	var message = &msg.Msg{}
	tw := timewheel.NewTimeWheel(time.Second, 120)

	var refreshAudioFlag bool
	var timewheelStatus bool
	for {
		n, err := y.commandStream.Read(readBuff)
		if err != nil {
			y.closenet()
			mlog.Println(err, y.quicSession.RemoteAddr().String(), y.quicSession.LocalAddr())
			return
		}
		// mlog.Println(readBuff[:n])
		err = ggproto.Unmarshal(readBuff[:n], message)
		if err != nil {
			mlog.Println(err)
			continue
		}
		// mlog.Println("|||||||", message)
		y.uid = message.GetUid()
		switch message.GetCmdID() {
		case msg.CMDID_SignIn:
			y.newConnect(message)
		case msg.CMDID_SubscribeTopic:
			y.newSubscribeTopic(message)
		case msg.CMDID_UnsubscribeTopic:
			if err = y.newUnsubscribeTopic(message); err == nil {
				mlog.Printf("error=%v uid=%d tid=%d unSubscribeTopic \n", err, y.uid, message.GetTid())
			}
			mlog.Printf("uid=%d tid=%d unSubscribeTopic success\n", y.uid, message.GetTid())
		case msg.CMDID_HoldMic:
			if err = y.newHoldMic(message); err != nil {
				mlog.Printf("error=%v uid=%d tid=%d holdMic\n", err, y.uid, message.GetTid())
			}
			tw.Start()
			timewheelStatus = true
			//添加定时任务
			//参数：interval 时间间隔
			//参数：times 执行次数 -1 表示周期任务 >0 执行指定次数
			//参数：key 任务唯一标识符 用户更新任务和删除任务
			//参数：taskData 回调函数参数
			//参数：job 回调函数
			tw.AddTask(60*time.Second, 1, "micLong", timewheel.TaskData{"name": "hold mic time long"},
				func(params timewheel.TaskData) {
					mlog.Println(params["name"], "mic setting 60s  removeMic")
					if timewheelStatus {
						y.newRemoveMic()
						tw.Stop()
					}
				})
			tw.AddTask(3*time.Second, -1, "audioData", timewheel.TaskData{"name": refreshAudioFlag},
				func(params timewheel.TaskData) {
					if !refreshAudioFlag {
						mlog.Println(params["name"], "long time had not receive audio data removeMic")
						if timewheelStatus == true {
							timewheelStatus = false
							y.newRemoveMic()
							tw.Stop()
						}
					}
					refreshAudioFlag = false
				})

			mlog.Printf("uid=%d tid=%d holdMic success\n", y.uid, message.GetTid())
		case msg.CMDID_ReleaseMic:
			//轮盘停止
			if timewheelStatus == true {
				y.newRemoveMic()
				tw.Stop()
			}
			if err = y.newReleaseMic(message); err != nil {
				mlog.Printf("error=%v uid=%d tid=%d releaseMic\n", err, y.uid, message.GetTid())
			}
			mlog.Printf("uid=%d tid=%d releaseMic success\n", y.uid, message.GetTid())
		case msg.CMDID_SignOutAck:
		case msg.CMDID_Audio:
			mlog.Println("audio data")
			refreshAudioFlag = true
			y.audioReceive(message)
		case msg.CMDID_Ping:
			// mlog.Println(message)
		default:
			mlog.Println("--------", message)
		}
	}

}
