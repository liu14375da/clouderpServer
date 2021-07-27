package svc

import (
	"ZeroProject/common/global"
	"ZeroProject/model/sql/unifiedLogin"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	zr        zrpc.RpcServerConf
	UserModel unifiedLogin.UserModel
}

/*
	nacos配置中心，获取配置信息
*/
func NewServiceContext(zr zrpc.RpcServerConf) *ServiceContext {
	// redis 参数连接
	cluster := cache.CacheConf{
		cache.NodeConf{
			RedisConf: redis.RedisConf{
				Host: global.Redis.Host,
				Pass: global.Redis.Pass,
			},
			Weight: 100,
		},
	}
	//获取数据库连接
	conn := sqlx.NewSqlConn("mssql", global.UnifiedLogin.SqlConn)
	return &ServiceContext{
		zr:        zr,
		UserModel: unifiedLogin.NewUserModel(conn, cluster),
	}
}

// 添加统一登录对应的配置文件（etcd的ip地址,key值和rpc服务名称,监听地址,日志）
func RpcServer(zr zrpc.RpcServerConf) zrpc.RpcServerConf {

	//zr.Etcd.Hosts = append(zr.Etcd.Hosts, global.UnifiedLoginRpcServer.Hosts)
	//zr.Etcd.Key = global.UnifiedLoginRpcServer.Key

	zr.Name = global.UnifiedLoginRpcServer.Name
	zr.ListenOn = global.UnifiedLoginRpcServer.ListenOn
	// 日志相关
	zr.ServiceConf.Log.Mode = global.UnifiedLoginRpcServer.LogMode
	zr.ServiceConf.Log.Path = global.UnifiedLoginRpcServer.Path
	zr.ServiceConf.Log.Level = global.UnifiedLoginRpcServer.Level
	// Prometheus（普罗米修斯报警配置信息）
	zr.ServiceConf.Prometheus.Host = global.UnifiedLoginRpcServer.Host
	zr.ServiceConf.Prometheus.Port = global.UnifiedLoginRpcServer.Port
	zr.ServiceConf.Prometheus.Path = global.UnifiedLoginRpcServer.Path
	return zr
}
