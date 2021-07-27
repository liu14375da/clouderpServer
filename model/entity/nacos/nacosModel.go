package nacos

/*
	存放在nacos配置列表中需要的参数
*/
type NaCosConfig struct {
	SqlConn          string // 统一登录数据库链接
	AccessSecret     string // jwt 密匙
	AccessExpire     int64  // jwt 过期时间
	Host             string // Redis IP
	Pass             string // Redis 密码
	Key              string // Etcd(rpc 对应 api etcd 的key值)
	Name             string // Rpc 服务名称
	ListenOn         string // Rpc 服务监听地址
	Hosts            string // Etcd ip地址
	LogMode          string // 日志模式
	Path             string // 日志路径
	Level            string // 日志级别
	Port             int    // api 端的端口
	Corpid           string // 企业微信 id
	Corpsecret       string // 企业微信 secret
	PrometheusHost   string //普罗米修斯 ip
	PrometheusPath   string //普罗米修斯 路径
	PrometheusPort   int    //普罗米修斯 端口
	BingDingName     string //企业微信绑定名称
	BingDingListenOn string //企业微信绑定监听端口
	BingDingHost     string //企业微信绑定 api 端的 ip
	BingDingPort     int    //企业微信绑定端口
	BingDingEtcdHost string //企业微信绑定etcd IP
	BingDingEtcdKey  string //企业微信绑定etcd key
}

/*
	注册到nacos中的服务所需要的参数
*/
type RegisterServiceConf struct {
	Name        []string // nacos 中定义的环境名称（dev,Prod）
	IP          []string // Naocs ip地址
	NamespaceId []string // 命名空间的id
	NacosProt   []uint64 // nacos 端口
	LogDir      string   // nacos 日志地址
	CacheDir    string   // nacos 缓冲地址
	RotateTime  []string // nacos 日志转换时间
	LogLevel    []string // nacos 日志级别
	Prot        uint64 // nacos 服务注册端口
	GroupName   []string // nacos 服务注册分组名称
	ClusterName string // nacos 服务注册集群名称

}


