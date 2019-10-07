package gateway

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

var config *toml.Tree

func init() {
	var err error
	config, err = toml.LoadFile("../conf/gateway.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
}
func logConfig() (string, string) {
	// retrieve data directly
	logFilePath := config.Get("log.logFilePath").(string)
	logLevel := config.Get("log.logLevel").(string)
	return logFilePath, logLevel
}
func quicServeAddr() string {
	t := config.Get("gateway.serverAddr")
	if t != nil {
		if s, ok := t.(string); ok {
			return s
		}
	}
	mlog.Fatalf("toml set error %v\n", t)
	return ""
}
func cmdBroadcastAddr() string {
	t := config.Get("gateway.cmdBroadcastListen")
	if t != nil {
		if s, ok := t.(string); ok {
			return s
		}
	}
	mlog.Fatalf("toml set error %v\n", t)
	return ""
}
func audioBroadcastAddr() string {
	t := config.Get("gateway.audioBroadcastListen")
	if t != nil {
		if s, ok := t.(string); ok {
			return s
		}
	}
	mlog.Fatalf("toml set error %v\n", t)
	return ""
}
func managerServerAddr() string {
	t := config.Get("gateway.managerServerAddr")
	if t != nil {
		if s, ok := t.(string); ok {
			return s
		}
	}
	mlog.Fatalf("toml set error %v\n", t)
	return ""
}
