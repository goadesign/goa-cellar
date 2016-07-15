// +build appengine

package main

import (
	"net/http"
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar/app"
	"github.com/goadesign/goa-cellar/controllers"
	"github.com/goadesign/goa/logging/log15"
	"github.com/goadesign/goa/middleware"
	"github.com/inconshreveable/log15"
)

const (
	projectID  = "goa-cellar"
	bucketName = "artifacts.cellar.goa.design"
)

func init() {
	// Configure logging for appengine
	logger := log15.New()
	logger.SetHandler(log15.StreamHandler(os.Stderr, log15.LogfmtFormat()))

	// Create goa application
	service := goa.New("cellar")
	service.WithLogger(goalog15.New(logger))

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount account controller onto application
	ac := controllers.NewAccount(service)
	app.MountAccountController(service, ac)

	// Mount bottle controller onto application
	bc := controllers.NewBottle(service)
	app.MountBottleController(service, bc)

	// Setup HTTP handler
	http.HandleFunc("/", service.Mux.ServeHTTP)
}
