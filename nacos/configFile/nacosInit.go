package configFile

import (
	"ZeroProject/model/entity/nacos"
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var (
	defaultConfig *viper.Viper
)

type LocalHostNacos struct {
	NacosIp     []string
	NacosPort   []uint64
	NamespaceId []string
}

func NacosConfigure(DataId, conn, Group string) *nacos.NaCosConfig {
	switch conn {
	case "sqlConn":
		host := GetConfig(DataId, Group)
		return Connect(host) //统一登录数据库链接
	case "jwt":
		host := GetConfig(DataId, Group)
		return NewJwt(host) // jwt认证信息
	case "redis":
		host := GetConfig(DataId, Group)
		return NewRedis(host) // Redis配置信息
	case "rpcServer":
		host := GetConfig(DataId, Group)
		return NewRpcServer(host) // rpcServer 配置信息
	case "apiClient":
		host := GetConfig(DataId, Group)
		return NewApiClient(host) // apiClient 配置信息
	case "QyWxRpc":
		host := GetConfig(DataId, Group)
		return NewQyWxRpc(host) // 企业微信 rpc 配置信息
	case "QyWxApi":
		host := GetConfig(DataId, Group)
		return NewQyWxApi(host) // 企业微信 api 配置信息
	default:
		return nil
	}
}

func RegisterServiceNacosConfigure(DataId, conn, Group string) *nacos.RegisterServiceConf {
	switch conn {
	case "nacos":
		//host := GetConfig(DataId, Group)
		return NewNcos() // apiClient 配置信息
	default:
		return nil
	}
}

func GetConfig(DataId string, group string) *viper.Viper {
	return initConfig(DataId, group)
}

func NewViper() *viper.Viper {
	config := viper.New()
	path, _ := os.Getwd()
	config.SetConfigName("configFile")
	config.AddConfigPath(path + "/config")
	config.SetConfigType("yaml")
	err := config.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败, 异常信息 : ", err)
	}
	return config
}

func initConfig(DataId string, group string) *viper.Viper {
	config := NewViper()
	st := new(LocalHostNacos)
	MarshalArrayString(config.Get("Nacos.IP"), &st.NacosIp)
	MarshalArrayInt(config.Get("Nacos.Port"), &st.NacosPort)
	MarshalArrayString(config.Get("Nacos.NamespaceId"), &st.NamespaceId)
	if len(st.NacosIp) == len(st.NacosPort) && len(st.NacosIp) == len(st.NamespaceId) {
		for i, v := range st.NacosIp {
			defaultConfig = viper.New()
			defaultConfig.SetConfigType("yaml")
			implement := true

			//配置模型
			serverConfigs := []constant.ServerConfig{
				{IpAddr: v, Port: st.NacosPort[i]},
			}
			//fmt.Println(v, st.NacosPort[i], st.NamespaceId[i])
			//客户端
			nacosClient, err := clients.NewConfigClient(
				vo.NacosClientParam{
					ClientConfig: &constant.ClientConfig{
						TimeoutMs:   5000,
						NamespaceId: st.NamespaceId[i], // 命名空间 名称id,例如 dev,Prod
					},
					ServerConfigs: serverConfigs,
				},
			)
			if err != nil {
				implement = false
				log.Fatal("nacos初始化错误:", err)
			}

			content, err := nacosClient.GetConfig(vo.ConfigParam{DataId: DataId, Group: group})
			if err != nil {
				implement = false
				fmt.Println("nacos读取配置错误:" + content + ",请检查Nacos是否在线")
			}

			err = defaultConfig.ReadConfig(strings.NewReader(content))
			if err != nil {
				implement = false
				log.Fatalln("Viper解析配置失败:", err)
			}

			err = nacosClient.ListenConfig(vo.ConfigParam{
				DataId: DataId,
				Group:  group,
				OnChange: func(namespace, group, dataId, data string) {
					fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
					err = defaultConfig.ReadConfig(strings.NewReader(data))
					if err != nil {
						implement = false
						log.Fatalln("Viper解析配置失败:", err)
					}
				},
			})
			if implement == true {
				break
			}
		}
	} else {
		log.Fatalln("Nacos的参数配置数量不一致")
	}
	return defaultConfig
}

