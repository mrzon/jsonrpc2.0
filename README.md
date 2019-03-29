# jsonrpc2.0
Json RPC 2.0 implementation using http protocol. This json rpc implementation hides the explicit rpc call. You just need to define a service skeleton, register it as client and the function in it can be called just like a native function. 

This will work if your existing server already support json rpc 2.0. You just need to port the skeleton method to Go.

# Example
## Math Common package
This struct will be the empty "interface" of math function. Math_common will be publicly exported.
The client will not care about the implementation because it will only depend on the math service.
```
package math_common

type MathService struct {
	Add       func(A int, B int) int `jsonrpc:"add"`
	Substract func(A int, B int) int
	Multiply  func(A int, B int) int
	Modulo    func(A int, B int) int
}
```
The tag `jsonrpc:"add"` will translate the function name of function "Add". This is beneficial if you call a non Golang rpc service (which they might name their method in camel case).

## Math Server package
This package will contains a function definition of the math functions. It will implement the math_common functions. 
```
package math_server

import ".../math-common"

func NewMathServiceImpl() math_common.MathService{
	mImpl := math_common.MathService{
		Add: func(A int, B int) int {
			return (A + B)
		},
		Substract: func(A int, B int) int {
			return A - B
		},
		Multiply: func(A int, B int) int {
			return A * B
		},
		Modulo: func(A int, B int) int {
			return A % B
		},
	}
	return mImpl
}
```

## Math Server main 
```
package main

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
```

## Math Client main 
```
package main

mathClient := &math_common.MathService{}

var mathRpcClient = &client.RpcClient{
  Service: mathClient,
  Config: client.Config{
    Host:     "localhost",
    Port:     7890,
    EndPoint: "/math",
    Timeout:  60 * time.Second,
  },
}

fmt.Println(mathClient.Add(1, 2))       //print 3
fmt.Println(mathClient.Modulo(4, 3))    //print 1
fmt.Println(mathClient.Substract(5, 2)) //print 3
fmt.Println(mathClient.Multiply(10, 2)) //print 20
```

For more inquiry: contact me.
