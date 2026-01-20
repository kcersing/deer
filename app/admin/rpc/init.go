package rpc

import "admin/rpc/client"

func InitRpc() {

	client.InitSystemRpc()
	client.InitUserRpc()
	client.InitProductRpc()
	client.InitMemberRpc()
	//client.InitCrmRpc()
	//client.InitOrderRpc()
}
