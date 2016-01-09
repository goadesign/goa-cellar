//************************************************************************//
// API "cellar" version 1.0: Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/raphael/goa-cellar
// --design=github.com/raphael/testd/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package v1

import "github.com/raphael/goa"

// BottleController is the controller interface for the Bottle actions.
type BottleController interface {
	goa.Controller
	Create(*CreateBottleContext) error
	Delete(*DeleteBottleContext) error
	List(*ListBottleContext) error
	Rate(*RateBottleContext) error
	Show(*ShowBottleContext) error
	Update(*UpdateBottleContext) error
}

// MountBottleController "mounts" a Bottle resource controller on the given service.
func MountBottleController(service goa.Service, ctrl BottleController) {
	var h goa.Handler
	mux := service.ServeMux().Version("1.0")
	h = func(c *goa.Context) error {
		ctx, err := NewCreateBottleContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Create(ctx)
	}
	mux.Handle("POST", "/1.0/cellar/accounts/:accountID/bottles", ctrl.HandleFunc("Create", h))
	service.Info("mount", "ctrl", "Bottle", "version", "1.0", "action", "Create", "route", "POST /1.0/cellar/accounts/:accountID/bottles")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteBottleContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	mux.Handle("DELETE", "/1.0/cellar/accounts/:accountID/bottles/:bottleID", ctrl.HandleFunc("Delete", h))
	service.Info("mount", "ctrl", "Bottle", "version", "1.0", "action", "Delete", "route", "DELETE /1.0/cellar/accounts/:accountID/bottles/:bottleID")
	h = func(c *goa.Context) error {
		ctx, err := NewListBottleContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	mux.Handle("GET", "/1.0/cellar/accounts/:accountID/bottles", ctrl.HandleFunc("List", h))
	service.Info("mount", "ctrl", "Bottle", "version", "1.0", "action", "List", "route", "GET /1.0/cellar/accounts/:accountID/bottles")
	h = func(c *goa.Context) error {
		ctx, err := NewRateBottleContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Rate(ctx)
	}
	mux.Handle("PUT", "/1.0/cellar/accounts/:accountID/bottles/:bottleID/actions/rate", ctrl.HandleFunc("Rate", h))
	service.Info("mount", "ctrl", "Bottle", "version", "1.0", "action", "Rate", "route", "PUT /1.0/cellar/accounts/:accountID/bottles/:bottleID/actions/rate")
	h = func(c *goa.Context) error {
		ctx, err := NewShowBottleContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	mux.Handle("GET", "/1.0/cellar/accounts/:accountID/bottles/:bottleID", ctrl.HandleFunc("Show", h))
	service.Info("mount", "ctrl", "Bottle", "version", "1.0", "action", "Show", "route", "GET /1.0/cellar/accounts/:accountID/bottles/:bottleID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateBottleContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	mux.Handle("PATCH", "/1.0/cellar/accounts/:accountID/bottles/:bottleID", ctrl.HandleFunc("Update", h))
	service.Info("mount", "ctrl", "Bottle", "version", "1.0", "action", "Update", "route", "PATCH /1.0/cellar/accounts/:accountID/bottles/:bottleID")
}
