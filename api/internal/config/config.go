package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql struct { // mysql配置
		DataSource string
	}
	CacheRedis cache.ClusterConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	LibraryRpc zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
}
