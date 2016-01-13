//************************************************************************//
// API "cellar": Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/raphael/goa-cellar
// --design=github.com/raphael/goa-cellar/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/raphael/goa"
)

// CreateAccountContext provides the account create action context.
type CreateAccountContext struct {
	*goa.Context
	Payload *CreateAccountPayload
}

// NewCreateAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller create action.
func NewCreateAccountContext(c *goa.Context) (*CreateAccountContext, error) {
	var err error
	ctx := CreateAccountContext{Context: c}
	p, err := NewCreateAccountPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateAccountPayload is the account create action payload.
type CreateAccountPayload struct {
	// Name of account
	Name string
}

// NewCreateAccountPayload instantiates a CreateAccountPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateAccountPayload(raw interface{}) (p *CreateAccountPayload, err error) {
	p, err = UnmarshalCreateAccountPayload(raw, err)
	return
}

// UnmarshalCreateAccountPayload unmarshals and validates a raw interface{} into an instance of CreateAccountPayload
func UnmarshalCreateAccountPayload(source interface{}, inErr error) (target *CreateAccountPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateAccountPayload)
		if v, ok := val["name"]; ok {
			var tmp1 string
			if val, ok := v.(string); ok {
				tmp1 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			target.Name = tmp1
		} else {
			err = goa.MissingAttributeError(`payload`, "name", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	if err2 := target.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
	}
	return
}

// Validate validates the type instance.
func (payload *CreateAccountPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MissingAttributeError(`raw`, "name", err)
	}

	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateAccountContext) Created() error {
	return ctx.Respond(201, nil)
}

// DeleteAccountContext provides the account delete action context.
type DeleteAccountContext struct {
	*goa.Context
	AccountID int
}

// NewDeleteAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller delete action.
func NewDeleteAccountContext(c *goa.Context) (*DeleteAccountContext, error) {
	var err error
	ctx := DeleteAccountContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteAccountContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteAccountContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ShowAccountContext provides the account show action context.
type ShowAccountContext struct {
	*goa.Context
	AccountID int
}

// NewShowAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller show action.
func NewShowAccountContext(c *goa.Context) (*ShowAccountContext, error) {
	var err error
	ctx := ShowAccountContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowAccountContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowAccountContext) OK(resp *Account, view AccountViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.account+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateAccountContext provides the account update action context.
type UpdateAccountContext struct {
	*goa.Context
	AccountID int
	Payload   *UpdateAccountPayload
}

// NewUpdateAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller update action.
func NewUpdateAccountContext(c *goa.Context) (*UpdateAccountContext, error) {
	var err error
	ctx := UpdateAccountContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	p, err := NewUpdateAccountPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateAccountPayload is the account update action payload.
type UpdateAccountPayload struct {
	// Name of account
	Name string
}

// NewUpdateAccountPayload instantiates a UpdateAccountPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateAccountPayload(raw interface{}) (p *UpdateAccountPayload, err error) {
	p, err = UnmarshalUpdateAccountPayload(raw, err)
	return
}

// UnmarshalUpdateAccountPayload unmarshals and validates a raw interface{} into an instance of UpdateAccountPayload
func UnmarshalUpdateAccountPayload(source interface{}, inErr error) (target *UpdateAccountPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateAccountPayload)
		if v, ok := val["name"]; ok {
			var tmp5 string
			if val, ok := v.(string); ok {
				tmp5 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			target.Name = tmp5
		} else {
			err = goa.MissingAttributeError(`payload`, "name", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	if err2 := target.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
	}
	return
}

// Validate validates the type instance.
func (payload *UpdateAccountPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MissingAttributeError(`raw`, "name", err)
	}

	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateAccountContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateAccountContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// CreateBottleContext provides the bottle create action context.
type CreateBottleContext struct {
	*goa.Context
	AccountID int
	Payload   *CreateBottlePayload
}

// NewCreateBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller create action.
func NewCreateBottleContext(c *goa.Context) (*CreateBottleContext, error) {
	var err error
	ctx := CreateBottleContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	p, err := NewCreateBottlePayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateBottlePayload is the bottle create action payload.
type CreateBottlePayload struct {
	Color     string
	Country   *string
	Name      string
	Region    *string
	Review    *string
	Sweetness *int
	Varietal  string
	Vineyard  string
	Vintage   int
}

// NewCreateBottlePayload instantiates a CreateBottlePayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateBottlePayload(raw interface{}) (p *CreateBottlePayload, err error) {
	p, err = UnmarshalCreateBottlePayload(raw, err)
	return
}

// UnmarshalCreateBottlePayload unmarshals and validates a raw interface{} into an instance of CreateBottlePayload
func UnmarshalCreateBottlePayload(source interface{}, inErr error) (target *CreateBottlePayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateBottlePayload)
		if v, ok := val["color"]; ok {
			var tmp7 string
			if val, ok := v.(string); ok {
				tmp7 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Color`, v, "string", err)
			}
			target.Color = tmp7
		} else {
			err = goa.MissingAttributeError(`payload`, "color", err)
		}
		if v, ok := val["country"]; ok {
			var tmp8 string
			if val, ok := v.(string); ok {
				tmp8 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Country`, v, "string", err)
			}
			target.Country = &tmp8
		}
		if v, ok := val["name"]; ok {
			var tmp9 string
			if val, ok := v.(string); ok {
				tmp9 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			target.Name = tmp9
		} else {
			err = goa.MissingAttributeError(`payload`, "name", err)
		}
		if v, ok := val["region"]; ok {
			var tmp10 string
			if val, ok := v.(string); ok {
				tmp10 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Region`, v, "string", err)
			}
			target.Region = &tmp10
		}
		if v, ok := val["review"]; ok {
			var tmp11 string
			if val, ok := v.(string); ok {
				tmp11 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Review`, v, "string", err)
			}
			target.Review = &tmp11
		}
		if v, ok := val["sweetness"]; ok {
			var tmp12 int
			if f, ok := v.(float64); ok {
				tmp12 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Sweetness`, v, "int", err)
			}
			target.Sweetness = &tmp12
		}
		if v, ok := val["varietal"]; ok {
			var tmp13 string
			if val, ok := v.(string); ok {
				tmp13 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Varietal`, v, "string", err)
			}
			target.Varietal = tmp13
		} else {
			err = goa.MissingAttributeError(`payload`, "varietal", err)
		}
		if v, ok := val["vineyard"]; ok {
			var tmp14 string
			if val, ok := v.(string); ok {
				tmp14 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Vineyard`, v, "string", err)
			}
			target.Vineyard = tmp14
		} else {
			err = goa.MissingAttributeError(`payload`, "vineyard", err)
		}
		if v, ok := val["vintage"]; ok {
			var tmp15 int
			if f, ok := v.(float64); ok {
				tmp15 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Vintage`, v, "int", err)
			}
			target.Vintage = tmp15
		} else {
			err = goa.MissingAttributeError(`payload`, "vintage", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	if err2 := target.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
	}
	return
}

