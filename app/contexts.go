//************************************************************************//
// API "cellar": Application Contexts
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/goadesign/goa-cellar/design
// --out=$(GOPATH)/src/github.com/goadesign/goa-cellar
// --version=v0.2.dev
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
func NewCreateAccountContext(ctx context.Context, service *goa.Service) (*CreateAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createAccountPayload is the account create action payload.
type createAccountPayload struct {
	// Name of account
	Name *string `json:"name,omitempty" xml:"name,omitempty" form:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createAccountPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	return
}

// Publicize creates CreateAccountPayload from createAccountPayload
func (payload *createAccountPayload) Publicize() *CreateAccountPayload {
	var pub CreateAccountPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// CreateAccountPayload is the account create action payload.
type CreateAccountPayload struct {
	// Name of account
	Name string `json:"name" xml:"name" form:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateAccountPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
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
func NewDeleteAccountContext(ctx context.Context, service *goa.Service) (*DeleteAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeleteAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountID := req.Params["accountID"]
	if len(paramAccountID) > 0 {
		rawAccountID := paramAccountID[0]
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
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

// ListAccountContext provides the account list action context.
type ListAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller list action.
func NewListAccountContext(ctx context.Context, service *goa.Service) (*ListAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListAccountContext) OK(r AccountCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.account+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ListAccountContext) OKTiny(r AccountTinyCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.account+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListAccountContext) NotFound() error {
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
func NewShowAccountContext(ctx context.Context, service *goa.Service) (*ShowAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountID := req.Params["accountID"]
	if len(paramAccountID) > 0 {
		rawAccountID := paramAccountID[0]
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowAccountContext) OK(r *Account) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.account+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ShowAccountContext) OKTiny(r *AccountTiny) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.account+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
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
func NewUpdateAccountContext(ctx context.Context, service *goa.Service) (*UpdateAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpdateAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountID := req.Params["accountID"]
	if len(paramAccountID) > 0 {
		rawAccountID := paramAccountID[0]
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	return &rctx, err
}

// updateAccountPayload is the account update action payload.
type updateAccountPayload struct {
	// Name of account
	Name *string `json:"name,omitempty" xml:"name,omitempty" form:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateAccountPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	return
}

// Publicize creates UpdateAccountPayload from updateAccountPayload
func (payload *updateAccountPayload) Publicize() *UpdateAccountPayload {
	var pub UpdateAccountPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// UpdateAccountPayload is the account update action payload.
type UpdateAccountPayload struct {
	// Name of account
	Name string `json:"name" xml:"name" form:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateAccountPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
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
func NewCreateBottleContext(ctx context.Context, service *goa.Service) (*CreateBottleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateBottleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountID := req.Params["accountID"]
	if len(paramAccountID) > 0 {
		rawAccountID := paramAccountID[0]
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	return &rctx, err
}

// createBottlePayload is the bottle create action payload.
type createBottlePayload struct {
	Color     *string `json:"color,omitempty" xml:"color,omitempty" form:"color,omitempty"`
	Country   *string `json:"country,omitempty" xml:"country,omitempty" form:"country,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty" form:"name,omitempty"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty" form:"region,omitempty"`
	Review    *string `json:"review,omitempty" xml:"review,omitempty" form:"review,omitempty"`
	Sweetness *int    `json:"sweetness,omitempty" xml:"sweetness,omitempty" form:"sweetness,omitempty"`
	Varietal  *string `json:"varietal,omitempty" xml:"varietal,omitempty" form:"varietal,omitempty"`
	Vineyard  *string `json:"vineyard,omitempty" xml:"vineyard,omitempty" form:"vineyard,omitempty"`
	Vintage   *int    `json:"vintage,omitempty" xml:"vintage,omitempty" form:"vintage,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createBottlePayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Vineyard == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "vineyard"))
	}
	if payload.Varietal == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "varietal"))
	}
	if payload.Vintage == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "vintage"))
	}
	if payload.Color == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "color"))
	}

	if payload.Color != nil {
		if !(*payload.Color == "red" || *payload.Color == "white" || *payload.Color == "rose" || *payload.Color == "yellow" || *payload.Color == "sparkling") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`raw.color`, *payload.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}))
		}
	}
	if payload.Country != nil {
		if len(*payload.Country) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.country`, *payload.Country, len(*payload.Country), 2, true))
		}
	}
	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true))
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 3, true))
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) > 300 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 300, false))
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 1, true))
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 5, false))
		}
	}
	if payload.Varietal != nil {
		if len(*payload.Varietal) < 4 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.varietal`, *payload.Varietal, len(*payload.Varietal), 4, true))
		}
	}
	if payload.Vineyard != nil {
		if len(*payload.Vineyard) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.vineyard`, *payload.Vineyard, len(*payload.Vineyard), 2, true))
		}
	}
	if payload.Vintage != nil {
		if *payload.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.vintage`, *payload.Vintage, 1900, true))
		}
	}
	if payload.Vintage != nil {
		if *payload.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.vintage`, *payload.Vintage, 2020, false))
		}
	}
	return
}

