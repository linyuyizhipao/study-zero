package logic

import (
	"context"
	"fmt"

	"black/rpc/internal/greet"
	"black/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *greet.SayRequest) (*greet.SayResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println(in.Ping,"hello 打出来的ping")
	return &greet.SayResponse{Pong: "say hello"}, nil
}
