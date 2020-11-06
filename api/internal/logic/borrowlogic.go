package logic

import (
	"context"

	"black/api/internal/svc"
	"black/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type BorrowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBorrowLogic(ctx context.Context, svcCtx *svc.ServiceContext) BorrowLogic {
	return BorrowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BorrowLogic) Borrow(req types.BorrowReq) error {
	// todo: add your logic here and delete this line

	return nil
}
