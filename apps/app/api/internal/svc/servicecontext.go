package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"oh-my-helper-go/apps/app/api/internal/config"
	"oh-my-helper-go/apps/bilibili/rpc/bilibiliclient"
	"oh-my-helper-go/apps/pushconfig/pushconfigclient"
	"oh-my-helper-go/apps/taskconfig/taskconfigclient"
)

type ServiceContext struct {
	Config        config.Config
	BiliRpc       bilibiliclient.Bilibili
	TaskConfigRpc taskconfigclient.TaskConfig
	PushConfigRpc pushconfigclient.PushConfig
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		BiliRpc:       bilibiliclient.NewBilibili(zrpc.MustNewClient(c.BiliRpc)),
		TaskConfigRpc: taskconfigclient.NewTaskConfig(zrpc.MustNewClient(c.TaskConfigRpc)),
		PushConfigRpc: pushconfigclient.NewPushConfig(zrpc.MustNewClient(c.PushConfigRpc)),
	}
}
