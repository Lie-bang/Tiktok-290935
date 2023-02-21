package rpc

import (
	"context"
	"douyin/kitex_gen/douyinvideo"
	"douyin/kitex_gen/douyinvideo/videoservice"
	"douyin/pkg/consts"
	"douyin/pkg/errno"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func initvideo() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.UserServiceName),
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
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func WorkAndFavoriteCount(ctx context.Context, req *douyinvideo.Douyin_Work_And_Favorite_CountRequest) (int64, int64, error) {
	resp, err := videoClient.WorkAndFavoriteCount(ctx, req)
	if err != nil {
		return 0, 0, err
	}
	if resp.StatusCode != 0 {
		return 0, 0, errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return resp.WorkCount, resp.FavCount, nil
}
