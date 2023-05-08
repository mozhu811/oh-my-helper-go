package bili

import (
	"context"
	"oh-my-helper-go/apps/bilibili/rpc/bilibiliclient"
	"oh-my-helper-go/pkg/cos"
	"oh-my-helper-go/pkg/delegate"

	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BiliQrCodeLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBiliQrCodeLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BiliQrCodeLoginLogic {
	return &BiliQrCodeLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BiliQrCodeLoginLogic) BiliQrCodeLogin(req *types.BiliLoginRequest) (resp *types.BiliLoginResponse, err error) {
	r, err := l.svcCtx.BiliRpc.Login(l.ctx, &bilibiliclient.BiliLoginRequest{
		QrCodeKey: req.QrCodeKey,
	})
	if err != nil {
		return nil, err
	}

	if r.Code == 0 {
		go func() {
			// todo 更新数据库biliuser
		}()
		go func() {
			// todo 更新cookie
		}()

		// 上传头像
		bd := delegate.NewDelegate(delegate.BiliTaskConfig{
			Dedeuserid: r.Dedeuserid,
			Sessdata:   r.Sessdata,
			BiliJct:    r.BiliJct,
		}, false)

		details, _ := bd.GetUserDetails()

		face := details["data"].(map[string]interface{})["face"].(string)
		_ = cos.Upload(face, r.Dedeuserid)
	}

	return &types.BiliLoginResponse{
		Code:       int(r.Code),
		Dedeuserid: r.Dedeuserid,
		Sessdata:   r.Sessdata,
		BiliJct:    r.BiliJct,
	}, nil
}
