package manager

import (
	"yt/ytproto/msg"

	tp "github.com/henrylee2cn/teleport"
)

func broadcast(tpsess tp.Session, bmsg *msg.Msg) {
	mlog.Println("broadcast .....")
	err := tpsess.Push("/broadcast/push", bmsg).ToError()
	if err != nil {
		mlog.Println(err)
	}
}
func systembroadcast(tpsess tp.Session, bmsg *msg.Msg) {
	mlog.Println("broadcast .....")
	err := tpsess.Push("/broadcast/systembroadcast", bmsg).ToError()
	if err != nil {
		mlog.Println(err)
	}
}
func releaseNotice() {

}
