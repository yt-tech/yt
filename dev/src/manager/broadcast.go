package manager

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

func broadcast(rpcsess tp.Session, bmsg *msg.Msg) {
	mlog.Println("broadcast .....")
	err := rpcsess.Push("/broadcast/push", bmsg).ToError()
	if err != nil {
		mlog.Println(err)
	}
}
