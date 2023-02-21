package rpc

import (
	"context"
	"douyin/kitex_gen/douyinfavorite"
	"douyin/kitex_gen/douyinfavorite/favoriteservice"
	"douyin/pkg/consts"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

//var videoClient videoservice.Client

var favoriteClient favoriteservice.Client

func InitFavorite() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.VideoServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := favoriteservice.NewClient(
		consts.FavoriteServiceName, // DestService
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServiceName}),
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

func FavoriteJudge(ctx context.Context, req *douyinfavorite.DouyinFavoriteJudgeRequest) (*douyinfavorite.DouyinFavoriteJudgeResponse, error) {
	resp, err := favoriteClient.FavoriteJudge(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
