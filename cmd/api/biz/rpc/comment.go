package rpc

import (
	"context"
	"douyin/kitex_gen/douyincomment"
	"douyin/kitex_gen/douyincomment/commentservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var commentClient commentservice.Client

func InitComment() {

	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	//c, err := videoservice.NewClient(
	//	consts.VideoServiceName,
	//	client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServiceName}),
	//	client.WithHostPorts(":8888"),
	//)

	c, err := commentservice.NewClient(
		consts.CommentServiceName,
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.CommentServiceName}),
	)

	if err != nil {
		panic(err)
	}
	commentClient = c
	//api发起了rpc调用
}

func CommentAction(ctx context.Context, req *douyincomment.DouyinCommentActionRequest, uid int64) (*douyincomment.DouyinCommentActionResponse, error) {
	resp, err := commentClient.CommentAction(ctx, req, uid)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return resp, nil
}

func CommentList(ctx context.Context, req *douyincomment.DouyinCommentListRequest) (*douyincomment.DouyinCommentListResponse, error) {
	resp, err := commentClient.CommentList(ctx, req)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return resp, nil
}
