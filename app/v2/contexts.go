//************************************************************************//
// API "cellar" version 2.0: Application Contexts
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
	"strconv"
	"strings"
)

// CreateGenericBottleContext provides the generic_bottle create action context.
type CreateGenericBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID  int
	Payload    *CreateGenericBottlePayload
	APIVersion string
}

// NewCreateGenericBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the generic_bottle controller create action.
func NewCreateGenericBottleContext(ctx context.Context) (*CreateGenericBottleContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := CreateGenericBottleContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	return &rctx, err
}

// CreateGenericBottlePayload is the generic_bottle create action payload.
type CreateGenericBottlePayload struct {
	Color   string  `json:"color" xml:"color"`
	Country *string `json:"country,omitempty" xml:"country,omitempty"`
	// Bottle kind
	Kind      string  `json:"kind" xml:"kind"`
	Name      string  `json:"name" xml:"name"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	Varietal  string  `json:"varietal" xml:"varietal"`
	Vineyard  string  `json:"vineyard" xml:"vineyard"`
	Vintage   int     `json:"vintage" xml:"vintage"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateGenericBottlePayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MissingAttributeError(`raw`, "name", err)
	}
	if payload.Kind == "" {
		err = goa.MissingAttributeError(`raw`, "kind", err)
	}
	if payload.Vineyard == "" {
		err = goa.MissingAttributeError(`raw`, "vineyard", err)
	}
	if payload.Varietal == "" {
		err = goa.MissingAttributeError(`raw`, "varietal", err)
	}

	if payload.Color == "" {
		err = goa.MissingAttributeError(`raw`, "color", err)
	}

	if !(payload.Color == "red" || payload.Color == "white" || payload.Color == "rose" || payload.Color == "yellow" || payload.Color == "sparkling") {
		err = goa.InvalidEnumValueError(`raw.color`, payload.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
	}
	if payload.Country != nil {
		if len(*payload.Country) < 2 {
			err = goa.InvalidLengthError(`raw.country`, *payload.Country, len(*payload.Country), 2, true, err)
		}
	}
	if !(payload.Kind == "wine" || payload.Kind == "rum") {
		err = goa.InvalidEnumValueError(`raw.kind`, payload.Kind, []interface{}{"wine", "rum"}, err)
	}
	if len(payload.Name) < 2 {
		err = goa.InvalidLengthError(`raw.name`, payload.Name, len(payload.Name), 2, true, err)
	}
	if payload.Review != nil {
		if len(*payload.Review) < 3 {
			err = goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 3, true, err)
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) > 300 {
			err = goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 300, false, err)
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness < 1 {
			err = goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 1, true, err)
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness > 5 {
			err = goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 5, false, err)
		}
	}
	if len(payload.Varietal) < 4 {
		err = goa.InvalidLengthError(`raw.varietal`, payload.Varietal, len(payload.Varietal), 4, true, err)
	}
	if len(payload.Vineyard) < 2 {
		err = goa.InvalidLengthError(`raw.vineyard`, payload.Vineyard, len(payload.Vineyard), 2, true, err)
	}
	if payload.Vintage < 1900 {
		err = goa.InvalidRangeError(`raw.vintage`, payload.Vintage, 1900, true, err)
	}
	if payload.Vintage > 2020 {
		err = goa.InvalidRangeError(`raw.vintage`, payload.Vintage, 2020, false, err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateGenericBottleContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// DeleteGenericBottleContext provides the generic_bottle delete action context.
type DeleteGenericBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID  int
	BottleID   int
	APIVersion string
}

// NewDeleteGenericBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the generic_bottle controller delete action.
func NewDeleteGenericBottleContext(ctx context.Context) (*DeleteGenericBottleContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := DeleteGenericBottleContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := req.Params.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.InvalidParamTypeError("bottleID", rawBottleID, "integer", err)
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteGenericBottleContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteGenericBottleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListGenericBottleContext provides the generic_bottle list action context.
type ListGenericBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID  int
	Years      []int
	APIVersion string
}

// NewListGenericBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the generic_bottle controller list action.
func NewListGenericBottleContext(ctx context.Context) (*ListGenericBottleContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ListGenericBottleContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawYears := req.Params.Get("years")
	if rawYears != "" {
		elemsYears := strings.Split(rawYears, ",")
		elemsYears2 := make([]int, len(elemsYears))
		for i, rawElem := range elemsYears {
			if elem, err2 := strconv.Atoi(rawElem); err2 == nil {
				elemsYears2[i] = elem
			} else {
				err = goa.InvalidParamTypeError("elem", rawElem, "integer", err)
			}
		}
		rctx.Years = elemsYears2
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListGenericBottleContext) OK(r GenericBottleCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.generic.bottle+json; type=collection")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ListGenericBottleContext) OKTiny(r GenericBottleTinyCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.generic.bottle+json; type=collection")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListGenericBottleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// RateGenericBottleContext provides the generic_bottle rate action context.
type RateGenericBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID  int
	BottleID   int
	Payload    *RateGenericBottlePayload
	APIVersion string
}

// NewRateGenericBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the generic_bottle controller rate action.
func NewRateGenericBottleContext(ctx context.Context) (*RateGenericBottleContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := RateGenericBottleContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := req.Params.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.InvalidParamTypeError("bottleID", rawBottleID, "integer", err)
		}
	}
	return &rctx, err
}

