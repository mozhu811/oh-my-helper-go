package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oh-my-helper-go/apps/bilibili/rpc/internal/config"
	"oh-my-helper-go/apps/bilibili/rpc/internal/model"
	"oh-my-helper-go/pkg/orm"
	"time"

	"github.com/jinzhu/gorm"
)

type ServiceContext struct {
	Config        config.Config
	BiliUserModel model.BilibiliUserModel
	Db            *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:        c,
		BiliUserModel: model.NewBilibiliUserModel(conn),
		Db: orm.NewMysql(&orm.Config{
			DSN:         c.DataSource,
			Active:      20,
			Idle:        10,
			IdleTimeout: time.Hour * 24,
		}),
	}
}
