package task

import (
	"context"
	"github.com/jinzhu/copier"
	"oh-my-helper-go/apps/taskconfig/taskconfigclient"

	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BiliTaskConfigEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBiliTaskConfigEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BiliTaskConfigEditLogic {
	return &BiliTaskConfigEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BiliTaskConfigEditLogic) BiliTaskConfigEdit(req *types.BiliTaskConfigEditRequest) (resp *types.BiliTaskConfigResponse, err error) {
	var newConfig = new(taskconfigclient.BiliConfigEditRequest)
	err = copier.Copy(newConfig, req)
	if err != nil {
		return nil, err
	}
	retConfig, err := l.svcCtx.TaskConfigRpc.EditBiliConfig(l.ctx, newConfig)
	if err != nil {
		return nil, err
	}
	var ret = new(types.BiliTaskConfigResponse)
	err = copier.Copy(&ret, retConfig)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
