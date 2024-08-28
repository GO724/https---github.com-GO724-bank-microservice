// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"bank-microservice/api/restapi/operations"
	"bank-microservice/api/restapi/operations/admins"
)

//go:generate swagger generate server --target ../../api --name BcardMicroserviceAPI --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.BcardMicroserviceAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BcardMicroserviceAPIAPI) http.Handler {
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

	if api.AdminsGetBanksHandler == nil {
		api.AdminsGetBanksHandler = admins.GetBanksHandlerFunc(func(params admins.GetBanksParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.GetBanks has not yet been implemented")
		})
	}
	if api.AdminsGetCardsHandler == nil {
		api.AdminsGetCardsHandler = admins.GetCardsHandlerFunc(func(params admins.GetCardsParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.GetCards has not yet been implemented")
		})
	}
	if api.AdminsGetPersonsHandler == nil {
		api.AdminsGetPersonsHandler = admins.GetPersonsHandlerFunc(func(params admins.GetPersonsParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.GetPersons has not yet been implemented")
		})
	}
	if api.AdminsAddCardHandler == nil {
		api.AdminsAddCardHandler = admins.AddCardHandlerFunc(func(params admins.AddCardParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.AddCard has not yet been implemented")
		})
	}
	if api.AdminsAddPersonHandler == nil {
		api.AdminsAddPersonHandler = admins.AddPersonHandlerFunc(func(params admins.AddPersonParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.AddPerson has not yet been implemented")
		})
	}
	if api.AdminsAddbankHandler == nil {
		api.AdminsAddbankHandler = admins.AddbankHandlerFunc(func(params admins.AddbankParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.Addbank has not yet been implemented")
		})
	}
	if api.AdminsDelBankHandler == nil {
		api.AdminsDelBankHandler = admins.DelBankHandlerFunc(func(params admins.DelBankParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.DelBank has not yet been implemented")
		})
	}
	if api.AdminsDelPersonHandler == nil {
		api.AdminsDelPersonHandler = admins.DelPersonHandlerFunc(func(params admins.DelPersonParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.DelPerson has not yet been implemented")
		})
	}
	if api.AdminsDelСardHandler == nil {
		api.AdminsDelСardHandler = admins.DelСardHandlerFunc(func(params admins.DelСardParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.DelСard has not yet been implemented")
		})
	}
	if api.AdminsGetBankHandler == nil {
		api.AdminsGetBankHandler = admins.GetBankHandlerFunc(func(params admins.GetBankParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.GetBank has not yet been implemented")
		})
	}
	if api.AdminsGetCardHandler == nil {
		api.AdminsGetCardHandler = admins.GetCardHandlerFunc(func(params admins.GetCardParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.GetCard has not yet been implemented")
		})
	}
	if api.AdminsGetPersonHandler == nil {
		api.AdminsGetPersonHandler = admins.GetPersonHandlerFunc(func(params admins.GetPersonParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.GetPerson has not yet been implemented")
		})
	}
	if api.AdminsIsValidHandler == nil {
		api.AdminsIsValidHandler = admins.IsValidHandlerFunc(func(params admins.IsValidParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.IsValid has not yet been implemented")
		})
	}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
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
