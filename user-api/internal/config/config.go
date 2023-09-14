package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Postgres struct {
		DSN         string
		LogMode     bool
		MaxOpenCons int
		MaxIdleCons int
	}
}
