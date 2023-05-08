package logic

import (
	"context"

	"oh-my-helper-go/apps/pushconfig/internal/svc"
	"oh-my-helper-go/apps/pushconfig/pushconfig"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPushConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPushConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPushConfigLogic {
	return &AddPushConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPushConfigLogic) AddPushConfig(in *pushconfig.PushConfigAddRequest) (*pushconfig.PushConfigResponse, error) {
	// todo: add your logic here and delete this line

	return &pushconfig.PushConfigResponse{}, nil
}
