package hander

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/micro/go-micro/util/log"

	pb "github.com/gomsa/aliyun/proto/aliyun"
)

// Aliyun 授权服务处理
type Aliyun struct {
	RegionId        string
	AccessKeyId     string
	AccessKeySecret string
}

// ProcessCommonRequest 处理公共请求
func (srv *Aliyun) ProcessCommonRequest(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 创建连接
	client, err := sdk.NewClientWithAccessKey(srv.RegionId, srv.AccessKeyId, srv.AccessKeySecret)
	if err != nil {
		log.Log(err)
		return err
	}
	// 配置参数
	request := requests.NewCommonRequest()
	request.Method = req.Method
	request.Scheme = req.Scheme
	request.Domain = req.Domain
	request.Version = req.Version
	request.ApiName = req.ApiName
	request.QueryParams = req.QueryParams

	// 请求
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		log.Log(err)
		return err
	}
	// 返回请求数据
	res.Content = response.GetHttpContentString()
	return nil
}
