package main

import (
	"flag"
	"fmt"

	"oh-my-helper-go/apps/pushconfig/internal/config"
	"oh-my-helper-go/apps/pushconfig/internal/server"
	"oh-my-helper-go/apps/pushconfig/internal/svc"
	"oh-my-helper-go/apps/pushconfig/pushconfig"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "apps/pushconfig/etc/pushconfig.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pushconfig.RegisterPushConfigServer(grpcServer, server.NewPushConfigServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
