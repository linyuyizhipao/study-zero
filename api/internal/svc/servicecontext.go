package svc

import (
	"black/api/internal/config"
	"black/api/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	um := model.NewUserModel(conn,c.CacheRedis)
	return &ServiceContext{
		Config:    c,
		UserModel: um,
	}
}