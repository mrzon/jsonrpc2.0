package main

import (
	"github.com/mrzon/jsonrpc2.0/client"
	"time"
	"github.com/mrzon/jsonrpc2.0/example/math/math-common"
	"fmt"
)

func main()  {
	mathServiceClient := math_common.MathService{}
	var mathRpcClient = &client.RpcClient{
		Service: &mathServiceClient,
		Config: client.Config{
			Host:     "localhost",
			Port:     7890,
			EndPoint: "/math",
			Timeout:  60 * time.Second,
		},
	}
	client.Register(mathRpcClient)
	fmt.Println(mathServiceClient.Add(6, 2))
	fmt.Println(mathServiceClient.Modulo(4, 3))
	fmt.Println(mathServiceClient.Substract(5, 2))
	fmt.Println(mathServiceClient.Multiply(10, 2))
}
