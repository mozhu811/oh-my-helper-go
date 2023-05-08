package logic

import (
	"context"

	"oh-my-helper-go/apps/taskconfig/internal/svc"
	"oh-my-helper-go/apps/taskconfig/taskconfig"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *taskconfig.Request) (*taskconfig.Response, error) {
	// todo: add your logic here and delete this line

	return &taskconfig.Response{}, nil
}
