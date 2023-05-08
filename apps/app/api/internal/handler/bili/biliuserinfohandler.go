package bili

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oh-my-helper-go/apps/app/api/internal/logic/bili"
	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"
)

func BiliUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dedeuserid, err := r.Cookie("dedeuserid")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		sessdata, err := r.Cookie("sessdata")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		biliJct, err := r.Cookie("biliJct")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		var req = types.BiliUserRequest{
			Dedeuserid: dedeuserid.Value,
			Sessdata:   sessdata.Value,
			BiliJct:    biliJct.Value,
		}
		l := bili.NewBiliUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.BiliUserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
