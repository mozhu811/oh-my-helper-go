package bili

import (
	"context"
	"oh-my-helper-go/apps/bilibili/rpc/bilibiliclient"

	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BiliQrCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBiliQrCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BiliQrCodeLogic {
	return &BiliQrCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BiliQrCodeLogic) BiliQrCode() (resp *types.QrCodeResponse, err error) {
	qrCodeResponse, err := l.svcCtx.BiliRpc.GetLoginQrCode(l.ctx, &bilibiliclient.QrCodeRequest{})
	if err != nil {
		return nil, err
	}

	return &types.QrCodeResponse{
		QrCodeUrl: qrCodeResponse.QrCodeUrl,
		QrCodeImg: qrCodeResponse.QrCodeImg,
		QrCodeKey: qrCodeResponse.QrCodeKey,
	}, nil
}
