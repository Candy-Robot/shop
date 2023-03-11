package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"user/handler"
	"user/model"
	pb "user/proto/user"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "user"
	version = "latest"
)

func main() {
	// 初始化 mysql 连接池
	model.InitDb()
	// 使用consul服务发现
	consulReg := consul.NewRegistry()

	// Create service
	srv := micro.NewService(
		micro.Address("127.0.0.1:12345"),
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
	)

	// Register handler
	pb.RegisterUserHandler(srv.Server(), new(handler.User))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
