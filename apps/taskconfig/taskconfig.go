package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/logx/logrusx"

	"oh-my-helper-go/apps/taskconfig/internal/config"
	"oh-my-helper-go/apps/taskconfig/internal/server"
	"oh-my-helper-go/apps/taskconfig/internal/svc"
	"oh-my-helper-go/apps/taskconfig/taskconfig"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "apps/taskconfig/etc/taskconfig.yaml", "the config file")

func main() {
	writer := logrusx.NewLogrusWriter(func(logger *logrus.Logger) {
		logger.SetFormatter(&logrus.TextFormatter{})
	})
	logx.SetWriter(writer)

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		taskconfig.RegisterTaskConfigServer(grpcServer, server.NewTaskConfigServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
