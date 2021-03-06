// Code generated by goctl. DO NOT EDIT!
// Source: greet.proto

//go:generate mockgen -destination ./greet_mock.go -package greet -source $GOFILE

package greet

import (
	"context"

	"black/rpc/internal/greet"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Request     = greet.Request
	Response    = greet.Response
	SayRequest  = greet.SayRequest
	SayResponse = greet.SayResponse

	Greet interface {
		Ping(ctx context.Context, in *Request) (*Response, error)
		SayHello(ctx context.Context, in *SayRequest) (*SayResponse, error)
	}

	defaultGreet struct {
		cli zrpc.Client
	}
)

func NewGreet(cli zrpc.Client) Greet {
	return &defaultGreet{
		cli: cli,
	}
}

func (m *defaultGreet) Ping(ctx context.Context, in *Request) (*Response, error) {
	client := greet.NewGreetClient(m.cli.Conn())
	return client.Ping(ctx, in)
}

func (m *defaultGreet) SayHello(ctx context.Context, in *SayRequest) (*SayResponse, error) {
	client := greet.NewGreetClient(m.cli.Conn())
	return client.SayHello(ctx, in)
}
