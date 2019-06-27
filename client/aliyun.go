package client

import (
	"os"

	"github.com/gomsa/tools/k8s/client"

	aliyunPB "github.com/gomsa/aliyun/proto/aliyun"
)

var (
	// Aliyun 用户客户端
	Aliyun aliyunPB.AliyunClient
)

func init() {
	aliyunSrvName := os.Getenv("ALIYUN_NAME")

	Aliyun = aliyunPB.NewAliyunClient(aliyunSrvName, client.DefaultClient)
}
