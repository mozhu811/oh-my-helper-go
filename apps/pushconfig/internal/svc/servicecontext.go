package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oh-my-helper-go/apps/pushconfig/internal/config"
	"oh-my-helper-go/apps/pushconfig/internal/model"
)

type ServiceContext struct {
	Config          config.Config
	PushConfigModel model.PushConfigModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:          c,
		PushConfigModel: model.NewPushConfigModel(conn),
	}
}
