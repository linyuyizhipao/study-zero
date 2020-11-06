package logic

import (
	"context"

	"black/api/internal/svc"
	"black/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReturnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnLogic {
	return ReturnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnLogic) Return(req types.ReturnReq) error {
	// todo: add your logic here and delete this line

	return nil
}
