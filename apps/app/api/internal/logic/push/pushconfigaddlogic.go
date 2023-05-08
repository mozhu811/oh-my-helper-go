package push

import (
	"context"

	"oh-my-helper-go/apps/app/api/internal/svc"
	"oh-my-helper-go/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushConfigAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPushConfigAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushConfigAddLogic {
	return &PushConfigAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PushConfigAddLogic) PushConfigAdd(req *types.PushConfigAddRequest) (resp *types.PushConfigResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
