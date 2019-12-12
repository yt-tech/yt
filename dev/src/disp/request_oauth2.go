package disp

import (
	"github.com/gogf/gf/net/ghttp"
	jsoniter "github.com/json-iterator/go"
)

// {"access_token":"WCN_Z-XJO4QCULXJ_GWDNW","expires_in":7200,"scope":"read","token_type":"Bearer"}
type tokenCode struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func requestOauth2() (string, int64, string, string) {
	response, err := ghttp.Get("http://localhost:9096/token?grant_type=client_credentials&client_id=yt123456&client_secret=123456&scope=read")
	if err != nil {
		return "auth error", 0, "", ""
	}
	var token tokenCode
	authCodeToken := response.ReadAll()
	err = jsoniter.Unmarshal(authCodeToken, &token)
	if err != nil {
		mlog.Println(err)
		return "auth error", 0, "", ""
	}
	return token.AccessToken, token.ExpiresIn, token.Scope, token.TokenType
}
