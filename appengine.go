// +build appengine

package main

import (
	"net/http"
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar/app"
	"github.com/goadesign/goa-cellar/controllers"
	"github.com/goadesign/goa-cellar/swagger"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/cors"
	"github.com/goadesign/goa/middleware/security/basicauth"
	"gopkg.in/inconshreveable/log15.v2"
)

const (
	projectID  = "goa-cellar"
	bucketName = "artifacts.cellar.goa.design"
)

func init() {
	// Configure logging for appengine
	goa.Log.SetHandler(log15.StreamHandler(os.Stderr, log15.LogfmtFormat()))

	// Create goa application
	service := goa.New("cellar")

	// Setup CORS to allow for swagger UI.
	spec, err := cors.New(func() {
		cors.Origin("*", func() {
			cors.Resource("*", func() {
				cors.Methods("GET", "POST", "PUT", "PATCH", "DELETE")
				cors.Headers("*")
			})
		})
	})
	if err != nil {
		panic(err)
	}

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(cors.Middleware(spec))
	service.Use(middleware.Recover())
	app.ConfigureAdminPassSecurity(service, basicauth.New("wine", "lover"))

	// Mount account controller onto application
	ac := controllers.NewAccount(service)
	app.MountAccountController(service, ac)

	// Mount bottle controller onto application
	bc := controllers.NewBottle(service)
	app.MountBottleController(service, bc)

	// Mount Swagger Spec controller onto application
	swagger.MountController(service)

	// Mount CORS preflight controllers
	cors.MountPreflightController(service, spec)

	// Setup HTTP handler
	http.HandleFunc("/", service.HTTPHandler().ServeHTTP)
}