// RateGenericBottlePayload is the generic_bottle rate action payload.
type RateGenericBottlePayload struct {
	// Rating of bottle between 1 and 5
	Rating int `json:"rating" xml:"rating"`
}

// Validate runs the validation rules defined in the design.
func (payload *RateGenericBottlePayload) Validate() (err error) {

	if payload.Rating < 1 {
		err = goa.InvalidRangeError(`raw.rating`, payload.Rating, 1, true, err)
	}
	if payload.Rating > 5 {
		err = goa.InvalidRangeError(`raw.rating`, payload.Rating, 5, false, err)
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *RateGenericBottleContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *RateGenericBottleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowGenericBottleContext provides the generic_bottle show action context.
type ShowGenericBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID  int
	BottleID   int
	APIVersion string
}

// NewShowGenericBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the generic_bottle controller show action.
func NewShowGenericBottleContext(ctx context.Context) (*ShowGenericBottleContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ShowGenericBottleContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := req.Params.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.InvalidParamTypeError("bottleID", rawBottleID, "integer", err)
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowGenericBottleContext) OK(r *GenericBottle) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.generic.bottle")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *ShowGenericBottleContext) OKFull(r *GenericBottleFull) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.generic.bottle")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ShowGenericBottleContext) OKTiny(r *GenericBottleTiny) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.generic.bottle")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowGenericBottleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateGenericBottleContext provides the generic_bottle update action context.
type UpdateGenericBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID  int
	BottleID   int
	Payload    *UpdateGenericBottlePayload
	APIVersion string
}

// NewUpdateGenericBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the generic_bottle controller update action.
func NewUpdateGenericBottleContext(ctx context.Context) (*UpdateGenericBottleContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := UpdateGenericBottleContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := req.Params.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.InvalidParamTypeError("bottleID", rawBottleID, "integer", err)
		}
	}
	return &rctx, err
}

// UpdateGenericBottlePayload is the generic_bottle update action payload.
type UpdateGenericBottlePayload struct {
	Color   *string `json:"color,omitempty" xml:"color,omitempty"`
	Country *string `json:"country,omitempty" xml:"country,omitempty"`
	// Bottle kind
	Kind      *string `json:"kind,omitempty" xml:"kind,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	Varietal  *string `json:"varietal,omitempty" xml:"varietal,omitempty"`
	Vineyard  *string `json:"vineyard,omitempty" xml:"vineyard,omitempty"`
	Vintage   *int    `json:"vintage,omitempty" xml:"vintage,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateGenericBottlePayload) Validate() (err error) {
	if payload.Color != nil {
		if !(*payload.Color == "red" || *payload.Color == "white" || *payload.Color == "rose" || *payload.Color == "yellow" || *payload.Color == "sparkling") {
			err = goa.InvalidEnumValueError(`raw.color`, *payload.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
		}
	}
	if payload.Country != nil {
		if len(*payload.Country) < 2 {
			err = goa.InvalidLengthError(`raw.country`, *payload.Country, len(*payload.Country), 2, true, err)
		}
	}
	if payload.Kind != nil {
		if !(*payload.Kind == "wine" || *payload.Kind == "rum") {
			err = goa.InvalidEnumValueError(`raw.kind`, *payload.Kind, []interface{}{"wine", "rum"}, err)
		}
	}
	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true, err)
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) < 3 {
			err = goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 3, true, err)
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) > 300 {
			err = goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 300, false, err)
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness < 1 {
			err = goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 1, true, err)
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness > 5 {
			err = goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 5, false, err)
		}
	}
	if payload.Varietal != nil {
		if len(*payload.Varietal) < 4 {
			err = goa.InvalidLengthError(`raw.varietal`, *payload.Varietal, len(*payload.Varietal), 4, true, err)
		}
	}
	if payload.Vineyard != nil {
		if len(*payload.Vineyard) < 2 {
			err = goa.InvalidLengthError(`raw.vineyard`, *payload.Vineyard, len(*payload.Vineyard), 2, true, err)
		}
	}
	if payload.Vintage != nil {
		if *payload.Vintage < 1900 {
			err = goa.InvalidRangeError(`raw.vintage`, *payload.Vintage, 1900, true, err)
		}
	}
	if payload.Vintage != nil {
		if *payload.Vintage > 2020 {
			err = goa.InvalidRangeError(`raw.vintage`, *payload.Vintage, 2020, false, err)
		}
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateGenericBottleContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateGenericBottleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
