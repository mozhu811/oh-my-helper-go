package push

import (
	"context"
	"encoding/json"
	"oh-my-helper-go/apps/pushconfig/pushconfigclient"

	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPushConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushConfigLogic {
	return &PushConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PushConfigLogic) PushConfig(req *types.PushConfigRequest) (resp *types.PushConfigResponse, err error) {
	pushConfig, err := l.svcCtx.PushConfigRpc.GetPushConfig(l.ctx, &pushconfigclient.PushConfigRequest{
		Dedeuserid: req.Dedeuserid,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.PushConfigResponse)
	err = json.Unmarshal([]byte(pushConfig.Config), &resp.Config)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	resp.Id = pushConfig.Id
	resp.UserId = pushConfig.UserId
	return resp, nil
}
