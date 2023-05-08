package logic

import (
	"context"

	"oh-my-helper-go/apps/bilibili/rpc/bilibili"
	"oh-my-helper-go/apps/bilibili/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersLogic {
	return &UsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UsersLogic) Users(in *bilibili.BiliUserListRequest) (*bilibili.BiliUserListResponse, error) {
	users, total, err := l.svcCtx.BiliUserModel.List(in.Page, in.Size)
	if err != nil {
		return nil, err
	}
	var biliUsers []*bilibili.BiliUserResponse
	for _, u := range users {
		biliUsers = append(biliUsers, &bilibili.BiliUserResponse{
			Id:            u.Id,
			Dedeuserid:    u.Dedeuserid.String,
			Username:      u.Username.String,
			Coins:         u.Coins.Float64,
			Level:         u.Level.Int64,
			CurrentExp:    u.CurrentExp.Int64,
			NextExp:       u.NextExp.Int64,
			UpgradeDays:   u.UpgradeDays.Int64,
			Medals:        u.Medals.String,
			VipStatus:     u.VipStatus.Int64,
			VipType:       u.VipType.Int64,
			VipLabelTheme: u.VipLabelTheme.String,
			Sign:          u.Sign.String,
			IsLogin:       u.IsLogin.Bool,
			LastRunTime:   u.LastRunTime.Time.Unix(),
			CreateTime:    u.CreateTime.Time.Unix(),
		})
	}
	return &bilibili.BiliUserListResponse{
		Total:     total,
		BiliUsers: biliUsers,
	}, nil
}
