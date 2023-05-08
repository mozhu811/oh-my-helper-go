package bili

import (
	"context"
	"oh-my-helper-go/apps/bilibili/rpc/bilibili"

	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BiliUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBiliUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BiliUserLogic {
	return &BiliUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BiliUserLogic) BiliUser(req *types.BiliUserListRequest) (resp *types.BiliUserListResponse, err error) {
	r, err := l.svcCtx.BiliRpc.Users(l.ctx, &bilibili.BiliUserListRequest{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	var users []types.BiliUserResponse
	for _, user := range r.BiliUsers {
		users = append(users, types.BiliUserResponse{
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
		})
	}
	return &types.BiliUserListResponse{
		Total: r.Total,
		Users: users,
	}, nil
}
