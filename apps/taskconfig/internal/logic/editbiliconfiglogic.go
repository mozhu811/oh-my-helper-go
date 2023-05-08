package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"oh-my-helper-go/apps/taskconfig/internal/svc"
	"oh-my-helper-go/apps/taskconfig/model"
	"oh-my-helper-go/apps/taskconfig/taskconfig"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditBiliConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditBiliConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditBiliConfigLogic {
	return &EditBiliConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditBiliConfigLogic) EditBiliConfig(in *taskconfig.BiliConfigEditRequest) (*taskconfig.BiliConfigResponse, error) {
	var config = new(model.BiliTaskConfig)
	err := copier.Copy(&config, in)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.BiliTaskConfigModel.Update(l.ctx, config)
	if err != nil {
		return nil, err
	}
	result, err := l.svcCtx.BiliTaskConfigModel.FindOne(l.ctx, config.Id)
	if err != nil {
		return nil, err
	}
	var ret = &taskconfig.BiliConfigResponse{}
	err = copier.Copy(&ret, result)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
