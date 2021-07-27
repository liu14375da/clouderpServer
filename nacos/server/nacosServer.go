package server

import (
	"ZeroProject/common/global"
	"ZeroProject/common/tool"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func RegisterService(registerName string, metadata map[string]string) {
	if len(global.RegisterServiceNacos.IP) == len(global.RegisterServiceNacos.NacosProt) && len(global.RegisterServiceNacos.IP) == len(global.RegisterServiceNacos.NamespaceId) {
		for i, v := range global.RegisterServiceNacos.IP {
			sc := []constant.ServerConfig{
				{
					IpAddr: v,
					Port:   global.RegisterServiceNacos.NacosProt[i],
				},
			}
			cc := constant.ClientConfig{
				NamespaceId:         global.RegisterServiceNacos.NamespaceId[i],
				TimeoutMs:           5000,
				NotLoadCacheAtStart: true,
				//LogDir:              global.RegisterServiceNacos.LogDir,
				//CacheDir:            global.RegisterServiceNacos.CacheDir,
				RotateTime: global.RegisterServiceNacos.RotateTime[i],
				MaxAge:     3,
				LogLevel:   global.RegisterServiceNacos.LogLevel[i],
			}
			client, err := clients.NewNamingClient(
				vo.NacosClientParam{
					ClientConfig:  &cc,
					ServerConfigs: sc,
				},
			)

			if err != nil {
				panic(err)
			}
			//服务注册
			ExampleserviceclientRegisterserviceinstance(client, vo.RegisterInstanceParam{
				Ip:          tool.LocalIP(),
				Port:        global.RegisterServiceNacos.Prot,
				ServiceName: registerName,
				Weight:      10,
				Enable:      true,
				Healthy:     true,
				Ephemeral:   true,
				Metadata:    metadata,
				GroupName:   registerName,                               //分组
				ClusterName: global.RegisterServiceNacos.ClusterName, //集群名
			})

			//获取健康实例
			//ExampleServiceClientSelectOneHealthyInstance(client, vo.SelectOneHealthInstanceParam{
			//	ServiceName: registerName,
			//	GroupName:   global.RegisterServiceNacos.GroupName[i], //分组
			//})
		}
	}
}

func ExampleserviceclientRegisterserviceinstance(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	success, _ := client.RegisterInstance(param)
	fmt.Printf("RegisterServiceInstance,result:%+v \n\n", success)
}

func ExampleServiceClientSelectOneHealthyInstance(client naming_client.INamingClient, param vo.SelectOneHealthInstanceParam) bool {
	instances, _ := client.SelectOneHealthyInstance(param)
	return instances.Healthy
	//fmt.Printf("SelectInstances,param:%+v, result:%+v \n\n", param, instances)
}
