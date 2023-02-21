package main

import (
	"douyin/cmd/relation/dal"
	"douyin/cmd/relation/rpc"
	douyinrelation "douyin/kitex_gen/douyinrelation/relationservice"
	"douyin/pkg/consts"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

func Init() {
	rpc.Init()
	dal.Init()
	klog.SetLevel(klog.LevelInfo)
	klog.SetLogger(logrus.NewLogger())
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.RelationServiceAddr)

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.RelationServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)

	Init()
	svr := douyinrelation.NewServer(new(RelationServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithMuxTransport(),
		server.WithLimit(&limit.Option{
			MaxConnections: 1000,
			MaxQPS:         100,
		}),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.RelationServiceName}),
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
