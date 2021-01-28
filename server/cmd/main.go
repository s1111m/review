package main

import (
	"net"
	"server/internal/config"

	"server/internal/handlers"
	"server/pkg/hashservice"

	"google.golang.org/grpc"
)

func main() {
	//config.Init()
	lis, err := net.Listen(config.Cfg.PROTO, config.Cfg.PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	server := &handlers.Server{}
	hashservice.RegisterHashServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
