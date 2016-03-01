// +build !appengine

package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar/app"
	"github.com/goadesign/goa-cellar/controllers"
	"github.com/goadesign/goa-cellar/js"
	"github.com/goadesign/goa-cellar/schema"
	"github.com/goadesign/goa-cellar/swagger"
	"github.com/goadesign/logging/logrus"
	"github.com/goadesign/middleware"
)

func main() {
	// Setup logger
	logger := logrus.New()
	goa.Log = goalogrus.New(logger)

	// Create goa service
	service := goa.New("cellar")

	// Setup basic middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.Recover())

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
		goa.Error(goa.RootContext, err.Error())
	}
}
