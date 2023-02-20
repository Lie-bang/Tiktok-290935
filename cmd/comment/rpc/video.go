package rpc

import (
	"context"
	"douyin/kitex_gen/douyinvideo"
	"douyin/kitex_gen/douyinvideo/videoservice"
	"douyin/pkg/consts"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

//var videoClient videoservice.Client

var videoClient videoservice.Client

func InitVideo() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.CommentServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := videoservice.NewClient(
		consts.VideoServiceName, // DestService
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.CommentServiceName}),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

//func GerUser(ctx context.Context, req *douyinuser.GetUserRequest) (*douyinuser.User, error) {
//	resp, err := userClient.GetUser(ctx, req)
//	if err != nil {
//		return nil, err
//	}
//	if resp.BaseResp.StatusCode != 0 {
//		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
//	}
//	return resp.User, nil
//}

func CommentCountUpdate(ctx context.Context, req *douyinvideo.DouyinCommentCountRequest) error {
	_, err := videoClient.CommentCountUpdate(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
