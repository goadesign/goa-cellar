//************************************************************************//
// API "cellar" version 2.0: Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/goa-cellar
// --design=github.com/goadesign/goa-cellar/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package v2

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// inited is true if initService has been called
var inited = false

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	if inited {
		return
	}
	inited = true
	// Setup encoders and decoders
	service.Version("2.0").SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.Version("2.0").SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.Version("2.0").SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.Version("2.0").SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.Version("2.0").SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.Version("2.0").SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")
}

// GenericBottleController is the controller interface for the GenericBottle actions.
type GenericBottleController interface {
	goa.Muxer
	Create(*CreateGenericBottleContext) error
	Delete(*DeleteGenericBottleContext) error
	List(*ListGenericBottleContext) error
	Rate(*RateGenericBottleContext) error
	Show(*ShowGenericBottleContext) error
	Update(*UpdateGenericBottleContext) error
}

// MountGenericBottleController "mounts" a GenericBottle resource controller on the given service.
func MountGenericBottleController(service *goa.Service, ctrl GenericBottleController) {
	initService(service)
	var h goa.Handler
	mux := service.Version("2.0").Mux
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateGenericBottleContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		rctx.APIVersion = "2.0"
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateGenericBottlePayload)
		}
		return ctrl.Create(rctx)
	}
	mux.Handle("POST", "/cellar/accounts/:accountID/bottles", ctrl.MuxHandler("Create", "2.0", h, unmarshalCreateGenericBottlePayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "GenericBottle"}, goa.KV{"version", "2.0"}, goa.KV{"action", "Create"}, goa.KV{"route", "POST /cellar/accounts/:accountID/bottles"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteGenericBottleContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		rctx.APIVersion = "2.0"
		return ctrl.Delete(rctx)
	}
	mux.Handle("DELETE", "/cellar/accounts/:accountID/bottles/:bottleID", ctrl.MuxHandler("Delete", "2.0", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "GenericBottle"}, goa.KV{"version", "2.0"}, goa.KV{"action", "Delete"}, goa.KV{"route", "DELETE /cellar/accounts/:accountID/bottles/:bottleID"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListGenericBottleContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		rctx.APIVersion = "2.0"
		return ctrl.List(rctx)
	}
	mux.Handle("GET", "/cellar/accounts/:accountID/bottles", ctrl.MuxHandler("List", "2.0", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "GenericBottle"}, goa.KV{"version", "2.0"}, goa.KV{"action", "List"}, goa.KV{"route", "GET /cellar/accounts/:accountID/bottles"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewRateGenericBottleContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		rctx.APIVersion = "2.0"
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*RateGenericBottlePayload)
		}
		return ctrl.Rate(rctx)
	}
	mux.Handle("PUT", "/cellar/accounts/:accountID/bottles/:bottleID/actions/rate", ctrl.MuxHandler("Rate", "2.0", h, unmarshalRateGenericBottlePayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "GenericBottle"}, goa.KV{"version", "2.0"}, goa.KV{"action", "Rate"}, goa.KV{"route", "PUT /cellar/accounts/:accountID/bottles/:bottleID/actions/rate"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowGenericBottleContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		rctx.APIVersion = "2.0"
		return ctrl.Show(rctx)
	}
	mux.Handle("GET", "/cellar/accounts/:accountID/bottles/:bottleID", ctrl.MuxHandler("Show", "2.0", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "GenericBottle"}, goa.KV{"version", "2.0"}, goa.KV{"action", "Show"}, goa.KV{"route", "GET /cellar/accounts/:accountID/bottles/:bottleID"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateGenericBottleContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		rctx.APIVersion = "2.0"
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateGenericBottlePayload)
		}
		return ctrl.Update(rctx)
	}
	mux.Handle("PATCH", "/cellar/accounts/:accountID/bottles/:bottleID", ctrl.MuxHandler("Update", "2.0", h, unmarshalUpdateGenericBottlePayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "GenericBottle"}, goa.KV{"version", "2.0"}, goa.KV{"action", "Update"}, goa.KV{"route", "PATCH /cellar/accounts/:accountID/bottles/:bottleID"})
}

// unmarshalCreateGenericBottlePayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateGenericBottlePayload(ctx context.Context, req *http.Request) error {
	var payload CreateGenericBottlePayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}

// unmarshalRateGenericBottlePayload unmarshals the request body into the context request data Payload field.
func unmarshalRateGenericBottlePayload(ctx context.Context, req *http.Request) error {
	var payload RateGenericBottlePayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}

// unmarshalUpdateGenericBottlePayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateGenericBottlePayload(ctx context.Context, req *http.Request) error {
	var payload UpdateGenericBottlePayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}
