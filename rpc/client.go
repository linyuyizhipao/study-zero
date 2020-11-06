package main

import (
	"black/rpc/internal/greet"
	"context"
	"fmt"
	"log"
	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/zrpc"
)

func main() {
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "rpc.rpc",
		},
	})

	hello := greet.NewGreetClient(client.Conn())

	reply, err := hello.SayHello(context.Background(), &greet.SayRequest{Ping: "go-zero-hello"})
	if err != nil {
		log.Fatal(err)
	}
	reply2, err := hello.Ping(context.Background(), &greet.Request{Ping: "go-zero-hello"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.Pong,"fdsfdfdf",reply2.Pong)
}