// Publicize creates CreateBottlePayload from createBottlePayload
func (payload *createBottlePayload) Publicize() *CreateBottlePayload {
	var pub CreateBottlePayload
	if payload.Color != nil {
		pub.Color = *payload.Color
	}
	if payload.Country != nil {
		pub.Country = payload.Country
	}
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	if payload.Region != nil {
		pub.Region = payload.Region
	}
	if payload.Review != nil {
		pub.Review = payload.Review
	}
	if payload.Sweetness != nil {
		pub.Sweetness = payload.Sweetness
	}
	if payload.Varietal != nil {
		pub.Varietal = *payload.Varietal
	}
	if payload.Vineyard != nil {
		pub.Vineyard = *payload.Vineyard
	}
	if payload.Vintage != nil {
		pub.Vintage = *payload.Vintage
	}
	return &pub
}

// CreateBottlePayload is the bottle create action payload.
type CreateBottlePayload struct {
	Color     string  `json:"color" xml:"color" form:"color"`
	Country   *string `json:"country,omitempty" xml:"country,omitempty" form:"country,omitempty"`
	Name      string  `json:"name" xml:"name" form:"name"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty" form:"region,omitempty"`
	Review    *string `json:"review,omitempty" xml:"review,omitempty" form:"review,omitempty"`
	Sweetness *int    `json:"sweetness,omitempty" xml:"sweetness,omitempty" form:"sweetness,omitempty"`
	Varietal  string  `json:"varietal" xml:"varietal" form:"varietal"`
	Vineyard  string  `json:"vineyard" xml:"vineyard" form:"vineyard"`
	Vintage   int     `json:"vintage" xml:"vintage" form:"vintage"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateBottlePayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Vineyard == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "vineyard"))
	}
	if payload.Varietal == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "varietal"))
	}
	if payload.Color == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "color"))
	}

	if !(payload.Color == "red" || payload.Color == "white" || payload.Color == "rose" || payload.Color == "yellow" || payload.Color == "sparkling") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`raw.color`, payload.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}))
	}
	if payload.Country != nil {
		if len(*payload.Country) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.country`, *payload.Country, len(*payload.Country), 2, true))
		}
	}
	if len(payload.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, payload.Name, len(payload.Name), 2, true))
	}
	if payload.Review != nil {
		if len(*payload.Review) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 3, true))
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) > 300 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 300, false))
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 1, true))
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 5, false))
		}
	}
	if len(payload.Varietal) < 4 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.varietal`, payload.Varietal, len(payload.Varietal), 4, true))
	}
	if len(payload.Vineyard) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.vineyard`, payload.Vineyard, len(payload.Vineyard), 2, true))
	}
	if payload.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.vintage`, payload.Vintage, 1900, true))
	}
	if payload.Vintage > 2020 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.vintage`, payload.Vintage, 2020, false))
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateBottleContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CreateBottleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
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
func NewDeleteBottleContext(ctx context.Context, service *goa.Service) (*DeleteBottleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeleteBottleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountID := req.Params["accountID"]
	if len(paramAccountID) > 0 {
		rawAccountID := paramAccountID[0]
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	paramBottleID := req.Params["bottleID"]
	if len(paramBottleID) > 0 {
		rawBottleID := paramBottleID[0]
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("bottleID", rawBottleID, "integer"))
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
func NewListBottleContext(ctx context.Context, service *goa.Service) (*ListBottleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListBottleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountID := req.Params["accountID"]
	if len(paramAccountID) > 0 {
		rawAccountID := paramAccountID[0]
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	paramYears := req.Params["years"]
	if len(paramYears) > 0 {
		var params []int
		for _, rawYears := range paramYears {
			elemsYears := strings.Split(rawYears, ",")
			elemsYears2 := make([]int, len(elemsYears))
			for i, rawElem := range elemsYears {
				if elem, err2 := strconv.Atoi(rawElem); err2 == nil {
					elemsYears2[i] = elem
				} else {
					err = goa.MergeErrors(err, goa.InvalidParamTypeError("elem", rawElem, "integer"))
				}
			}
			params = elemsYears2
			rctx.Years = append(rctx.Years, params...)
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListBottleContext) OK(r BottleCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bottle+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ListBottleContext) OKTiny(r BottleTinyCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bottle+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
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
func NewRateBottleContext(ctx context.Context, service *goa.Service) (*RateBottleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := RateBottleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountID := req.Params["accountID"]
	if len(paramAccountID) > 0 {
		rawAccountID := paramAccountID[0]
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	paramBottleID := req.Params["bottleID"]
	if len(paramBottleID) > 0 {
		rawBottleID := paramBottleID[0]
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("bottleID", rawBottleID, "integer"))
		}
	}
	return &rctx, err
}

// rateBottlePayload is the bottle rate action payload.
type rateBottlePayload struct {
	// Rating of bottle between 1 and 5
	Rating *int `json:"rating,omitempty" xml:"rating,omitempty" form:"rating,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *rateBottlePayload) Validate() (err error) {
	if payload.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "rating"))
	}

	if payload.Rating != nil {
		if *payload.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 1, true))
		}
	}
	if payload.Rating != nil {
		if *payload.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 5, false))
		}
	}
	return
}

// Publicize creates RateBottlePayload from rateBottlePayload
func (payload *rateBottlePayload) Publicize() *RateBottlePayload {
	var pub RateBottlePayload
	if payload.Rating != nil {
		pub.Rating = *payload.Rating
	}
	return &pub
}

// RateBottlePayload is the bottle rate action payload.
type RateBottlePayload struct {
	// Rating of bottle between 1 and 5
	Rating int `json:"rating" xml:"rating" form:"rating"`
}

// Validate runs the validation rules defined in the design.
func (payload *RateBottlePayload) Validate() (err error) {
	if payload.Rating < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, payload.Rating, 1, true))
	}
	if payload.Rating > 5 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, payload.Rating, 5, false))
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
func NewShowBottleContext(ctx context.Context, service *goa.Service) (*ShowBottleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowBottleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountID := req.Params["accountID"]
	if len(paramAccountID) > 0 {
		rawAccountID := paramAccountID[0]
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	paramBottleID := req.Params["bottleID"]
	if len(paramBottleID) > 0 {
		rawBottleID := paramBottleID[0]
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("bottleID", rawBottleID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowBottleContext) OK(r *Bottle) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bottle+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *ShowBottleContext) OKFull(r *BottleFull) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bottle+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ShowBottleContext) OKTiny(r *BottleTiny) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.bottle+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
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
func NewUpdateBottleContext(ctx context.Context, service *goa.Service) (*UpdateBottleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpdateBottleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountID := req.Params["accountID"]
	if len(paramAccountID) > 0 {
		rawAccountID := paramAccountID[0]
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	paramBottleID := req.Params["bottleID"]
	if len(paramBottleID) > 0 {
		rawBottleID := paramBottleID[0]
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("bottleID", rawBottleID, "integer"))
		}
	}
	return &rctx, err
}

// updateBottlePayload is the bottle update action payload.
type updateBottlePayload struct {
	Color     *string `json:"color,omitempty" xml:"color,omitempty" form:"color,omitempty"`
	Country   *string `json:"country,omitempty" xml:"country,omitempty" form:"country,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty" form:"name,omitempty"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty" form:"region,omitempty"`
	Review    *string `json:"review,omitempty" xml:"review,omitempty" form:"review,omitempty"`
	Sweetness *int    `json:"sweetness,omitempty" xml:"sweetness,omitempty" form:"sweetness,omitempty"`
	Varietal  *string `json:"varietal,omitempty" xml:"varietal,omitempty" form:"varietal,omitempty"`
	Vineyard  *string `json:"vineyard,omitempty" xml:"vineyard,omitempty" form:"vineyard,omitempty"`
	Vintage   *int    `json:"vintage,omitempty" xml:"vintage,omitempty" form:"vintage,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateBottlePayload) Validate() (err error) {
	if payload.Color != nil {
		if !(*payload.Color == "red" || *payload.Color == "white" || *payload.Color == "rose" || *payload.Color == "yellow" || *payload.Color == "sparkling") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`raw.color`, *payload.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}))
		}
	}
	if payload.Country != nil {
		if len(*payload.Country) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.country`, *payload.Country, len(*payload.Country), 2, true))
		}
	}
	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true))
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 3, true))
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) > 300 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 300, false))
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 1, true))
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 5, false))
		}
	}
	if payload.Varietal != nil {
		if len(*payload.Varietal) < 4 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.varietal`, *payload.Varietal, len(*payload.Varietal), 4, true))
		}
	}
	if payload.Vineyard != nil {
		if len(*payload.Vineyard) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.vineyard`, *payload.Vineyard, len(*payload.Vineyard), 2, true))
		}
	}
	if payload.Vintage != nil {
		if *payload.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.vintage`, *payload.Vintage, 1900, true))
		}
	}
	if payload.Vintage != nil {
		if *payload.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.vintage`, *payload.Vintage, 2020, false))
		}
	}
	return
}

