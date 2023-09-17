package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpcConf zrpc.RpcClientConf
	Postgres    struct {
		DSN         string
		LogMode     bool
		MaxOpenCons int
		MaxIdleCons int
	}
}
