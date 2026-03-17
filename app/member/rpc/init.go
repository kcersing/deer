package rpc

import "member/rpc/client"

func Init() {

	client.InitProductRpc()
	client.InitOrderRpc()

}
