package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// rpc
	UserFavoriteRpc zrpc.RpcClientConf
	UserRpc         zrpc.RpcClientConf
	UserCommentRpc  zrpc.RpcClientConf
	// kq
	UserCommentOptServiceConf  kq.KqConf
	UserFavoriteOptServiceConf kq.KqConf
}
