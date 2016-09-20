package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/context"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar-ep/app"
	"github.com/goadesign/goa-cellar-ep/controllers"
	"github.com/goadesign/goa-cellar-ep/store"
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

	// Setup security middleware
	epMiddleware := endpointsMiddleware()
	app.UseAPIKeyMiddleware(service, epMiddleware)
	app.UseJWTMiddleware(service, epMiddleware)

	// Instantiate DB
	db := store.NewDB()

	// Mount account controller onto service
	ac := controllers.NewAccount(service, db)
	app.MountAccountController(service, ac)

	// Mount bottle controller onto service
	bc := controllers.NewBottle(service, db)
	app.MountBottleController(service, bc)

	// Mount auth controller onto service
	uc := controllers.NewAuth(service)
	app.MountAuthController(service, uc)

	// Mount health-check controller onto service
	hc := controllers.NewHealth(service, db)
	app.MountHealthController(service, hc)

	// Mount public controller onto service
	pc := controllers.NewPublic(service)
	app.MountPublicController(service, pc)

	// Mount js controller onto service
	jc := controllers.NewJs(service)
	app.MountJsController(service, jc)

	// Mount swagger controller onto service
	sc := controllers.NewSwagger(service)
	app.MountSwaggerController(service, sc)

	// Run service
	port := "8080"
	if s := os.Getenv("PORT"); s != "" {
		port = s
	}
	if err := service.ListenAndServe(":" + port); err != nil {
		service.LogError(err.Error())
	}

	service.LogInfo("exiting")
}

// endpointsMiddleware extracts the user information initialized by Google Cloud Endpoints
func endpointsMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			info := req.Header.Get("X-Endpoint-API-UserInfo")
			var res map[string]interface{}
			if info == "" {
				res = map[string]interface{}{"id": "anonymous"}
			} else {
				js, err := base64.StdEncoding.DecodeString(info)
				if err != nil {
					return fmt.Errorf("X-Endpoint-API-UserInfo contains invalid base64 encoding: %s", err)
				}
				err = json.Unmarshal(js, &res)
				if err != nil {
					return fmt.Errorf("X-Endpoint-API-UserInfo contains invalid JSON: %s", err)
				}
			}
			ctx = context.WithValue(ctx, "userinfo", res)
			return h(ctx, rw, req)
		}
	}
}
