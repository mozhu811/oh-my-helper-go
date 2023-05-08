package main

import (
	"flag"
	"fmt"

	"oh-my-helper-go/apps/bilibili/rpc/bilibili"
	"oh-my-helper-go/apps/bilibili/rpc/internal/config"
	"oh-my-helper-go/apps/bilibili/rpc/internal/server"
	"oh-my-helper-go/apps/bilibili/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "apps/bilibili/rpc/etc/bilibili.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		bilibili.RegisterBilibiliServer(grpcServer, server.NewBilibiliServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
