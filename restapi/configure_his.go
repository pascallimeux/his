package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/pascallimeux/his/restapi/operations"
	"github.com/pascallimeux/his/restapi/operations/consents"
	"github.com/pascallimeux/his/restapi/operations/version"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../swagger.json

func configureFlags(api *operations.HisAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HisAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.ConsentsCreateConsentHandler = consents.CreateConsentHandlerFunc(func(params consents.CreateConsentParams) middleware.Responder {
		return middleware.NotImplemented("operation consents.CreateConsent has not yet been implemented")
	})
	api.ConsentsDeleteConsentHandler = consents.DeleteConsentHandlerFunc(func(params consents.DeleteConsentParams) middleware.Responder {
		return middleware.NotImplemented("operation consents.DeleteConsent has not yet been implemented")
	})
	api.ConsentsDeleteConsentsHandler = consents.DeleteConsentsHandlerFunc(func(params consents.DeleteConsentsParams) middleware.Responder {
		return middleware.NotImplemented("operation consents.DeleteConsents has not yet been implemented")
	})
	api.ConsentsGetConsentHandler = consents.GetConsentHandlerFunc(func(params consents.GetConsentParams) middleware.Responder {
		return middleware.NotImplemented("operation consents.GetConsent has not yet been implemented")
	})
	api.ConsentsGetConsentsHandler = consents.GetConsentsHandlerFunc(func(params consents.GetConsentsParams) middleware.Responder {
		return middleware.NotImplemented("operation consents.GetConsents has not yet been implemented")
	})
	api.ConsentsGetConsents4ConsumerHandler = consents.GetConsents4ConsumerHandlerFunc(func(params consents.GetConsents4ConsumerParams) middleware.Responder {
		return middleware.NotImplemented("operation consents.GetConsents4Consumer has not yet been implemented")
	})
	api.ConsentsGetConsents4ConsumerOwnerHandler = consents.GetConsents4ConsumerOwnerHandlerFunc(func(params consents.GetConsents4ConsumerOwnerParams) middleware.Responder {
		return middleware.NotImplemented("operation consents.GetConsents4ConsumerOwner has not yet been implemented")
	})
	api.ConsentsGetConsents4OwnerHandler = consents.GetConsents4OwnerHandlerFunc(func(params consents.GetConsents4OwnerParams) middleware.Responder {
		return middleware.NotImplemented("operation consents.GetConsents4Owner has not yet been implemented")
	})
	api.VersionGetVersionHandler = version.GetVersionHandlerFunc(func(params version.GetVersionParams) middleware.Responder {
		return middleware.NotImplemented("operation version.GetVersion has not yet been implemented")
	})
	api.ConsentsIsConsentHandler = consents.IsConsentHandlerFunc(func(params consents.IsConsentParams) middleware.Responder {
		return middleware.NotImplemented("operation consents.IsConsent has not yet been implemented")
	})

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
func configureServer(s *graceful.Server, scheme, addr string) {
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
