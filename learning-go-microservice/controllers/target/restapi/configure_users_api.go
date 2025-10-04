// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"learning-go-microservice/controllers/target/restapi/operations"
)

//go:generate swagger generate server --target ../../target --name UsersAPI --spec ../../../routes/swagger/userSwagger.yaml --principal any

func configureFlags(api *operations.UsersAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
	_ = api
}

func configureAPI(api *operations.UsersAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...any)
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.DeleteUserUserIDHandler == nil {
		api.DeleteUserUserIDHandler = operations.DeleteUserUserIDHandlerFunc(func(params operations.DeleteUserUserIDParams) middleware.Responder {
			_ = params

			return middleware.NotImplemented("operation operations.DeleteUserUserID has not yet been implemented")
		})
	}
	if api.GetUserUserIDHandler == nil {
		api.GetUserUserIDHandler = operations.GetUserUserIDHandlerFunc(func(params operations.GetUserUserIDParams) middleware.Responder {
			_ = params

			return middleware.NotImplemented("operation operations.GetUserUserID has not yet been implemented")
		})
	}
	if api.GetUsersHandler == nil {
		api.GetUsersHandler = operations.GetUsersHandlerFunc(func(params operations.GetUsersParams) middleware.Responder {
			_ = params

			return middleware.NotImplemented("operation operations.GetUsers has not yet been implemented")
		})
	}
	if api.PostUserHandler == nil {
		api.PostUserHandler = operations.PostUserHandlerFunc(func(params operations.PostUserParams) middleware.Responder {
			_ = params

			return middleware.NotImplemented("operation operations.PostUser has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
	_ = tlsConfig
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(server *http.Server, scheme, addr string) {
	_ = server
	_ = scheme
	_ = addr
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
