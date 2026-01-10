package rpc

import "admin/rpc/client"

func InitRpc() {
	//client.InitCrmRpc()
	client.InitMemberRpc()
	client.InitOrderRpc()
	client.InitSystemRpc()
	client.InitUserRpc()
	client.InitProductRpc()
}
