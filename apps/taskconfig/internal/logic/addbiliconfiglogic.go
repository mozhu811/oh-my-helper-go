package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"oh-my-helper-go/apps/taskconfig/model"

	"oh-my-helper-go/apps/taskconfig/internal/svc"
	"oh-my-helper-go/apps/taskconfig/taskconfig"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBiliConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBiliConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBiliConfigLogic {
	return &AddBiliConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddBiliConfigLogic) AddBiliConfig(in *taskconfig.BiliConfigAddRequest) (*taskconfig.BiliConfigResponse, error) {
	var biliConfig = new(model.BiliTaskConfig)
	err := copier.Copy(&biliConfig, &in)
	if err != nil {
		return nil, err
	}
	res, err := l.svcCtx.BiliTaskConfigModel.Insert(l.ctx, biliConfig)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	biliConfig, err = l.svcCtx.BiliTaskConfigModel.FindOne(l.ctx, id)
	if err != nil {
		return nil, err
	}
	var ret = new(taskconfig.BiliConfigResponse)
	err = copier.Copy(&ret, &biliConfig)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
