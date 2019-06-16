package main

import (
	// 公共引入
	"github.com/micro/go-log"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"


	pb "github.com/gomsa/aliyun/proto/aliyun"

)

func main() {
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
	)
	srv.Init()

	// 用户服务实现
	pb.RegisterAliyunHandler(srv.Server(), &hander.Aliyun{
		AccessKeyId:ACCESS_KEY_ID,
		AccessKeySecret:ACCESS_KEY_SECRET,
	})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
