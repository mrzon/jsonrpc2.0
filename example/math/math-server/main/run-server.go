package main

import (
	"github.com/mrzon/jsonrpc2.0/example/math/math-server"
	"github.com/mrzon/jsonrpc2.0/server"
)

func main()  {
	var rpcServerConn = server.NewRpcServerConnection()
	var mathRpcServer = &server.RpcServer{
		Config: server.Config{
			Port:     7890,
			EndPoint: "/math",
		},
		Service: math_server.NewMathServiceImpl(),
	}
	rpcServerConn.Register(mathRpcServer)
	rpcServerConn.StartAndServe()
}
