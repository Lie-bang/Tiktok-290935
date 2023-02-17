package rpc

import (
	"context"
	"douyin/kitex_gen/douyinrelation"
	"douyin/kitex_gen/douyinrelation/relationservice"
	"douyin/pkg/consts"
	"douyin/pkg/errno"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func InitRelation() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.RelationServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)

	c, err := relationservice.NewClient(
		consts.RelationServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithInstanceMW(mw.CommonMiddleware),
		client.WithMiddleware(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.RelationServiceName}),
	)
	if err != nil {
		panic(err)
	}

	relationClient = c
}

func Action(ctx context.Context, req *douyinrelation.ActionRequest) error {
	resp, err := relationClient.Action(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMessage)
	}
	return nil
}

func FollowList(ctx context.Context, req *douyinrelation.FollowListRequest) ([]*douyinrelation.User, error) {
	resp, err := relationClient.FollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMessage)
	}
	return resp.UserList, nil
}

func FollowerList(ctx context.Context, req *douyinrelation.FollowerListRequest) ([]*douyinrelation.User, error) {
	resp, err := relationClient.FollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMessage)
	}
	return resp.UserList, nil
}

func FriendList(ctx context.Context, req *douyinrelation.FriendListRequest) ([]*douyinrelation.FriendUser, error) {
	resp, err := relationClient.FriendList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMessage)
	}
	return resp.UserList, nil
}
