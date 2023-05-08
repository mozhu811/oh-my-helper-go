package task

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oh-my-helper-go/apps/app/api/internal/logic/task"
	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"
)

func BiliTaskConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
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

		var req = types.BiliTaskConfigRequest{
			Dedeuserid: dedeuserid.Value,
			Sessdata:   sessdata.Value,
			BiliJct:    biliJct.Value,
		}

		l := task.NewBiliTaskConfigLogic(r.Context(), svcCtx)
		resp, err := l.BiliTaskConfig(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
