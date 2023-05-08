package bili

import (
	"context"
	"oh-my-helper-go/apps/bilibili/rpc/bilibiliclient"

	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BiliUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBiliUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BiliUserInfoLogic {
	return &BiliUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BiliUserInfoLogic) BiliUserInfo(req *types.BiliUserRequest) (resp *types.BiliUserResponse, err error) {
	user, err := l.svcCtx.BiliRpc.User(l.ctx, &bilibiliclient.BiliUserRequest{
		Dedeuserid: req.Dedeuserid,
		Sessdata:   req.Sessdata,
		BiliJct:    req.BiliJct,
	})
	if err != nil {
		return nil, err
	}

	return &types.BiliUserResponse{
		Id:            user.Id,
		Dedeuserid:    user.Dedeuserid,
		Username:      user.Username,
		Coins:         user.Coins,
		CurrentExp:    user.CurrentExp,
		NextExp:       user.NextExp,
		IsLogin:       user.IsLogin,
		UpgradeDays:   user.UpgradeDays,
		Level:         user.Level,
		Medals:        user.Medals,
		VipStatus:     user.VipStatus,
		VipType:       user.VipType,
		VipLabelTheme: user.VipLabelTheme,
		Sign:          user.Sign,
		LastRunTime:   user.LastRunTime,
		CreateTime:    user.CreateTime,
	}, nil
}
