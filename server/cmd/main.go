package main

import (
	"net"
	"os"
	"os/signal"
	"server/internal/config"
	"syscall"

	"server/internal/handlers"
	"server/pkg/hashservice"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen(config.Cfg.PROTO, config.Cfg.PORT)
	if err != nil {
		logrus.WithError(err).Error("Can't init net")
	}
	//инициализируем севис
	s := grpc.NewServer()
	server := &handlers.Server{}
	hashservice.RegisterHashServiceServer(s, server)

	if err := s.Serve(lis); err != nil {

		logrus.WithError(err).Error("Can't serve")
	}

	c := make(chan os.Signal, 1)
	//ловим сигнал завершения
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	//Вызываем штатный завершальщик
	s.GracefulStop()
}
