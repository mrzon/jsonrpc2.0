package client

import (
	"time"
	"strconv"
)

type Config struct {
	Host          string
	Port          int
	EndPoint      string
	Timeout       time.Duration
	EnableLogging bool
}

type RpcClient struct {
	Config  Config
	Service interface{}
}

func (c Config) GetAddr() string {
	if c.Port == 0 {
		c.Port = 80
	}

	protocol := "http"
	portStr := ":" + strconv.Itoa(c.Port);
	if c.Port == 443 {
		protocol = "https"
		portStr = ""
	}

	return protocol + "://" + c.Host + portStr + c.EndPoint;
}
