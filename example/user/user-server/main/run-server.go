package main

import (
	"github.com/mrzon/jsonrpc2.0/server"
	"github.com/mrzon/jsonrpc2.0/example/user/user-server"
)

func main()  {
	var rpcServerConn = server.NewRpcServerConnection()

	userService := user_server.NewUserServiceImpl()

	var userRpcServer = &server.RpcServer{
		Config: server.Config{
			Port:     8101,
			EndPoint: "/user",
		},
		Service: userService,
	}

	rpcServerConn.Register(userRpcServer)
	rpcServerConn.StartAndServe()
}
