package client

import (
	"gen/kitex_gen/member/memberservice"
)

type RpcServer struct{}

func NewRpcServer(memberClient memberservice.Client) *RpcServer {
	return &RpcServer{}
}

func (r *RpcServer) MemberRpc() (rpc memberservice.Client, err error) {
	MemberClient, _ = r.MemberRpc()
	rpc = MemberClient
	return
}
