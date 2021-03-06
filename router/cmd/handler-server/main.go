// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"os"
	"os/signal"
	"router/internal/config"
	"router/internal/handler"
	"router/internal/handler/operations"
	"syscall"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Embedded(handler.SwaggerJSON, handler.FlatSwaggerJSON)
	if err != nil {
		config.Logger.WithError(err)
	}

	api := operations.NewHandlerAPI(swaggerSpec)
	server := handler.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Итоговое задание. Хэши."
	parser.LongDescription = "Данный сервис должен, взаимодействуя с сервисом считающим хэши (по выбранному вами протоколу), получать из входящих строк их хэши, сохранять их в свою БД (выбор так же за вами) с присвоем id, по которым далее можно будет запрашивать хэши."
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			config.Logger.WithError(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}
	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		config.Logger.WithError(err)
	}
	//shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	server.Shutdown()

}
