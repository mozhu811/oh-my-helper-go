package logic

import (
	"context"

	"oh-my-helper-go/apps/taskconfig/internal/svc"
	"oh-my-helper-go/apps/taskconfig/taskconfig"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBiliConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBiliConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBiliConfigLogic {
	return &GetBiliConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBiliConfigLogic) GetBiliConfig(in *taskconfig.BiliConfigRequest) (*taskconfig.BiliConfigResponse, error) {
	btc, err := l.svcCtx.BiliTaskConfigModel.GetConfigByDedeuserid(in.Dedeuserid)
	if err != nil {
		return nil, err
	}

	return &taskconfig.BiliConfigResponse{
		Id:                 btc.Id,
		Dedeuserid:         btc.Dedeuserid,
		Sessdata:           btc.Sessdata,
		BiliJct:            btc.BiliJct,
		DonateCoins:        btc.DonateCoins,
		ReserveCoins:       btc.ReserveCoins,
		AutoCharge:         btc.AutoCharge,
		DonateGift:         btc.DonateGift,
		AutoChargeTarget:   btc.AutoChargeTarget,
		DonateGiftTarget:   btc.DonateGiftTarget,
		DevicePlatform:     btc.DevicePlatform,
		DonateCoinStrategy: btc.DonateCoinStrategy,
		UserAgent:          btc.UserAgent,
		SkipTask:           btc.SkipTask,
		FollowDeveloper:    btc.FollowDeveloper,
		CreateTime:         btc.CreateTime.Unix(),
		UpdateTime:         btc.UpdateTime.Unix(),
	}, nil
}
