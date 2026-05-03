package config

import (
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

type Config struct {
	kitex.RpcServerConf
	UserClient kitex.RpcClientConf
	ChatClient kitex.RpcClientConf
	SyncClient kitex.RpcClientConf
}
