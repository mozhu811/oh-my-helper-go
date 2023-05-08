package logic

import (
	"context"
	"oh-my-helper-go/apps/bilibili/rpc/internal/model"
	bizerr "oh-my-helper-go/pkg/bizerr"
	"oh-my-helper-go/pkg/cos"
	"oh-my-helper-go/pkg/delegate"

	"oh-my-helper-go/apps/bilibili/rpc/bilibili"
	"oh-my-helper-go/apps/bilibili/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLogic) User(in *bilibili.BiliUserRequest) (*bilibili.BiliUserResponse, error) {

	d := delegate.NewDelegate(delegate.BiliTaskConfig{
		Dedeuserid: in.Dedeuserid,
		Sessdata:   in.Sessdata,
		BiliJct:    in.BiliJct,
	}, false)
	ud, err := d.GetUserDetails()
	if err != nil {
		if err == bizerr.ErrorCookieExpired {
			//logx.Errorf("账号[%v]Cookie已过期，请访问 https://ohmyhelper.com/bilibili/ 重新扫码登陆更新Cookie ❌", in.Dedeuserid)
			l.svcCtx.Db.Model(&model.BilibiliUser{}).
				Where("dedeuserid = ?", in.Dedeuserid).
				Update("is_login", false)
		}
		return nil, err
	}
	face := ud["data"].(map[string]interface{})["face"].(string)
	go func() {
		_ = cos.Upload(face, in.Dedeuserid)
	}()
	bu, err := l.svcCtx.BiliUserModel.FindByDedeuserid(in.Dedeuserid)
	if err != nil {
		return nil, err
	}
	return &bilibili.BiliUserResponse{
		Id:            bu.Id,
		Dedeuserid:    bu.Dedeuserid.String,
		Username:      bu.Username.String,
		Coins:         bu.Coins.Float64,
		Level:         bu.Level.Int64,
		CurrentExp:    bu.CurrentExp.Int64,
		NextExp:       bu.NextExp.Int64,
		UpgradeDays:   bu.UpgradeDays.Int64,
		Medals:        bu.Medals.String,
		VipStatus:     bu.VipStatus.Int64,
		VipType:       bu.VipType.Int64,
		VipLabelTheme: bu.VipLabelTheme.String,
		Sign:          bu.Sign.String,
		IsLogin:       bu.IsLogin.Bool,
		LastRunTime:   bu.LastRunTime.Time.Unix(),
		CreateTime:    bu.CreateTime.Time.Unix(),
	}, nil
}
