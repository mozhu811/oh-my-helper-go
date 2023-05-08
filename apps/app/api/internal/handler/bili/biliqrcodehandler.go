package bili

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oh-my-helper-go/apps/app/api/internal/logic/bili"
	"oh-my-helper-go/apps/app/api/internal/svc"
)

func BiliQrCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := bili.NewBiliQrCodeLogic(r.Context(), svcCtx)
		resp, err := l.BiliQrCode()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
