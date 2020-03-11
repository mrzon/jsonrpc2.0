package server

import "time"

type Config struct {
	Port          int
	EndPoint      string
	Timeout       time.Duration
	EnableLogging bool
}

type RpcServer struct {
	Config  Config
	Service interface{}
}
