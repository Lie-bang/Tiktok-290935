package main

import (
	"douyin/cmd/favorite/dal"
	"douyin/cmd/favorite/rpc"
	"douyin/kitex_gen/douyinfavorite/favoriteservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func Init() {
	rpc.Init()
	dal.Init()
}

func main() {
	//svr := douyinfavorite.NewServer(new(FavoriteServiceImpl))
	//
	//err := svr.Run()
	//
	//if err != nil {
	//	log.Println(err.Error())
	//}

	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.FavoriteServiceAddr)

	Init()
	svr := favoriteservice.NewServer(new(FavoriteServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.FavoriteServiceName}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
