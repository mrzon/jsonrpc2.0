package main

import (
	"github.com/mrzon/jsonrpc2.0/client"
	"github.com/mrzon/jsonrpc2.0/example/user/user-common"
	"time"
	"fmt"
)

func main()  {
	userClient := &user_common.UserService{}
	var userRpcClient = &client.RpcClient{
		Service: userClient,
		Config: client.Config{
			Host:     "localhost",
			Port:     8101,
			EndPoint: "/user",
			Timeout:  60 * time.Second,
		},
	}
	client.Register(userRpcClient)

	fmt.Println(userClient.Register(user_common.RegisterUserSpec{
		"emerson",
		"123",
	}))

	fmt.Println(userClient.Login(user_common.LoginUserSpec{
		"mrzon",
		"123",
	}))

	fmt.Println(userClient.Login(user_common.LoginUserSpec{
		"emerson",
		"1234",
	}))

	fmt.Println(userClient.Login(user_common.LoginUserSpec{
		"emerson",
		"123",
	}))
}