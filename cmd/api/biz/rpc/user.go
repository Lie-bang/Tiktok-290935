package rpc

import (
	"context"
	"douyin/kitex_gen/douyinuser"
	"douyin/kitex_gen/douyinuser/userservice"
	"douyin/pkg/consts"
	"douyin/pkg/errno"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	provider.NewOpenTelemetryProvider(
		provider.WithInsecure(),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithServiceName(consts.ApiServiceName),
	)

	c, err := userservice.NewClient(
		consts.UserServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.ClientMiddleware),
		client.WithInstanceMW(mw.CommonMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}),
	)

	if err != nil {
		panic(err)
	}

	userClient = c
}

func CreateUser(ctx context.Context, req *douyinuser.CreateUserRequest) error {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMessage)
	}
	return nil
}

func CheckUser(ctx context.Context, req *douyinuser.CheckUserRequest) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMessage)
	}
	return resp.UserId, nil
}

func GerUser(ctx context.Context, req *douyinuser.GetUserRequest) (*douyinuser.User, error) {
	resp, err := userClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMessage)
	}
	return resp.User, nil
}
