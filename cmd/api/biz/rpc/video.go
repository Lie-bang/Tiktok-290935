package rpc

import (
	"context"
	"douyin/kitex_gen/douyinvideo"
	"douyin/kitex_gen/douyinvideo/videoservice"
	"douyin/pkg/consts"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var videoClient videoservice.Client

func InitVideo() {

	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	//c, err := videoservice.NewClient(
	//	consts.VideoServiceName,
	//	client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServiceName}),
	//	client.WithHostPorts(":8888"),
	//)

	c, err := videoservice.NewClient(
		consts.VideoServiceName,
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServiceName}),
	)

	if err != nil {
		panic(err)
	}
	videoClient = c
	//api发起了rpc调用
}

func FeedVideo(ctx context.Context, req *douyinvideo.DouyinFeedRequest) (*douyinvideo.DouyinFeedResponse, error) {
	//调用RPC Service

	// goto kitex handler
	fmt.Println("int rpc.FeedVideo: userid = ", *req.Token)
	resp, err := videoClient.FeedVideo(ctx, req)
	//此处的返回值实际上还是*[]db_video type

	if err != nil {
		log.Print(err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		log.Printf("resp.StatusCode != 0")
		return nil, nil
	}
	return resp, err
}

func PublishVideo(ctx context.Context, req *douyinvideo.DouyinPublishActionRequest) (*douyinvideo.DouyinPublishActionResponse, error) {
	fmt.Println("get in videoClient.PublishVIdeo")
	resp, err := videoClient.PublishVideo(ctx, req)
	fmt.Println("get out videoClient.PublishVIdeo")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, err
}

func PublishListVideo(ctx context.Context, req *douyinvideo.DouyinPublishListRequest) (*douyinvideo.DouyinPublishListResponse, error) {
	resp, err := videoClient.PublishListVideo(ctx, req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, err
}
