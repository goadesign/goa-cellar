// +build appengine

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	sg "google.golang.org/api/storage/v1"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"

	"appengine"

	"github.com/raphael/goa"
	"github.com/raphael/goa-middleware"
	"github.com/raphael/goa/cors"
	"github.com/raphael/goa/examples/cellar/app"
	"github.com/raphael/goa/examples/cellar/controllers"
	"github.com/raphael/goa/examples/cellar/swagger"
	"gopkg.in/inconshreveable/log15.v2"
)

const (
	projectID  = "goa-cellar"
	bucketName = "artifacts.cellar.goa.design"
	scope      = sg.DevstorageReadOnlyScope
)

var (
	// bucket that stores DP key
	bucket *storage.BucketHandle

	// DP key
	dpKey string
)

func init() {
	// Configure logging for appengine
	goa.Log.SetHandler(log15.MultiHandler(
		log15.StreamHandler(os.Stderr, log15.LogfmtFormat()),
		AppEngineLogHandler()),
	)

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

	// Load DP key
	err = loadDPKey()
	fmt.Fprintf(os.Stderr, "DPKEY: %s", dpKey)
	if err != nil {
		panic(err)
	}

	// Setup basic middleware
	service.Use(goa.RequestID())
	service.Use(AppEngineLogCtx())
	service.Use(cors.Middleware(spec))
	service.Use(middleware.DeferPanic(service, dpKey))

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

// Load DP key
func loadDPKey() error {
	defClient, err := google.DefaultClient(context.Background(), scope)
	if err != nil {
		return fmt.Errorf("Unable to get default client: %v", err)
	}
	ctx := cloud.NewContext(projectID, defClient)
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("Unable to get storage client: %v", err)
	}
	bucket = client.Bucket(bucketName)
	object := bucket.Object(".dpkey")
	rc, err := object.NewReader(ctx)
	if err != nil {
		return fmt.Errorf("Unable to read DP key: %v", err)
	}
	defer rc.Close()
	b, err := ioutil.ReadAll(rc)
	if err != nil {
		return fmt.Errorf("Unable to load DP key: %v", err)
	}
	dpKey = string(b)
	return nil
}

// AppEngineLogHandler sends logs to AppEngine.
// The record must contain the appengine request context.
func AppEngineLogHandler() log15.Handler {
	logFormat := log15.JsonFormat()
	return log15.FuncHandler(func(r *log15.Record) error {
		var c appengine.Context
		index := 0
		for i, e := range r.Ctx {
			if ct, ok := e.(appengine.Context); ok {
				c = ct
				index = i
				break
			}
		}
		if c == nil {
			// not in the context of a request
			return nil
		}
		r.Ctx = append(r.Ctx[:index-1], r.Ctx[index+1:]...)
		log := string(logFormat.Format(r))
		switch r.Lvl {
		case log15.LvlCrit:
			c.Criticalf(log)
		case log15.LvlError:
			c.Errorf(log)
		case log15.LvlWarn:
			c.Warningf(log)
		case log15.LvlInfo:
			c.Infof(log)
		case log15.LvlDebug:
			c.Debugf(log)
		}
		return nil
	})
}

// AppEngineLogCtx returns a goa middleware that sets the appengine context in the log records.
func AppEngineLogCtx() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx *goa.Context) error {
			actx := appengine.NewContext(ctx.Request())
			ctx.SetValue(goa.ReqIDKey, appengine.RequestID(actx))
			ctx.Logger = ctx.Logger.New("aeCtx", actx)
			return h(ctx)
		}
	}
}
