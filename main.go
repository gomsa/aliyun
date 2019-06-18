package main

import (
	// 公共引入
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	k8s "github.com/micro/kubernetes/go/micro"

	"github.com/gomsa/tools/env"
	"github.com/gomsa/aliyun/hander"
	pb "github.com/gomsa/aliyun/proto/aliyun"
)

func main() {
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
	)
	srv.Init()

	// aliyun服务实现
	pb.RegisterAliyunHandler(srv.Server(), &hander.Aliyun{
		RegionId:     	 env.Getenv("REGION_ID","default"),
		AccessKeyId:     env.Getenv("ACCESS_KEY_ID",""),
		AccessKeySecret: env.Getenv("ACCESS_KEY_SECRET",""),
	})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
