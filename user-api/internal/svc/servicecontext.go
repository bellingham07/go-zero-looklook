package svc

import (
	"database/sql"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-looklook/user-api/internal/config"
	"go-zero-looklook/user-api/internal/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"log"
)

type ServiceContext struct {
	Config         config.Config
	DB             *gorm.DB
	TestMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := Init(c)
	return &ServiceContext{
		Config:         c,
		DB:             db,
		TestMiddleware: middleware.NewTestMiddleware().Handle,
	}
}

func Init(c config.Config) (db *gorm.DB) {
	var (
		sqlDB *sql.DB
		err   error
	)
	postgresConf := postgres.Config{
		DSN: c.Postgres.DSN,
	}
	gormConfig := configLog(c.Postgres.LogMode)
	if db, err = gorm.Open(postgres.New(postgresConf), gormConfig); err != nil {
		log.Fatal("opens database failed: ", err)
	}
	if sqlDB, err = db.DB(); err != nil {
		log.Fatal("db.db() failed: ", err)
	}

	sqlDB.SetMaxIdleConns(c.Postgres.MaxIdleCons)
	sqlDB.SetMaxOpenConns(c.Postgres.MaxOpenCons)
	return
}

// configLog 根据配置决定是否开启日志
func configLog(mod bool) (c *gorm.Config) {
	if mod {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名不加复数形式，false默认加
			},
		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名不加复数形式，false默认加
			},
		}
	}
	return
}
