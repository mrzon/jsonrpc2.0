package main

import (
	"fmt"
	"time"

	common "github.com/mrzon/jsonrpc2.0/example/math/math-common"

	client "github.com/mrzon/jsonrpc2.0/client"
)

func main() {
	mathServiceClient := common.MathService{}
	var mathRpcClient = &client.RpcClient{
		Service: &mathServiceClient,
		Config: client.Config{
			Host:          "localhost",
			Port:          7890,
			EndPoint:      "/math",
			Timeout:       60 * time.Second,
			EnableLogging: true,
		},
	}
	client.Register(mathRpcClient)
	fmt.Println(mathServiceClient.Add(6, 2))
	fmt.Println(mathServiceClient.Modulo(4, 3))
	fmt.Println(mathServiceClient.Substract(5, 2))
	fmt.Println(mathServiceClient.Multiply(10, 2))
}
