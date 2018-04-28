package client

import (
	"time"
	"strconv"
)

type Config struct {
	Host string
	Port int
	EndPoint string
	Timeout time.Duration
}

type RpcClient struct{
	Config Config
	Service interface{}
}

func (c Config) GetAddr() string {
	if c.Port == 0 {
		c.Port = 80
	}
	return c.Host + ":" + strconv.Itoa(c.Port) + c.EndPoint;
}