package client

import (
	"github.com/gogf/gf/net/ghttp"
)

func getDisp() string {
	r, e := ghttp.Get("http://127.0.0.1:7811")
	if e != nil {
		panic(e)
	}
	defer r.Close()
	return r.ReadAllString()
}
