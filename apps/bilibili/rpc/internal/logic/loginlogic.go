package logic

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"oh-my-helper-go/apps/bilibili/rpc/bilibili"
	"oh-my-helper-go/apps/bilibili/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type QrCodeStat struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		URL          string `json:"url"`
		RefreshToken string `json:"refresh_token"`
		Timestamp    int    `json:"timestamp"`
		Code         int    `json:"code"`
		Message      string `json:"message"`
	} `json:"data"`
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *bilibili.BiliLoginRequest) (*bilibili.BiliLoginResponse, error) {
	r, err := http.Get("https://passport.bilibili.com/x/passport-login/web/qrcode/poll" + "?qrcode_key=" + in.QrCodeKey)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			l.Logger.Error(err)
		}
	}(r.Body)

	var qrCodeStat QrCodeStat
	err = json.NewDecoder(r.Body).Decode(&qrCodeStat)
	if err != nil {
		return nil, err
	}
	code := qrCodeStat.Data.Code
	if code == 0 {
		// 解析Cookie
		var dedeuserid string
		var sessdata string
		var biliJct string
		for _, cookie := range r.Cookies() {
			switch cookie.Name {
			case "DedeUserID":
				dedeuserid = cookie.Value
			case "SESSDATA":
				sessdata = cookie.Value
			case "bili_jct":
				biliJct = cookie.Value
			}
		}

		return &bilibili.BiliLoginResponse{
			Code:       uint32(code),
			Dedeuserid: dedeuserid,
			Sessdata:   sessdata,
			BiliJct:    biliJct,
		}, nil
	}

	return &bilibili.BiliLoginResponse{
		Code: uint32(code),
	}, nil
}
