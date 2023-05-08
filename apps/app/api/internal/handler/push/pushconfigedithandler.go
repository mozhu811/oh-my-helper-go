package push

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oh-my-helper-go/apps/app/api/internal/logic/push"
	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"
)

func PushConfigEditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PushConfigEditRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := push.NewPushConfigEditLogic(r.Context(), svcCtx)
		resp, err := l.PushConfigEdit(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
