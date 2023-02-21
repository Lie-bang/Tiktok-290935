package rpc

import (
	"context"
	"douyin/kitex_gen/douyinfavorite"
	"douyin/kitex_gen/douyinfavorite/favoriteservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var favoriteClient favoriteservice.Client

func InitFavorite() {

	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	//c, err := videoservice.NewClient(
	//	consts.VideoServiceName,
	//	client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServiceName}),
	//	client.WithHostPorts(":8888"),
	//)

	c, err := favoriteservice.NewClient(
		consts.FavoriteServiceName,
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.FavoriteServiceName}),
	)

	if err != nil {
		panic(err)
	}
	favoriteClient = c
	//api发起了rpc调用
}

func FavoriteAction(ctx context.Context, req *douyinfavorite.DouyinFavoriteActionRequest) (*douyinfavorite.DouyinFavoriteActionResponse, error) {
	resp, err := favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return resp, nil
}

func FavoriteList(ctx context.Context, req *douyinfavorite.DouyinFavoriteListRequest) (*douyinfavorite.DouyinFavoriteListResponse, error) {
	resp, err := favoriteClient.FavoriteList(ctx, req)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return resp, nil
}