// Publicize creates UpdateBottlePayload from updateBottlePayload
func (payload *updateBottlePayload) Publicize() *UpdateBottlePayload {
	var pub UpdateBottlePayload
	if payload.Color != nil {
		pub.Color = payload.Color
	}
	if payload.Country != nil {
		pub.Country = payload.Country
	}
	if payload.Name != nil {
		pub.Name = payload.Name
	}
	if payload.Region != nil {
		pub.Region = payload.Region
	}
	if payload.Review != nil {
		pub.Review = payload.Review
	}
	if payload.Sweetness != nil {
		pub.Sweetness = payload.Sweetness
	}
	if payload.Varietal != nil {
		pub.Varietal = payload.Varietal
	}
	if payload.Vineyard != nil {
		pub.Vineyard = payload.Vineyard
	}
	if payload.Vintage != nil {
		pub.Vintage = payload.Vintage
	}
	return &pub
}

// UpdateBottlePayload is the bottle update action payload.
type UpdateBottlePayload struct {
	Color     *string `json:"color,omitempty" xml:"color,omitempty" form:"color,omitempty"`
	Country   *string `json:"country,omitempty" xml:"country,omitempty" form:"country,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty" form:"name,omitempty"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty" form:"region,omitempty"`
	Review    *string `json:"review,omitempty" xml:"review,omitempty" form:"review,omitempty"`
	Sweetness *int    `json:"sweetness,omitempty" xml:"sweetness,omitempty" form:"sweetness,omitempty"`
	Varietal  *string `json:"varietal,omitempty" xml:"varietal,omitempty" form:"varietal,omitempty"`
	Vineyard  *string `json:"vineyard,omitempty" xml:"vineyard,omitempty" form:"vineyard,omitempty"`
	Vintage   *int    `json:"vintage,omitempty" xml:"vintage,omitempty" form:"vintage,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateBottlePayload) Validate() (err error) {
	if payload.Color != nil {
		if !(*payload.Color == "red" || *payload.Color == "white" || *payload.Color == "rose" || *payload.Color == "yellow" || *payload.Color == "sparkling") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`raw.color`, *payload.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}))
		}
	}
	if payload.Country != nil {
		if len(*payload.Country) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.country`, *payload.Country, len(*payload.Country), 2, true))
		}
	}
	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true))
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 3, true))
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) > 300 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 300, false))
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 1, true))
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 5, false))
		}
	}
	if payload.Varietal != nil {
		if len(*payload.Varietal) < 4 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.varietal`, *payload.Varietal, len(*payload.Varietal), 4, true))
		}
	}
	if payload.Vineyard != nil {
		if len(*payload.Vineyard) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.vineyard`, *payload.Vineyard, len(*payload.Vineyard), 2, true))
		}
	}
	if payload.Vintage != nil {
		if *payload.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.vintage`, *payload.Vintage, 1900, true))
		}
	}
	if payload.Vintage != nil {
		if *payload.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.vintage`, *payload.Vintage, 2020, false))
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
func NewWatchBottleContext(ctx context.Context, service *goa.Service) (*WatchBottleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := WatchBottleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountID := req.Params["accountID"]
	if len(paramAccountID) > 0 {
		rawAccountID := paramAccountID[0]
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			rctx.AccountID = accountID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountID", rawAccountID, "integer"))
		}
	}
	paramBottleID := req.Params["bottleID"]
	if len(paramBottleID) > 0 {
		rawBottleID := paramBottleID[0]
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("bottleID", rawBottleID, "integer"))
		}
	}
	return &rctx, err
}
