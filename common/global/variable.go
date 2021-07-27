package global

import (
	"ZeroProject/nacos/configFile"
)

/*
	该全局对应nacos的配置列表
*/
var (
	// 获取统一登录数据库链接
	UnifiedLogin = configFile.NacosConfigure("sqlserver-10.28.0.6", "sqlConn", "unifiedLogin")

	// 获取jwt密钥
	JwtAuth = configFile.NacosConfigure("jwt-secret", "jwt", "jwt")

	// redis 配置信息
	Redis = configFile.NacosConfigure("redis-unifiedLogin", "redis", "Rpc")

	// unifiedLoginRpcServer 统一登录 RPC 配置信息
	UnifiedLoginRpcServer = configFile.NacosConfigure("unifiedLoginRpc", "rpcServer", "Rpc")

	// unifiedLoginRpcServer 统一登录 API 配置信息
	UnifiedLoginApiClient = configFile.NacosConfigure("unifiedLoginApi", "apiClient", "Api")

	//注册nacos服务，所需的配置参数
	RegisterServiceNacos = configFile.RegisterServiceNacosConfigure("nacos", "nacos", "rpc")

	// WxUnfollowwApiClient API 配置信息
	WxUnfollowwApiClient = configFile.NacosConfigure("wxUnfollowwApiClient", "apiClient", "api")

	// 企业微信 配置
	QyWxRpc = configFile.NacosConfigure("Wechat-bindingRpc", "QyWxRpc", "Rpc")

	// 企业微信的 api 配置
	QyWxApi = configFile.NacosConfigure("Wechat-bindingApi", "QyWxApi", "Api")

	// auth 的配置参数
	AuthToken = configFile.NacosConfigure("TokenVerificationApi", "apiClient", "Api")
)
