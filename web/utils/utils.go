package utils

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"

)

// 初始化micro客户端
func InitMicro() micro.Service{
	consulReg := consul.NewRegistry()
	consulService := micro.NewService(
		micro.Registry(consulReg),
		)
	return consulService
}
