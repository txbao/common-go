package config

import (
    {{.authImport}}
    "github.com/zeromicro/go-zero/core/stores/cache"
    "github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	{{.auth}}

	BaseRpc zrpc.RpcClientConf

    Mysql struct {
        DataSource string
    }
    CacheRedis cache.CacheConf
}
