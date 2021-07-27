package rpcServer

import (
	"errors"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

type (
	RpcServerConf struct {
		service.ServiceConf
		ListenOn      string
		Nacos         NacosConf    `json:",optional"`
		Auth          bool               `json:",optional"`
		Redis         redis.RedisKeyConf `json:",optional"`
		StrictControl bool               `json:",optional"`
		Timeout       int64              `json:",default=2000"`
		CpuThreshold  int64              `json:",default=900,range=[0:1000]"`
	}

	// A RpcClientConf is a rpc client config.
	RpcClientConf struct {
		Nacos     NacosConf `json:",optional"`
		Endpoints []string        `json:",optional=!Etcd"`
		App       string          `json:",optional"`
		Token     string          `json:",optional"`
		Timeout   int64           `json:",default=2000"`
	}
)


type NacosConf struct {
	Hosts []string
	Key   string
}

// Validate validates c.
func (c NacosConf) Validate() error {
	if len(c.Hosts) == 0 {
		return errors.New("empty etcd hosts")
	} else if len(c.Key) == 0 {
		return errors.New("empty etcd key")
	} else {
		return nil
	}
}

// HasEtcd checks if there is etcd settings in config.
func (sc RpcServerConf) HasNacos() bool {
	return len(sc.Nacos.Hosts) > 0 && len(sc.Nacos.Key) > 0
}

// Validate validates the config.
func (sc RpcServerConf) Validate() error {
	if !sc.Auth {
		return nil
	}

	return sc.Redis.Validate()
}
