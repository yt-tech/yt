package client

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

var config *toml.Tree

func init() {
	var err error
	config, err = toml.LoadFile("../conf/client.toml")
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
	return config.Get("client.quicServeAddr").(string)
}
