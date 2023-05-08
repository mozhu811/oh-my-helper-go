package task

import (
	"context"
	"oh-my-helper-go/apps/taskconfig/taskconfig"
	"oh-my-helper-go/pkg/delegate"

	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BiliTaskConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBiliTaskConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BiliTaskConfigLogic {
	return &BiliTaskConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BiliTaskConfigLogic) BiliTaskConfig(req *types.BiliTaskConfigRequest) (resp *types.BiliTaskConfigResponse, err error) {
	bd := delegate.NewDelegate(delegate.BiliTaskConfig{
		Dedeuserid: req.Dedeuserid,
		Sessdata:   req.Sessdata,
		BiliJct:    req.BiliJct,
	}, false)
	_, err = bd.GetUserDetails()
	if err != nil {
		return nil, err
	}

	bc, err := l.svcCtx.TaskConfigRpc.GetBiliConfig(l.ctx, &taskconfig.BiliConfigRequest{Dedeuserid: req.Dedeuserid})
	if err != nil {
		return nil, err
	}

	return &types.BiliTaskConfigResponse{
		Id:                 bc.Id,
		DonateCoins:        bc.DonateCoins,
		ReserveCoins:       bc.ReserveCoins,
		AutoCharge:         bc.AutoCharge,
		DonateGift:         bc.DonateGift,
		AutoChargeTarget:   bc.AutoChargeTarget,
		DonateGiftTarget:   bc.DonateGiftTarget,
		DevicePlatform:     bc.DevicePlatform,
		DonateCoinStrategy: bc.DonateCoinStrategy,
		UserAgent:          bc.UserAgent,
		SkipTask:           bc.SkipTask,
		FollowDeveloper:    bc.FollowDeveloper,
		CreateTime:         bc.CreateTime,
		UpdateTime:         bc.UpdateTime,
	}, nil
}
