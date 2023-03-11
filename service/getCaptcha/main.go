package main

import (
	"getCaptcha/handler"
	pb "getCaptcha/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"

)

var (
	service = "getcaptcha"
	version = "latest"
)

func main() {
	consulReg := consul.NewRegistry()
	// Create service
	srv := micro.NewService(
		micro.Address("127.0.0.1:13579"),	// 防止随机生成port
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
	)
	srv.Init()

	// Register handler
	pb.RegisterGetCaptchaHandler(srv.Server(), new(handler.GetCaptcha))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
