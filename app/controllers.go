//************************************************************************//
// API "cellar": Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/goa-cellar
// --design=github.com/goadesign/goa-cellar/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

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
	service.Encoder(goa.NewJSONEncoder, "application/json")
	service.Encoder(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder(goa.NewXMLEncoder, "application/xml")
	service.Decoder(goa.NewJSONDecoder, "application/json")
	service.Decoder(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder(goa.NewJSONEncoder, "*/*")
	service.Decoder(goa.NewJSONDecoder, "*/*")
}

// AccountController is the controller interface for the Account actions.
type AccountController interface {
	goa.Muxer
	Create(*CreateAccountContext) error
	Delete(*DeleteAccountContext) error
	Show(*ShowAccountContext) error
	Update(*UpdateAccountContext) error
}

// MountAccountController "mounts" a Account resource controller on the given service.
func MountAccountController(service *goa.Service, ctrl AccountController) {
	initService(service)
	var h goa.Handler
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateAccountContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateAccountPayload)
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/cellar/accounts", ctrl.MuxHandler("Create", h, unmarshalCreateAccountPayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Account"}, goa.KV{"action", "Create"}, goa.KV{"route", "POST /cellar/accounts"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteAccountContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/cellar/accounts/:accountID", ctrl.MuxHandler("Delete", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Account"}, goa.KV{"action", "Delete"}, goa.KV{"route", "DELETE /cellar/accounts/:accountID"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowAccountContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/cellar/accounts/:accountID", ctrl.MuxHandler("Show", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Account"}, goa.KV{"action", "Show"}, goa.KV{"route", "GET /cellar/accounts/:accountID"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateAccountContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateAccountPayload)
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PUT", "/cellar/accounts/:accountID", ctrl.MuxHandler("Update", h, unmarshalUpdateAccountPayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Account"}, goa.KV{"action", "Update"}, goa.KV{"route", "PUT /cellar/accounts/:accountID"})
}

// unmarshalCreateAccountPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateAccountPayload(ctx context.Context, req *http.Request) error {
	var payload CreateAccountPayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}

// unmarshalUpdateAccountPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateAccountPayload(ctx context.Context, req *http.Request) error {
	var payload UpdateAccountPayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}

// BottleController is the controller interface for the Bottle actions.
type BottleController interface {
	goa.Muxer
	Create(*CreateBottleContext) error
	Delete(*DeleteBottleContext) error
	List(*ListBottleContext) error
	Rate(*RateBottleContext) error
	Show(*ShowBottleContext) error
	Update(*UpdateBottleContext) error
}

// MountBottleController "mounts" a Bottle resource controller on the given service.
func MountBottleController(service *goa.Service, ctrl BottleController) {
	initService(service)
	var h goa.Handler
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateBottleContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateBottlePayload)
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/cellar/accounts/:accountID/bottles", ctrl.MuxHandler("Create", h, unmarshalCreateBottlePayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Bottle"}, goa.KV{"action", "Create"}, goa.KV{"route", "POST /cellar/accounts/:accountID/bottles"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteBottleContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/cellar/accounts/:accountID/bottles/:bottleID", ctrl.MuxHandler("Delete", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Bottle"}, goa.KV{"action", "Delete"}, goa.KV{"route", "DELETE /cellar/accounts/:accountID/bottles/:bottleID"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListBottleContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/cellar/accounts/:accountID/bottles", ctrl.MuxHandler("List", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Bottle"}, goa.KV{"action", "List"}, goa.KV{"route", "GET /cellar/accounts/:accountID/bottles"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewRateBottleContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*RateBottlePayload)
		}
		return ctrl.Rate(rctx)
	}
	service.Mux.Handle("PUT", "/cellar/accounts/:accountID/bottles/:bottleID/actions/rate", ctrl.MuxHandler("Rate", h, unmarshalRateBottlePayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Bottle"}, goa.KV{"action", "Rate"}, goa.KV{"route", "PUT /cellar/accounts/:accountID/bottles/:bottleID/actions/rate"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowBottleContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/cellar/accounts/:accountID/bottles/:bottleID", ctrl.MuxHandler("Show", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Bottle"}, goa.KV{"action", "Show"}, goa.KV{"route", "GET /cellar/accounts/:accountID/bottles/:bottleID"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateBottleContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateBottlePayload)
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PATCH", "/cellar/accounts/:accountID/bottles/:bottleID", ctrl.MuxHandler("Update", h, unmarshalUpdateBottlePayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Bottle"}, goa.KV{"action", "Update"}, goa.KV{"route", "PATCH /cellar/accounts/:accountID/bottles/:bottleID"})
}

// unmarshalCreateBottlePayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateBottlePayload(ctx context.Context, req *http.Request) error {
	var payload CreateBottlePayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}

// unmarshalRateBottlePayload unmarshals the request body into the context request data Payload field.
func unmarshalRateBottlePayload(ctx context.Context, req *http.Request) error {
	var payload RateBottlePayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}

// unmarshalUpdateBottlePayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateBottlePayload(ctx context.Context, req *http.Request) error {
	var payload UpdateBottlePayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}
