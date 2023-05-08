package push

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oh-my-helper-go/apps/app/api/internal/logic/push"
	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"
)

func PushConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PushConfigRequest
		cookie, err := r.Cookie("dedeuserid")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		req.Dedeuserid = cookie.Value

		l := push.NewPushConfigLogic(r.Context(), svcCtx)
		resp, err := l.PushConfig(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
