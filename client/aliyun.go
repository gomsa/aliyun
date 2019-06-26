package client

import (
	"context"
	"os"

	"github.com/gomsa/tools/config"
	"github.com/gomsa/tools/k8s/client"

	aliyunPB "github.com/gomsa/aliyun/proto/aliyun"
)

var (
	// Aliyun 用户客户端
	Aliyun aliyunPB.AliyunsClient
)

func init() {
	aliyunSrvName := os.Getenv("ALIYUN_NAME")

	Aliyun = aliyunPB.NewAliyunsClient(aliyunSrvName, client.DefaultClient)
}