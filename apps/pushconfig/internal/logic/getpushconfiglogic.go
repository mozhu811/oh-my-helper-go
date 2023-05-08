package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"oh-my-helper-go/apps/pushconfig/internal/svc"
	"oh-my-helper-go/apps/pushconfig/pushconfig"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPushConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPushConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPushConfigLogic {
	return &GetPushConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPushConfigLogic) GetPushConfig(in *pushconfig.PushConfigRequest) (*pushconfig.PushConfigResponse, error) {
	pushConfig, err := l.svcCtx.PushConfigModel.FindOneByUserId(l.ctx, in.Dedeuserid)
	if err != nil {
		return nil, err
	}
	var ret = &pushconfig.PushConfigResponse{}
	err = copier.Copy(ret, pushConfig)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
