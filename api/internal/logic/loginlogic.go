package logic

import (
	"black/api/model"
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"

	"black/api/internal/svc"
	"black/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.UserReply, error) {
	// todo: add your logic here and delete this line
	// 忽略逻辑校验
	userInfo, err := l.svcCtx.UserModel.FindOneByName(req.Username)
	switch err {
	case nil:
		if userInfo.Password != req.Password {
			return nil, errorIncorrectPassword
		}
		now := time.Now().Unix()
		accessExpire := l.svcCtx.Config.Auth.AccessExpire
		jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire)
		if err != nil {
			return nil, err
		}

		return &types.UserReply{
			Id:       userInfo.Id,
			Username: userInfo.Name,
			Mobile:   userInfo.Mobile,
			Nickname: userInfo.Nickname,
			Gender:   userInfo.Gender,
			JwtToken: types.JwtToken{
				AccessToken:  jwtToken,
				AccessExpire: now + accessExpire,
				RefreshAfter: now + accessExpire/2,
			},
		}, nil
	case model.ErrNotFound:
		return nil, errorUsernameUnRegister
	default:
		return nil, err
	}
}


func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