// Validate validates the type instance.
func (payload *CreateBottlePayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MissingAttributeError(`raw`, "name", err)
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
func (ctx *CreateBottleContext) Created() error {
	return ctx.Respond(201, nil)
}

// DeleteBottleContext provides the bottle delete action context.
type DeleteBottleContext struct {
	*goa.Context
	AccountID int
	BottleID  int
}

// NewDeleteBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller delete action.
func NewDeleteBottleContext(c *goa.Context) (*DeleteBottleContext, error) {
	var err error
	ctx := DeleteBottleContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := c.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			ctx.BottleID = int(bottleID)
		} else {
			err = goa.InvalidParamTypeError("bottleID", rawBottleID, "integer", err)
		}
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteBottleContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteBottleContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ListBottleContext provides the bottle list action context.
type ListBottleContext struct {
	*goa.Context
	AccountID int
	Years     []int
}

// NewListBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller list action.
func NewListBottleContext(c *goa.Context) (*ListBottleContext, error) {
	var err error
	ctx := ListBottleContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawYears := c.Get("years")
	if rawYears != "" {
		elemsYears := strings.Split(rawYears, ",")
		elemsYears2 := make([]int, len(elemsYears))
		for i, rawElem := range elemsYears {
			if elem, err2 := strconv.Atoi(rawElem); err2 == nil {
				elemsYears2[i] = int(elem)
			} else {
				err = goa.InvalidParamTypeError("elem", rawElem, "integer", err)
			}
		}
		ctx.Years = elemsYears2
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListBottleContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ListBottleContext) OK(resp BottleCollection, view BottleCollectionViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.bottle+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// RateBottleContext provides the bottle rate action context.
type RateBottleContext struct {
	*goa.Context
	AccountID int
	BottleID  int
	Payload   *RateBottlePayload
}

// NewRateBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller rate action.
func NewRateBottleContext(c *goa.Context) (*RateBottleContext, error) {
	var err error
	ctx := RateBottleContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := c.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			ctx.BottleID = int(bottleID)
		} else {
			err = goa.InvalidParamTypeError("bottleID", rawBottleID, "integer", err)
		}
	}
	p, err := NewRateBottlePayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// RateBottlePayload is the bottle rate action payload.
type RateBottlePayload struct {
	// Rating of bottle between 1 and 5
	Rating int
}

// NewRateBottlePayload instantiates a RateBottlePayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewRateBottlePayload(raw interface{}) (p *RateBottlePayload, err error) {
	p, err = UnmarshalRateBottlePayload(raw, err)
	return
}

// UnmarshalRateBottlePayload unmarshals and validates a raw interface{} into an instance of RateBottlePayload
func UnmarshalRateBottlePayload(source interface{}, inErr error) (target *RateBottlePayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(RateBottlePayload)
		if v, ok := val["rating"]; ok {
			var tmp22 int
			if f, ok := v.(float64); ok {
				tmp22 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Rating`, v, "int", err)
			}
			target.Rating = tmp22
		} else {
			err = goa.MissingAttributeError(`payload`, "rating", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	if err2 := target.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
	}
	return
}

// Validate validates the type instance.
func (payload *RateBottlePayload) Validate() (err error) {

	if payload.Rating < 1 {
		err = goa.InvalidRangeError(`raw.rating`, payload.Rating, 1, true, err)
	}
	if payload.Rating > 5 {
		err = goa.InvalidRangeError(`raw.rating`, payload.Rating, 5, false, err)
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *RateBottleContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *RateBottleContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ShowBottleContext provides the bottle show action context.
type ShowBottleContext struct {
	*goa.Context
	AccountID int
	BottleID  int
}

// NewShowBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller show action.
func NewShowBottleContext(c *goa.Context) (*ShowBottleContext, error) {
	var err error
	ctx := ShowBottleContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := c.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			ctx.BottleID = int(bottleID)
		} else {
			err = goa.InvalidParamTypeError("bottleID", rawBottleID, "integer", err)
		}
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowBottleContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowBottleContext) OK(resp *Bottle, view BottleViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.bottle+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateBottleContext provides the bottle update action context.
type UpdateBottleContext struct {
	*goa.Context
	AccountID int
	BottleID  int
	Payload   *UpdateBottlePayload
}

// NewUpdateBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller update action.
func NewUpdateBottleContext(c *goa.Context) (*UpdateBottleContext, error) {
	var err error
	ctx := UpdateBottleContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := c.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			ctx.BottleID = int(bottleID)
		} else {
			err = goa.InvalidParamTypeError("bottleID", rawBottleID, "integer", err)
		}
	}
	p, err := NewUpdateBottlePayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateBottlePayload is the bottle update action payload.
type UpdateBottlePayload struct {
	Color     *string
	Country   *string
	Name      *string
	Region    *string
	Review    *string
	Sweetness *int
	Varietal  *string
	Vineyard  *string
	Vintage   *int
}

// NewUpdateBottlePayload instantiates a UpdateBottlePayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateBottlePayload(raw interface{}) (p *UpdateBottlePayload, err error) {
	p, err = UnmarshalUpdateBottlePayload(raw, err)
	return
}

// UnmarshalUpdateBottlePayload unmarshals and validates a raw interface{} into an instance of UpdateBottlePayload
func UnmarshalUpdateBottlePayload(source interface{}, inErr error) (target *UpdateBottlePayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateBottlePayload)
		if v, ok := val["color"]; ok {
			var tmp27 string
			if val, ok := v.(string); ok {
				tmp27 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Color`, v, "string", err)
			}
			target.Color = &tmp27
		}
		if v, ok := val["country"]; ok {
			var tmp28 string
			if val, ok := v.(string); ok {
				tmp28 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Country`, v, "string", err)
			}
			target.Country = &tmp28
		}
		if v, ok := val["name"]; ok {
			var tmp29 string
			if val, ok := v.(string); ok {
				tmp29 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			target.Name = &tmp29
		}
		if v, ok := val["region"]; ok {
			var tmp30 string
			if val, ok := v.(string); ok {
				tmp30 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Region`, v, "string", err)
			}
			target.Region = &tmp30
		}
		if v, ok := val["review"]; ok {
			var tmp31 string
			if val, ok := v.(string); ok {
				tmp31 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Review`, v, "string", err)
			}
			target.Review = &tmp31
		}
		if v, ok := val["sweetness"]; ok {
			var tmp32 int
			if f, ok := v.(float64); ok {
				tmp32 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Sweetness`, v, "int", err)
			}
			target.Sweetness = &tmp32
		}
		if v, ok := val["varietal"]; ok {
			var tmp33 string
			if val, ok := v.(string); ok {
				tmp33 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Varietal`, v, "string", err)
			}
			target.Varietal = &tmp33
		}
		if v, ok := val["vineyard"]; ok {
			var tmp34 string
			if val, ok := v.(string); ok {
				tmp34 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Vineyard`, v, "string", err)
			}
			target.Vineyard = &tmp34
		}
		if v, ok := val["vintage"]; ok {
			var tmp35 int
			if f, ok := v.(float64); ok {
				tmp35 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Vintage`, v, "int", err)
			}
			target.Vintage = &tmp35
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	if err2 := target.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
	}
	return
}

// Validate validates the type instance.
func (payload *UpdateBottlePayload) Validate() (err error) {
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
func (ctx *UpdateBottleContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateBottleContext) NotFound() error {
	return ctx.Respond(404, nil)
}
