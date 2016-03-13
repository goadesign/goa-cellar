//************************************************************************//
// API "cellar": Application Contexts
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
	"strconv"
	"strings"
)

// CreateAccountContext provides the account create action context.
type CreateAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateAccountPayload
}

// NewCreateAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller create action.
func NewCreateAccountContext(ctx context.Context) (*CreateAccountContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := CreateAccountContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
}

// CreateAccountPayload is the account create action payload.
type CreateAccountPayload struct {
	// Name of account
	Name string `json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateAccountPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.StackErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateAccountContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// DeleteAccountContext provides the account delete action context.
type DeleteAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID int
}

// NewDeleteAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller delete action.
func NewDeleteAccountContext(ctx context.Context) (*DeleteAccountContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := DeleteAccountContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteAccountContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteAccountContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowAccountContext provides the account show action context.
type ShowAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID int
}

// NewShowAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller show action.
func NewShowAccountContext(ctx context.Context) (*ShowAccountContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ShowAccountContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowAccountContext) OK(r *Account) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.account")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ShowAccountContext) OKTiny(r *AccountTiny) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.account")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowAccountContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateAccountContext provides the account update action context.
type UpdateAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID int
	Payload   *UpdateAccountPayload
}

// NewUpdateAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller update action.
func NewUpdateAccountContext(ctx context.Context) (*UpdateAccountContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := UpdateAccountContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	return &rctx, err
}

// UpdateAccountPayload is the account update action payload.
type UpdateAccountPayload struct {
	// Name of account
	Name string `json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateAccountPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.StackErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateAccountContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateAccountContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateBottleContext provides the bottle create action context.
type CreateBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID int
	Payload   *CreateBottlePayload
}

// NewCreateBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller create action.
func NewCreateBottleContext(ctx context.Context) (*CreateBottleContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := CreateBottleContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	return &rctx, err
}

// CreateBottlePayload is the bottle create action payload.
type CreateBottlePayload struct {
	Color     string  `json:"color" xml:"color"`
	Country   *string `json:"country,omitempty" xml:"country,omitempty"`
	Name      string  `json:"name" xml:"name"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	Varietal  string  `json:"varietal" xml:"varietal"`
	Vineyard  string  `json:"vineyard" xml:"vineyard"`
	Vintage   int     `json:"vintage" xml:"vintage"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateBottlePayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.StackErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Vineyard == "" {
		err = goa.StackErrors(err, goa.MissingAttributeError(`raw`, "vineyard"))
	}
	if payload.Varietal == "" {
		err = goa.StackErrors(err, goa.MissingAttributeError(`raw`, "varietal"))
	}

	if payload.Color == "" {
		err = goa.StackErrors(err, goa.MissingAttributeError(`raw`, "color"))
	}

	if !(payload.Color == "red" || payload.Color == "white" || payload.Color == "rose" || payload.Color == "yellow" || payload.Color == "sparkling") {
		err = goa.StackErrors(err, goa.InvalidEnumValueError(`raw.color`, payload.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}))
	}
	if payload.Country != nil {
		if len(*payload.Country) < 2 {
			err = goa.StackErrors(err, goa.InvalidLengthError(`raw.country`, *payload.Country, len(*payload.Country), 2, true))
		}
	}
	if len(payload.Name) < 2 {
		err = goa.StackErrors(err, goa.InvalidLengthError(`raw.name`, payload.Name, len(payload.Name), 2, true))
	}
	if payload.Review != nil {
		if len(*payload.Review) < 3 {
			err = goa.StackErrors(err, goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 3, true))
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) > 300 {
			err = goa.StackErrors(err, goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 300, false))
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness < 1 {
			err = goa.StackErrors(err, goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 1, true))
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness > 5 {
			err = goa.StackErrors(err, goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 5, false))
		}
	}
	if len(payload.Varietal) < 4 {
		err = goa.StackErrors(err, goa.InvalidLengthError(`raw.varietal`, payload.Varietal, len(payload.Varietal), 4, true))
	}
	if len(payload.Vineyard) < 2 {
		err = goa.StackErrors(err, goa.InvalidLengthError(`raw.vineyard`, payload.Vineyard, len(payload.Vineyard), 2, true))
	}
	if payload.Vintage < 1900 {
		err = goa.StackErrors(err, goa.InvalidRangeError(`raw.vintage`, payload.Vintage, 1900, true))
	}
	if payload.Vintage > 2020 {
		err = goa.StackErrors(err, goa.InvalidRangeError(`raw.vintage`, payload.Vintage, 2020, false))
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateBottleContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// DeleteBottleContext provides the bottle delete action context.
type DeleteBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID int
	BottleID  int
}

// NewDeleteBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller delete action.
func NewDeleteBottleContext(ctx context.Context) (*DeleteBottleContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := DeleteBottleContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	rawBottleID := req.Params.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("bottleID", rawBottleID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteBottleContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteBottleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListBottleContext provides the bottle list action context.
type ListBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID int
	Years     []int
}

// NewListBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller list action.
func NewListBottleContext(ctx context.Context) (*ListBottleContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ListBottleContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
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
				err = goa.StackErrors(err, goa.InvalidParamTypeError("elem", rawElem, "integer"))
			}
		}
		rctx.Years = elemsYears2
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListBottleContext) OK(r BottleCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bottle+json; type=collection")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ListBottleContext) OKTiny(r BottleTinyCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bottle+json; type=collection")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListBottleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// RateBottleContext provides the bottle rate action context.
type RateBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID int
	BottleID  int
	Payload   *RateBottlePayload
}

// NewRateBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller rate action.
func NewRateBottleContext(ctx context.Context) (*RateBottleContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := RateBottleContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	rawBottleID := req.Params.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("bottleID", rawBottleID, "integer"))
		}
	}
	return &rctx, err
}

// RateBottlePayload is the bottle rate action payload.
type RateBottlePayload struct {
	// Rating of bottle between 1 and 5
	Rating int `json:"rating" xml:"rating"`
}

// Validate runs the validation rules defined in the design.
func (payload *RateBottlePayload) Validate() (err error) {

	if payload.Rating < 1 {
		err = goa.StackErrors(err, goa.InvalidRangeError(`raw.rating`, payload.Rating, 1, true))
	}
	if payload.Rating > 5 {
		err = goa.StackErrors(err, goa.InvalidRangeError(`raw.rating`, payload.Rating, 5, false))
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *RateBottleContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *RateBottleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowBottleContext provides the bottle show action context.
type ShowBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID int
	BottleID  int
}

// NewShowBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller show action.
func NewShowBottleContext(ctx context.Context) (*ShowBottleContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ShowBottleContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	rawBottleID := req.Params.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("bottleID", rawBottleID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowBottleContext) OK(r *Bottle) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bottle")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *ShowBottleContext) OKFull(r *BottleFull) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bottle")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ShowBottleContext) OKTiny(r *BottleTiny) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bottle")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowBottleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateBottleContext provides the bottle update action context.
type UpdateBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID int
	BottleID  int
	Payload   *UpdateBottlePayload
}

// NewUpdateBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller update action.
func NewUpdateBottleContext(ctx context.Context) (*UpdateBottleContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := UpdateBottleContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	rawBottleID := req.Params.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("bottleID", rawBottleID, "integer"))
		}
	}
	return &rctx, err
}

// UpdateBottlePayload is the bottle update action payload.
type UpdateBottlePayload struct {
	Color     *string `json:"color,omitempty" xml:"color,omitempty"`
	Country   *string `json:"country,omitempty" xml:"country,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	Varietal  *string `json:"varietal,omitempty" xml:"varietal,omitempty"`
	Vineyard  *string `json:"vineyard,omitempty" xml:"vineyard,omitempty"`
	Vintage   *int    `json:"vintage,omitempty" xml:"vintage,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateBottlePayload) Validate() (err error) {
	if payload.Color != nil {
		if !(*payload.Color == "red" || *payload.Color == "white" || *payload.Color == "rose" || *payload.Color == "yellow" || *payload.Color == "sparkling") {
			err = goa.StackErrors(err, goa.InvalidEnumValueError(`raw.color`, *payload.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}))
		}
	}
	if payload.Country != nil {
		if len(*payload.Country) < 2 {
			err = goa.StackErrors(err, goa.InvalidLengthError(`raw.country`, *payload.Country, len(*payload.Country), 2, true))
		}
	}
	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.StackErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true))
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) < 3 {
			err = goa.StackErrors(err, goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 3, true))
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) > 300 {
			err = goa.StackErrors(err, goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 300, false))
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness < 1 {
			err = goa.StackErrors(err, goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 1, true))
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness > 5 {
			err = goa.StackErrors(err, goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 5, false))
		}
	}
	if payload.Varietal != nil {
		if len(*payload.Varietal) < 4 {
			err = goa.StackErrors(err, goa.InvalidLengthError(`raw.varietal`, *payload.Varietal, len(*payload.Varietal), 4, true))
		}
	}
	if payload.Vineyard != nil {
		if len(*payload.Vineyard) < 2 {
			err = goa.StackErrors(err, goa.InvalidLengthError(`raw.vineyard`, *payload.Vineyard, len(*payload.Vineyard), 2, true))
		}
	}
	if payload.Vintage != nil {
		if *payload.Vintage < 1900 {
			err = goa.StackErrors(err, goa.InvalidRangeError(`raw.vintage`, *payload.Vintage, 1900, true))
		}
	}
	if payload.Vintage != nil {
		if *payload.Vintage > 2020 {
			err = goa.StackErrors(err, goa.InvalidRangeError(`raw.vintage`, *payload.Vintage, 2020, false))
		}
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateBottleContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateBottleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// WatchBottleContext provides the bottle watch action context.
type WatchBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountID int
	BottleID  int
}

// NewWatchBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller watch action.
func NewWatchBottleContext(ctx context.Context) (*WatchBottleContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := WatchBottleContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawAccountID := req.Params.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	rawBottleID := req.Params.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("bottleID", rawBottleID, "integer"))
		}
	}
	return &rctx, err
}
