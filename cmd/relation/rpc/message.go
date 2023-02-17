package rpc

import (
	"context"
	"douyin/kitex_gen/douyinmessage"
	"douyin/kitex_gen/douyinmessage/messageservice"
	"douyin/pkg/consts"
	"douyin/pkg/errno"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var messageClient messageservice.Client

func InitMessage() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.MessageServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)

	c, err := messageservice.NewClient(
		consts.MessageServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithInstanceMW(mw.CommonMiddleware),
		client.WithMiddleware(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.MessageServiceName}),
	)
	if err != nil {
		panic(err)
	}

	messageClient = c
}

func GetFirstMessages(ctx context.Context, req *douyinmessage.GetFirstMessagesRequest) ([]*douyinmessage.Message, error) {
	resp, err := messageClient.GetFirstMessages(ctx, req)
	if err != nil {
		return nil, nil
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMessage)
	}
	return resp.Messages, nil
}
