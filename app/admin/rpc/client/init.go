package client

import "github.com/cloudwego/hertz/pkg/common/hlog"

func Init() {
	hlog.Info("加载rpc")
	initSystemRpc()
	initUserRpc()
	initMemberRpc()
	//initProductRpc()
	//initOrderRpc()
}