func NewRedis(host *viper.Viper) *nacos.NaCosConfig {
	st := new(nacos.NaCosConfig)
	st.Host = host.GetString("host")
	st.Pass = host.GetString("pass")
	return st
}

func Connect(host *viper.Viper) *nacos.NaCosConfig {
	connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;encrypt=disable",
		host.GetString("Sqlserver.Server"), host.GetString("Sqlserver.Database"),
		host.GetString("Sqlserver.User"), host.GetString("Sqlserver.Password"),
		//host.GetString("Sqlserver.Port"),
	)
	st := new(nacos.NaCosConfig)
	st.SqlConn = connString
	return st
}

func NewJwt(host *viper.Viper) *nacos.NaCosConfig {
	st := new(nacos.NaCosConfig)
	st.AccessExpire = host.GetInt64("AccessExpire")  // jwt 过期时间
	st.AccessSecret = host.GetString("AccessSecret") // jwt 密匙
	return st
}

func NewRpcServer(host *viper.Viper) *nacos.NaCosConfig {
	st := new(nacos.NaCosConfig)
	st.Hosts = host.GetString("Nacos.Hosts")  // etcd 的ip
	st.Key = host.GetString("Nacos.Key")      // etcd 对应的key值(rpc 对应 api)

	st.Name = host.GetString("Name")         // rpc 服务名称
	st.ListenOn = host.GetString("ListenOn") // Rpc 服务监听地址

	st.LogMode = host.GetString("Log.Mode")  // 日志模式
	st.Path = host.GetString("Log.Path")     // 日志路径
	st.Level = host.GetString("Log.Level")   // 日志级别

	// Prometheus（普罗米修斯报警配置信息）
	st.Host = host.GetString("Prometheus.Host") //普罗米修斯 ip地址
	st.Port = host.GetInt("Prometheus.Port")    //普罗米修斯 端口
	st.Path = host.GetString("Prometheus.Path") //普罗米修斯 路径
	return st
}

// 获取naCos 配置中心的对应的企业微信rpc配置参数
func NewQyWxRpc(host *viper.Viper) *nacos.NaCosConfig {
	st := new(nacos.NaCosConfig)
	st.Corpid = host.GetString("QyWx.Corpid")         // 企业微信 id
	st.Corpsecret = host.GetString("QyWx.Corpsecret") // 企业微信 secret

	// 企业微信绑定
	st.BingDingName = host.GetString("qywxbinding.Name")
	st.BingDingListenOn = host.GetString("qywxbinding.ListenOn")

	st.BingDingEtcdHost = host.GetString("qywxbinding.Etcd.Hosts")
	st.BingDingEtcdKey = host.GetString("qywxbinding.Etcd.Key")

	st.LogMode = host.GetString("qywxbinding.Log.Mode") // 日志模式
	st.Path = host.GetString("qywxbinding.Log.Path")    // 日志路径
	st.Level = host.GetString("qywxbinding.Log.Level")  // 日志级别

	// Prometheus（普罗米修斯报警配置信息）
	st.PrometheusHost = host.GetString("qywxbinding.Prometheus.Host") //普罗米修斯 ip地址
	st.PrometheusPort = host.GetInt("qywxbinding.Prometheus.Port")    //普罗米修斯 端口
	st.PrometheusPath = host.GetString("qywxbinding.Prometheus.Path") //普罗米修斯 路径
	return st
}

