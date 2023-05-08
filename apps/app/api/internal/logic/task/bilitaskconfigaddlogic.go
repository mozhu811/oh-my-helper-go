package task

import (
	"context"
	"github.com/jinzhu/copier"
	"oh-my-helper-go/apps/taskconfig/taskconfig"

	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BiliTaskConfigAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBiliTaskConfigAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BiliTaskConfigAddLogic {
	return &BiliTaskConfigAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BiliTaskConfigAddLogic) BiliTaskConfigAdd(req *types.BiliTaskConfigAddRequest) (resp *types.BiliTaskConfigResponse, err error) {
	var newConfig = new(taskconfig.BiliConfigAddRequest)
	err = copier.Copy(&newConfig, &req)
	if err != nil {
		return nil, err
	}
	config, err := l.svcCtx.TaskConfigRpc.AddBiliConfig(l.ctx, newConfig)
	if err != nil {
		return nil, err
	}
	var ret = new(types.BiliTaskConfigResponse)
	err = copier.Copy(&ret, &config)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
