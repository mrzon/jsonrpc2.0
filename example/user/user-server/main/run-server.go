package main

import (
	"github.com/mrzon/jsonrpc2.0/server"
	"github.com/mrzon/jsonrpc2.0/example/user/user-server"
)

func main()  {
	var rpcServerConn = server.NewRpcServerConnection()
	var userRpcServer = &server.RpcServer{
		Config: server.Config{
			Port:     8101,
			EndPoint: "/user",
		},
		Service: user_server.NewUserServiceImpl(),
	}
	rpcServerConn.Register(userRpcServer)
	rpcServerConn.StartAndServe()
}