// 获取naCos 配置中心的对应的企业微信api配置参数
func NewQyWxApi(host *viper.Viper) *nacos.NaCosConfig {
	st := new(nacos.NaCosConfig)
	// 企业微信绑定
	st.BingDingName = host.GetString("qywxbinding.Name")
	st.BingDingHost = host.GetString("qywxbinding.Host")
	st.BingDingPort = host.GetInt("qywxbinding.Port")

	st.BingDingEtcdHost = host.GetString("qywxbinding.QyWxRpc.Etcd.Hosts")
	st.BingDingEtcdKey = host.GetString("qywxbinding.QyWxRpc.Etcd.Key")

	st.LogMode = host.GetString("qywxbinding.Log.Mode") // 日志模式
	st.Path = host.GetString("qywxbinding.Log.Path")    // 日志路径
	st.Level = host.GetString("qywxbinding.Log.Level")  // 日志级别

	// Prometheus（普罗米修斯报警配置信息）
	st.PrometheusHost = host.GetString("qywxbinding.Prometheus.Host") //普罗米修斯 ip地址
	st.PrometheusPort = host.GetInt("qywxbinding.Prometheus.Port")    //普罗米修斯 端口
	st.PrometheusPath = host.GetString("qywxbinding.Prometheus.Path") //普罗米修斯 路径
	return st
}

func NewApiClient(host *viper.Viper) *nacos.NaCosConfig {
	st := new(nacos.NaCosConfig)
	st.Host = host.GetString("Host")                // api 端的ip
	st.Name = host.GetString("Name")                // api 端的名称
	st.Port = host.GetInt("Port")                   // api 的端口

	st.Hosts = host.GetString("UserRpc.Etcd.Hosts") //etcd 的ip
	st.Key = host.GetString("UserRpc.Etcd.Key")     //  api 对应 rpc（etcd的key值）

	st.LogMode = host.GetString("Log.Mode")         // 日志模式
	st.Path = host.GetString("Log.Path")            // 日志路径
	st.Level = host.GetString("Log.Level")          // 日志级别
	// Prometheus（普罗米修斯报警配置信息）
	st.PrometheusHost = host.GetString("Prometheus.Host") //普罗米修斯 ip地址
	st.PrometheusPort = host.GetInt("Prometheus.Port")    //普罗米修斯 端口
	st.PrometheusPath = host.GetString("Prometheus.Path") //普罗米修斯 路径
	return st
}

//注册nacos服务所需的参数
func NewNcos() *nacos.RegisterServiceConf {
	host := NewViper()
	st := new(nacos.RegisterServiceConf)
	MarshalArrayString(host.Get("Nacos.IP"), &st.IP)
	MarshalArrayInt(host.Get("Nacos.Port"), &st.NacosProt)
	MarshalArrayString(host.Get("Nacos.NamespaceId"), &st.NamespaceId)
	MarshalArrayString(host.Get("Nacos.Name"), &st.Name)
	//st.LogDir = host.GetString("Nacos.LogDir")
	//st.CacheDir = host.GetString("Nacos.CacheDir")
	MarshalArrayString(host.Get("Nacos.RotateTime"), &st.RotateTime)
	MarshalArrayString(host.Get("Nacos.LogLevel"), &st.LogLevel)

	MarshalInt(host.Get("RegisterService.Prot"), &st.Prot)
	MarshalArrayString(host.Get("RegisterService.GroupName"), &st.GroupName)
	MarshalString(host.Get("RegisterService.ClusterName"), &st.ClusterName)
	return st
}

func MarshalArrayString(data interface{}, st *[]string) {
	resByre, _ := json.Marshal(data)
	_ = json.Unmarshal(resByre, &st)
}

func MarshalArrayInt(data interface{}, st *[]uint64) {
	resByre, _ := json.Marshal(data)
	_ = json.Unmarshal(resByre, &st)
}

func MarshalString(data interface{}, st *string) {
	resByre, _ := json.Marshal(data)
	_ = json.Unmarshal(resByre, &st)
}

func MarshalInt(data interface{}, st *uint64) {
	resByre, _ := json.Marshal(data)
	_ = json.Unmarshal(resByre, &st)
}