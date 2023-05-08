package logic

import (
	"context"

	"oh-my-helper-go/apps/pushconfig/internal/svc"
	"oh-my-helper-go/apps/pushconfig/pushconfig"

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

func (l *PingLogic) Ping(in *pushconfig.Request) (*pushconfig.Response, error) {
	// todo: add your logic here and delete this line

	return &pushconfig.Response{}, nil
}
