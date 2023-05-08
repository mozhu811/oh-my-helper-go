package push

import (
	"context"
	"encoding/json"
	"oh-my-helper-go/apps/pushconfig/pushconfigclient"

	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushConfigEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPushConfigEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushConfigEditLogic {
	return &PushConfigEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PushConfigEditLogic) PushConfigEdit(req *types.PushConfigEditRequest) (resp *types.PushConfigResponse, err error) {
	jsonStr, err := json.Marshal(req.Config)
	result, err := l.svcCtx.PushConfigRpc.EditPushConfig(l.ctx, &pushconfigclient.PushConfigEditRequest{
		Id:     req.Id,
		UserId: req.UserId,
		Config: string(jsonStr),
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.PushConfigResponse)
	err = json.Unmarshal([]byte(result.Config), &resp.Config)
	if err != nil {
		return nil, err
	}
	resp.Id = result.Id
	resp.UserId = result.UserId
	return resp, nil
}
