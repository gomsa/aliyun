package main

import (
	"github.com/gomsa/tools/config"
)

// Conf 配置
var Conf config.Config

func init() {
	Conf = config.Config{
		Service: "aliyun",
		Version: "latest",
	}
}
