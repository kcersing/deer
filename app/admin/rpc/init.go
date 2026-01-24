package rpc

import "admin/rpc/client"

func InitRpc() {

	client.InitSystemRpc()
	client.InitUserRpc()
	client.InitProductRpc()
	client.InitMemberRpc()
	client.InitMessageRpc()
	//client.InitCrmRpc()
	//client.InitOrderRpc()
}
