package logic

import (
	"black/api/model"
	"context"

	"black/api/internal/svc"
	"black/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) error {
	// todo: add your logic here and delete this line

	// 为了简单，这里只做一下简单的逻辑校验
	_, err := l.svcCtx.UserModel.FindOneByMobile(req.Username)
	if err == nil {
		return errorDuplicateUsername
	}

	_, err = l.svcCtx.UserModel.FindOneByMobile(req.Mobile)
	if err == nil {
		return errorDuplicateMobile
	}

	_, err = l.svcCtx.UserModel.Insert(model.User{
		Name:     req.Username,
		Password: req.Password,
		Mobile:   req.Mobile,
		Gender:   "男",
		Nickname: "admin",
	})
	return err
}
