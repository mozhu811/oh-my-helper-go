package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oh-my-helper-go/apps/taskconfig/internal/config"
	"oh-my-helper-go/apps/taskconfig/model"
)

type ServiceContext struct {
	Config              config.Config
	BiliTaskConfigModel model.BiliTaskConfigModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:              c,
		BiliTaskConfigModel: model.NewBiliTaskConfigModel(conn),
	}
}
