package task

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oh-my-helper-go/apps/app/api/internal/logic/task"
	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"
)

func BiliTaskConfigEditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = types.BiliTaskConfigEditRequest{}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := task.NewBiliTaskConfigEditLogic(r.Context(), svcCtx)
		resp, err := l.BiliTaskConfigEdit(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			w.WriteHeader(201)
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
