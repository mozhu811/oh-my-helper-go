package logic

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/skip2/go-qrcode"
	"io"
	"net/http"

	"oh-my-helper-go/apps/bilibili/rpc/bilibili"
	"oh-my-helper-go/apps/bilibili/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoginQrCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type biliQrCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		URL       string `json:"url"`
		QrcodeKey string `json:"qrcode_key"`
	} `json:"data"`
}

func NewGetLoginQrCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginQrCodeLogic {
	return &GetLoginQrCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLoginQrCodeLogic) GetLoginQrCode(in *bilibili.QrCodeRequest) (*bilibili.QrCodeResponse, error) {
	resp, err := http.Get("https://passport.bilibili.com/x/passport-login/web/qrcode/generate")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			l.Logger.Error(err)
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	/*
		{
		    "code": 0,
		    "message": "0",
		    "ttl": 1,
		    "data": {
		        "url": "https://passport.bilibili.com/h5-app/passport/login/scan?navhide=1&qrcode_key=cc47ba24d6660b2e1dd9aa36ca0b7b13&from=",
		        "qrcode_key": "cc47ba24d6660b2e1dd9aa36ca0b7b13"
		    }
		}
	*/
	var qrCode = &biliQrCode{}
	err = json.Unmarshal(body, &qrCode)
	if err != nil {
		return nil, err
	}

	url := qrCode.Data.URL
	// 将url进行Base64编码
	img, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return &bilibili.QrCodeResponse{
		QrCodeUrl: url,
		QrCodeImg: "data:image/png;base64," + base64.StdEncoding.EncodeToString(img),
		QrCodeKey: qrCode.Data.QrcodeKey,
	}, nil
}
