package main

import (
	"douyin/cmd/video/dal"
	"douyin/cmd/video/rpc"
	douyinvideo "douyin/kitex_gen/douyinvideo/videoservice"
	"douyin/pkg/consts"
	"douyin/pkg/minio"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func Init() {
	rpc.Init()
	dal.Init()
	minio.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.VideoServiceAddr)

	Init()
	svr := douyinvideo.NewServer(new(VideoServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServiceName}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
