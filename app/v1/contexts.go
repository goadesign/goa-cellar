//************************************************************************//
// API "cellar" version 1.0: Application Contexts
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

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/raphael/goa"
)

// CreateBottleContext provides the bottle create action context.
type CreateBottleContext struct {
	*goa.Context
	AccountID *int
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
			tmp17 := int(accountID)
			tmp16 := &tmp17
			ctx.AccountID = tmp16
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
			var tmp18 string
			if val, ok := v.(string); ok {
				tmp18 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Color`, v, "string", err)
			}
			if err == nil {
				if !(tmp18 == "red" || tmp18 == "white" || tmp18 == "rose" || tmp18 == "yellow" || tmp18 == "sparkling") {
					err = goa.InvalidEnumValueError(`payload.Color`, tmp18, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
				}
			}
			target.Color = tmp18
		} else {
			err = goa.MissingAttributeError(`payload`, "color", err)
		}
		if v, ok := val["country"]; ok {
			var tmp19 string
			if val, ok := v.(string); ok {
				tmp19 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Country`, v, "string", err)
			}
			if err == nil {
				if len(tmp19) < 2 {
					err = goa.InvalidLengthError(`payload.Country`, tmp19, len(tmp19), 2, true, err)
				}
			}
			target.Country = &tmp19
		}
		if v, ok := val["name"]; ok {
			var tmp20 string
			if val, ok := v.(string); ok {
				tmp20 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp20) < 2 {
					err = goa.InvalidLengthError(`payload.Name`, tmp20, len(tmp20), 2, true, err)
				}
			}
			target.Name = tmp20
		} else {
			err = goa.MissingAttributeError(`payload`, "name", err)
		}
		if v, ok := val["region"]; ok {
			var tmp21 string
			if val, ok := v.(string); ok {
				tmp21 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Region`, v, "string", err)
			}
			target.Region = &tmp21
		}
		if v, ok := val["review"]; ok {
			var tmp22 string
			if val, ok := v.(string); ok {
				tmp22 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Review`, v, "string", err)
			}
			if err == nil {
				if len(tmp22) < 10 {
					err = goa.InvalidLengthError(`payload.Review`, tmp22, len(tmp22), 10, true, err)
				}
				if len(tmp22) > 300 {
					err = goa.InvalidLengthError(`payload.Review`, tmp22, len(tmp22), 300, false, err)
				}
			}
			target.Review = &tmp22
		}
		if v, ok := val["sweetness"]; ok {
			var tmp23 int
			if f, ok := v.(float64); ok {
				tmp23 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Sweetness`, v, "int", err)
			}
			if err == nil {
				if tmp23 < 1 {
					err = goa.InvalidRangeError(`payload.Sweetness`, tmp23, 1, true, err)
				}
				if tmp23 > 5 {
					err = goa.InvalidRangeError(`payload.Sweetness`, tmp23, 5, false, err)
				}
			}
			target.Sweetness = &tmp23
		}
		if v, ok := val["varietal"]; ok {
			var tmp24 string
			if val, ok := v.(string); ok {
				tmp24 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Varietal`, v, "string", err)
			}
			if err == nil {
				if len(tmp24) < 4 {
					err = goa.InvalidLengthError(`payload.Varietal`, tmp24, len(tmp24), 4, true, err)
				}
			}
			target.Varietal = tmp24
		} else {
			err = goa.MissingAttributeError(`payload`, "varietal", err)
		}
		if v, ok := val["vineyard"]; ok {
			var tmp25 string
			if val, ok := v.(string); ok {
				tmp25 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Vineyard`, v, "string", err)
			}
			if err == nil {
				if len(tmp25) < 2 {
					err = goa.InvalidLengthError(`payload.Vineyard`, tmp25, len(tmp25), 2, true, err)
				}
			}
			target.Vineyard = tmp25
		} else {
			err = goa.MissingAttributeError(`payload`, "vineyard", err)
		}
		if v, ok := val["vintage"]; ok {
			var tmp26 int
			if f, ok := v.(float64); ok {
				tmp26 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Vintage`, v, "int", err)
			}
			if err == nil {
				if tmp26 < 1900 {
					err = goa.InvalidRangeError(`payload.Vintage`, tmp26, 1900, true, err)
				}
				if tmp26 > 2020 {
					err = goa.InvalidRangeError(`payload.Vintage`, tmp26, 2020, false, err)
				}
			}
			target.Vintage = tmp26
		} else {
			err = goa.MissingAttributeError(`payload`, "vintage", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
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
	AccountID *int
	BottleID  *int
}

// NewDeleteBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller delete action.
func NewDeleteBottleContext(c *goa.Context) (*DeleteBottleContext, error) {
	var err error
	ctx := DeleteBottleContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			tmp28 := int(accountID)
			tmp27 := &tmp28
			ctx.AccountID = tmp27
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := c.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			tmp30 := int(bottleID)
			tmp29 := &tmp30
			ctx.BottleID = tmp29
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
	AccountID *int
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
			tmp32 := int(accountID)
			tmp31 := &tmp32
			ctx.AccountID = tmp31
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
	AccountID *int
	BottleID  *int
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
			tmp35 := int(accountID)
			tmp34 := &tmp35
			ctx.AccountID = tmp34
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := c.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			tmp37 := int(bottleID)
			tmp36 := &tmp37
			ctx.BottleID = tmp36
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
			var tmp38 int
			if f, ok := v.(float64); ok {
				tmp38 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Rating`, v, "int", err)
			}
			if err == nil {
				if tmp38 < 1 {
					err = goa.InvalidRangeError(`payload.Rating`, tmp38, 1, true, err)
				}
				if tmp38 > 5 {
					err = goa.InvalidRangeError(`payload.Rating`, tmp38, 5, false, err)
				}
			}
			target.Rating = tmp38
		} else {
			err = goa.MissingAttributeError(`payload`, "rating", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
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
	AccountID *int
	BottleID  *int
}

// NewShowBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller show action.
func NewShowBottleContext(c *goa.Context) (*ShowBottleContext, error) {
	var err error
	ctx := ShowBottleContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			tmp40 := int(accountID)
			tmp39 := &tmp40
			ctx.AccountID = tmp39
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := c.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			tmp42 := int(bottleID)
			tmp41 := &tmp42
			ctx.BottleID = tmp41
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
	AccountID *int
	BottleID  *int
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
			tmp44 := int(accountID)
			tmp43 := &tmp44
			ctx.AccountID = tmp43
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := c.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			tmp46 := int(bottleID)
			tmp45 := &tmp46
			ctx.BottleID = tmp45
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
			var tmp47 string
			if val, ok := v.(string); ok {
				tmp47 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Color`, v, "string", err)
			}
			if err == nil {
				if !(tmp47 == "red" || tmp47 == "white" || tmp47 == "rose" || tmp47 == "yellow" || tmp47 == "sparkling") {
					err = goa.InvalidEnumValueError(`payload.Color`, tmp47, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
				}
			}
			target.Color = &tmp47
		}
		if v, ok := val["country"]; ok {
			var tmp48 string
			if val, ok := v.(string); ok {
				tmp48 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Country`, v, "string", err)
			}
			if err == nil {
				if len(tmp48) < 2 {
					err = goa.InvalidLengthError(`payload.Country`, tmp48, len(tmp48), 2, true, err)
				}
			}
			target.Country = &tmp48
		}
		if v, ok := val["name"]; ok {
			var tmp49 string
			if val, ok := v.(string); ok {
				tmp49 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp49) < 2 {
					err = goa.InvalidLengthError(`payload.Name`, tmp49, len(tmp49), 2, true, err)
				}
			}
			target.Name = &tmp49
		}
		if v, ok := val["region"]; ok {
			var tmp50 string
			if val, ok := v.(string); ok {
				tmp50 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Region`, v, "string", err)
			}
			target.Region = &tmp50
		}
		if v, ok := val["review"]; ok {
			var tmp51 string
			if val, ok := v.(string); ok {
				tmp51 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Review`, v, "string", err)
			}
			if err == nil {
				if len(tmp51) < 10 {
					err = goa.InvalidLengthError(`payload.Review`, tmp51, len(tmp51), 10, true, err)
				}
				if len(tmp51) > 300 {
					err = goa.InvalidLengthError(`payload.Review`, tmp51, len(tmp51), 300, false, err)
				}
			}
			target.Review = &tmp51
		}
		if v, ok := val["sweetness"]; ok {
			var tmp52 int
			if f, ok := v.(float64); ok {
				tmp52 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Sweetness`, v, "int", err)
			}
			if err == nil {
				if tmp52 < 1 {
					err = goa.InvalidRangeError(`payload.Sweetness`, tmp52, 1, true, err)
				}
				if tmp52 > 5 {
					err = goa.InvalidRangeError(`payload.Sweetness`, tmp52, 5, false, err)
				}
			}
			target.Sweetness = &tmp52
		}
		if v, ok := val["varietal"]; ok {
			var tmp53 string
			if val, ok := v.(string); ok {
				tmp53 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Varietal`, v, "string", err)
			}
			if err == nil {
				if len(tmp53) < 4 {
					err = goa.InvalidLengthError(`payload.Varietal`, tmp53, len(tmp53), 4, true, err)
				}
			}
			target.Varietal = &tmp53
		}
		if v, ok := val["vineyard"]; ok {
			var tmp54 string
			if val, ok := v.(string); ok {
				tmp54 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Vineyard`, v, "string", err)
			}
			if err == nil {
				if len(tmp54) < 2 {
					err = goa.InvalidLengthError(`payload.Vineyard`, tmp54, len(tmp54), 2, true, err)
				}
			}
			target.Vineyard = &tmp54
		}
		if v, ok := val["vintage"]; ok {
			var tmp55 int
			if f, ok := v.(float64); ok {
				tmp55 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Vintage`, v, "int", err)
			}
			if err == nil {
				if tmp55 < 1900 {
					err = goa.InvalidRangeError(`payload.Vintage`, tmp55, 1900, true, err)
				}
				if tmp55 > 2020 {
					err = goa.InvalidRangeError(`payload.Vintage`, tmp55, 2020, false, err)
				}
			}
			target.Vintage = &tmp55
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
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
