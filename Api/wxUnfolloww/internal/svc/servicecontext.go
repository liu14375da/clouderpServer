package svc

import (
	"ZeroProject/Api/wxUnfolloww/internal/config"
	"ZeroProject/Api/wxUnfolloww/internal/model/sql"
	"ZeroProject/common/global"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/rest"
)

type ServiceContext struct {
	Config    Config
	UnFollowModel sql.UnFollowSql
}

type Config struct {
	rest.RestConf
}



func NewServiceContext(c Config) *ServiceContext {
	// redis 参数连接
	cluster := cache.CacheConf{
		cache.NodeConf{
			RedisConf:
			redis.RedisConf{
				Host: global.Redis.Host,
				Pass: global.Redis.Pass,
			},
			Weight: 100,
		},
	}
	//获取数据库连接
	conn := sqlx.NewSqlConn("mssql", global.UnifiedLogin.SqlConn)
	return &ServiceContext{
		Config: c,
		UnFollowModel: sql.NewUnFollowSql(conn,cluster),
	}
}


func ClientConfig(c Config) config.Config {
	// api 服务对应的（名称,ip,端口）
	c.Name = global.WxUnfollowwApiClient.Name
	c.Host = global.WxUnfollowwApiClient.Host
	c.Port = global.WxUnfollowwApiClient.Port
	// 日志
	c.Log.Mode = global.WxUnfollowwApiClient.LogMode
	c.Log.Path = global.WxUnfollowwApiClient.Path
	c.Log.Level = global.WxUnfollowwApiClient.Level
	return (config.Config)(c)
}