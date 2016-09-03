// +build !appengine,!appenginevm

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar/app"
	"github.com/goadesign/goa-cellar/controllers"
	"github.com/goadesign/goa-cellar/store"
	"github.com/goadesign/goa/logging/log15"
	"github.com/goadesign/goa/middleware"
	"github.com/inconshreveable/log15"
)

func main() {
	// Create goa service
	service := goa.New("cellar")

	// Setup logger
	logger := log15.New()
	service.WithLogger(goalog15.New(logger))

	// Setup basic middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Setup database connection
	db := store.NewDB()

	// Mount account controller onto service
	ac := controllers.NewAccount(service, db)
	app.MountAccountController(service, ac)

	// Mount bottle controller onto service
	bc := controllers.NewBottle(service, db)
	app.MountBottleController(service, bc)

	// Run service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError(err.Error())
	}
}
