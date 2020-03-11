package main

import (
	math_server "github.com/mrzon/jsonrpc2.0/example/math/math-server"
	server "github.com/mrzon/jsonrpc2.0/server"
)

func main() {
	var rpcServerConn = server.NewRpcServerConnection()
	var mathRpcServer = &server.RpcServer{
		Config: server.Config{
			Port:          7890,
			EndPoint:      "/math",
			EnableLogging: true,
		},
		Service: math_server.NewMathServiceImpl(),
	}
	rpcServerConn.Register(mathRpcServer)
	rpcServerConn.StartAndServe()
}
