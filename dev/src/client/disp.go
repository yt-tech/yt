package client

import (
	"github.com/gogf/gf/net/ghttp"
	jsoniter "github.com/json-iterator/go"
)

type dispResponse struct {
	UserID      string `json:"user_id"`
	RemoteInfo  string `json:"remote_info"`
	RsqTime     string `json:"response_time"`
	ReturnCode  uint8  `json:"return_code"`
	GatewayAddr string `json:"gateway_addr"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}
type requestDisp struct {
	user   string
	passwd string
}

//http://localhost:9001/disp/v1?account=user1&passwd=abc
func (r requestDisp) getDisp() *dispResponse {
	g, e := ghttp.Get("http://localhost:9001/disp/v1?account=" + r.user + "&passwd=" + r.passwd)
	if e != nil {
		panic(e)
	}
	buff := g.ReadAll()
	g.Close()
	var receive dispResponse
	err := jsoniter.Unmarshal(buff, &receive)
	if err != nil {
		return nil
	}
	return &receive
}
