package manager

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

var config *toml.Tree

func init() {
	var err error
	config, err = toml.LoadFile("../conf/manager.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
}

func serveAddr() string {
	t := config.Get("manager.serverAddr")
	if t != nil {
		if s, ok := t.(string); ok {
			return s
		}
	}
	mlog.Fatalf("toml set error %v\n", t)
	return ""
}
