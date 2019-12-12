package disp

import (
	"log"
	"os"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	jsoniter "github.com/json-iterator/go"
)

var mlog = log.New(os.Stdout, "disp ", log.LstdFlags|log.Lshortfile)

const dispAddr = 9001

type dispResponse struct {
	RemoteInfo  string `json:"remote_info"`
	RsqTime     string `json:"response_time"`
	ReturnCode  uint8  `json:"return_code"`
	GatewayAddr string `json:"gateway_addr"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

//StartDisp ..
func StartDisp() {
	// go gatewayWatcher()
	s := g.Server()
	s.SetPort(dispAddr)
	s.BindHandler("/disp/v1", func(r *ghttp.Request) {
		var responseInfo = dispResponse{
			RemoteInfo:  "d01",
			RsqTime:     time.Now().Format("2006-01-02 15:04:05"),
			GatewayAddr: "123.123.123.123:1231",
			AccessToken: "failed",
		}
		var requestUserInfo = userSignUpInfo{
			userAccount: r.GetString("account"),
			password:    r.GetString("passwd"),
		}
		mlog.Println(requestUserInfo)
		rc := queryUser(requestUserInfo)
		switch rc {
		case 1:
			responseInfo.AccessToken, responseInfo.ExpiresIn, responseInfo.Scope, responseInfo.TokenType = requestOauth2()
			mlog.Println(responseInfo)
		default:
			mlog.Println(responseInfo, rc)
		}
		responseInfo.ReturnCode = rc
		buff, err := jsoniter.Marshal(responseInfo)
		if err != nil {
			mlog.Println(err)
			buff = []byte("server error")
		}
		r.Response.Write(buff)
	})
	s.Run()
}
