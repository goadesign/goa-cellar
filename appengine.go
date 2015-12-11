// +build appengine

package main

import (
	"net/http"
	"os"

	"github.com/raphael/goa"
	"github.com/raphael/goa-middleware/cors"
	"github.com/raphael/goa/examples/cellar/app"
	"github.com/raphael/goa/examples/cellar/controllers"
	"github.com/raphael/goa/examples/cellar/swagger"
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
	service.Use(goa.RequestID())
	service.Use(cors.Middleware(spec))
	service.Use(goa.Recover())

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
