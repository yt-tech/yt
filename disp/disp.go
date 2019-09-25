package disp

import (
	"log"
	"os"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var mlog = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

const dispAddr = 9001

//StartDisp ..
func StartDisp() {
	go gatewayWatcher()
	s := g.Server()
	s.SetPort(dispAddr)
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write(getGateway())
	})
	s.Run()
}
