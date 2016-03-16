// +build !appengine

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar/app"
	"github.com/goadesign/goa-cellar/controllers"
	"github.com/goadesign/goa-cellar/js"
	"github.com/goadesign/goa-cellar/schema"
	"github.com/goadesign/goa-cellar/swagger"
	"github.com/goadesign/logging/log15"
	"github.com/goadesign/middleware"
	"github.com/goadesign/middleware/security/basicauth"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {
	// Create goa service
	service := goa.New("cellar")

	// Setup logger
	logger := log15.New()
	service.UseLogger(goalog15.New(logger))

	// Setup basic middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.Recover())
	app.ConfigureAdminPassSecurity(service, basicauth.New("wine", "lover"))

	// Mount account controller onto service
	ac := controllers.NewAccount(service)
	app.MountAccountController(service, ac)

	// Mount bottle controller onto service
	bc := controllers.NewBottle(service)
	app.MountBottleController(service, bc)

	// Mount Swagger Spec controller onto service
	swagger.MountController(service)

	// Mount JSON Schema controller onto service
	schema.MountController(service)

	// Mount JavaScript example
	js.MountController(service)

	// Run service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.Error(err.Error())
	}
}
