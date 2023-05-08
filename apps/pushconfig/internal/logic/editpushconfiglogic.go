package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"oh-my-helper-go/apps/pushconfig/internal/model"

	"oh-my-helper-go/apps/pushconfig/internal/svc"
	"oh-my-helper-go/apps/pushconfig/pushconfig"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditPushConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditPushConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditPushConfigLogic {
	return &EditPushConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditPushConfigLogic) EditPushConfig(in *pushconfig.PushConfigEditRequest) (*pushconfig.PushConfigResponse, error) {
	err := l.svcCtx.PushConfigModel.Update(l.ctx, &model.PushConfig{
		Id:     in.Id,
		UserId: in.UserId,
		Config: in.Config,
	})
	if err != nil {
		return nil, err
	}
	pushConfig, err := l.svcCtx.PushConfigModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	var ret = new(pushconfig.PushConfigResponse)
	err = copier.Copy(&ret, pushConfig)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
