// This file is safe to edit. Once it exists it will not be overwritten

package handler

import (
	"crypto/tls"
	"net/http"
	"router/internal/db"
	"router/internal/handler/operations"
	"router/models"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

//go:generate swagger generate server --target ..\..\..\router --name Handler --spec ..\..\..\api\api.yaml --server-package ./internal/handler --principal interface{}

func configureFlags(api *operations.HandlerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HandlerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetCheckHandler = operations.GetCheckHandlerFunc(func(params operations.GetCheckParams) middleware.Responder {
		var result models.ArrayOfHash
		//функция gorm умеет искать сразу по массиву, конвертим строковой params.Ids в [] int
		// пропускаем ошибки конвертации (просто не ищем эти записи), и если в итоге строк 0 - считаем что функция вернула ошибку и возвращаем NoContent
		result, err := db.DbFind(params.Ids)
		//ничего не нашли
		if err != nil {
			response := operations.NewGetCheckNoContent()
			return response
		}
		//если нашли
		response := operations.NewPostSendOK()
		response.SetPayload(result)
		return response
	})

	api.PostSendHandler = operations.PostSendHandlerFunc(func(params operations.PostSendParams) middleware.Responder {
		//	return middleware.NotImplemented("operation operations.PostSend has not yet been implemented")
		//var result models.ArrayOfHash

		result, err := sendToHashservice(params.Params)
		if err != nil {
			return operations.NewPostSendInternalServerError()
		}
		payload := db.DbWrite(convertProtoHashArrayToArrayOfHashes(result))
		response := operations.NewPostSendOK()
		response.SetPayload(payload)
		return response
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {

}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
