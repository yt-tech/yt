package disp

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//StartDisp ..
func StartDisp() {
	s := g.Server()
	s.SetPort(7811)
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("127.0.0.1:7822")
	})
	s.Run()
}
