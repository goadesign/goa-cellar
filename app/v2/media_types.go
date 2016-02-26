//************************************************************************//
// API "cellar" version 2.0: Application Media Types
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
	"github.com/goadesign/goa-cellar/app"
)

// A bottle of rum or wine
// Identifier: application/vnd.generic.bottle+json
type GenericBottle struct {
	// Account that owns bottle
	Account *app.Account `json:"account,omitempty" xml:"account,omitempty"`
	// API href of bottle
	Href string `json:"href" xml:"href"`
	// ID of bottle
	ID int `json:"id" xml:"id"`
	// The kind of bottle
	Kind string `json:"kind" xml:"kind"`
	// Links to related resources
	Links *GenericBottleLinks `json:"links,omitempty" xml:"links,omitempty"`
	Name  string              `json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating   *int   `json:"rating,omitempty" xml:"rating,omitempty"`
	Varietal string `json:"varietal" xml:"varietal"`
	Vineyard string `json:"vineyard" xml:"vineyard"`
	Vintage  int    `json:"vintage" xml:"vintage"`
}

// Validate validates the media type instance.
func (mt *GenericBottle) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Kind == "" {
		err = goa.MissingAttributeError(`response`, "kind", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}
	if mt.Vineyard == "" {
		err = goa.MissingAttributeError(`response`, "vineyard", err)
	}
	if mt.Varietal == "" {
		err = goa.MissingAttributeError(`response`, "varietal", err)
	}

	if mt.Account != nil {
		if mt.Account.CreatedAt != nil {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.Account.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response.account.created_at`, *mt.Account.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if mt.Account.CreatedBy != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Account.CreatedBy); err2 != nil {
				err = goa.InvalidFormatError(`response.account.created_by`, *mt.Account.CreatedBy, goa.FormatEmail, err2, err)
			}
		}
	}
	if !(mt.Kind == "wine" || mt.Kind == "rum") {
		err = goa.InvalidEnumValueError(`response.kind`, mt.Kind, []interface{}{"wine", "rum"}, err)
	}
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 2, true, err)
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false, err)
		}
	}
	if len(mt.Varietal) < 4 {
		err = goa.InvalidLengthError(`response.varietal`, mt.Varietal, len(mt.Varietal), 4, true, err)
	}
	if len(mt.Vineyard) < 2 {
		err = goa.InvalidLengthError(`response.vineyard`, mt.Vineyard, len(mt.Vineyard), 2, true, err)
	}
	if mt.Vintage < 1900 {
		err = goa.InvalidRangeError(`response.vintage`, mt.Vintage, 1900, true, err)
	}
	if mt.Vintage > 2020 {
		err = goa.InvalidRangeError(`response.vintage`, mt.Vintage, 2020, false, err)
	}
	return
}

// A bottle of rum or wine, full view
// Identifier: application/vnd.generic.bottle+json
type GenericBottleFull struct {
	// Account that owns bottle
	Account *app.Account `json:"account,omitempty" xml:"account,omitempty"`
	Color   string       `json:"color" xml:"color"`
	Country *string      `json:"country,omitempty" xml:"country,omitempty"`
	// Date of creation
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// API href of bottle
	Href string `json:"href" xml:"href"`
	// ID of bottle
	ID int `json:"id" xml:"id"`
	// The kind of bottle
	Kind string `json:"kind" xml:"kind"`
	// Links to related resources
	Links *GenericBottleLinks `json:"links,omitempty" xml:"links,omitempty"`
	Name  string              `json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating    *int    `json:"rating,omitempty" xml:"rating,omitempty"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	// Date of last update
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	Varietal  string  `json:"varietal" xml:"varietal"`
	Vineyard  string  `json:"vineyard" xml:"vineyard"`
	Vintage   int     `json:"vintage" xml:"vintage"`
}

// Validate validates the media type instance.
func (mt *GenericBottleFull) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Kind == "" {
		err = goa.MissingAttributeError(`response`, "kind", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}
	if mt.Vineyard == "" {
		err = goa.MissingAttributeError(`response`, "vineyard", err)
	}
	if mt.Varietal == "" {
		err = goa.MissingAttributeError(`response`, "varietal", err)
	}

	if mt.Color == "" {
		err = goa.MissingAttributeError(`response`, "color", err)
	}

	if mt.Account != nil {
		if mt.Account.CreatedAt != nil {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.Account.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response.account.created_at`, *mt.Account.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if mt.Account.CreatedBy != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Account.CreatedBy); err2 != nil {
				err = goa.InvalidFormatError(`response.account.created_by`, *mt.Account.CreatedBy, goa.FormatEmail, err2, err)
			}
		}
	}
	if !(mt.Color == "red" || mt.Color == "white" || mt.Color == "rose" || mt.Color == "yellow" || mt.Color == "sparkling") {
		err = goa.InvalidEnumValueError(`response.color`, mt.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
	}
	if mt.Country != nil {
		if len(*mt.Country) < 2 {
			err = goa.InvalidLengthError(`response.country`, *mt.Country, len(*mt.Country), 2, true, err)
		}
	}
	if mt.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.created_at`, *mt.CreatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if !(mt.Kind == "wine" || mt.Kind == "rum") {
		err = goa.InvalidEnumValueError(`response.kind`, mt.Kind, []interface{}{"wine", "rum"}, err)
	}
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 2, true, err)
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false, err)
		}
	}
	if mt.Review != nil {
		if len(*mt.Review) < 3 {
			err = goa.InvalidLengthError(`response.review`, *mt.Review, len(*mt.Review), 3, true, err)
		}
	}
	if mt.Review != nil {
		if len(*mt.Review) > 300 {
			err = goa.InvalidLengthError(`response.review`, *mt.Review, len(*mt.Review), 300, false, err)
		}
	}
	if mt.Sweetness != nil {
		if *mt.Sweetness < 1 {
			err = goa.InvalidRangeError(`response.sweetness`, *mt.Sweetness, 1, true, err)
		}
	}
	if mt.Sweetness != nil {
		if *mt.Sweetness > 5 {
			err = goa.InvalidRangeError(`response.sweetness`, *mt.Sweetness, 5, false, err)
		}
	}
	if mt.UpdatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.UpdatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.updated_at`, *mt.UpdatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if len(mt.Varietal) < 4 {
		err = goa.InvalidLengthError(`response.varietal`, mt.Varietal, len(mt.Varietal), 4, true, err)
	}
	if len(mt.Vineyard) < 2 {
		err = goa.InvalidLengthError(`response.vineyard`, mt.Vineyard, len(mt.Vineyard), 2, true, err)
	}
	if mt.Vintage < 1900 {
		err = goa.InvalidRangeError(`response.vintage`, mt.Vintage, 1900, true, err)
	}
	if mt.Vintage > 2020 {
		err = goa.InvalidRangeError(`response.vintage`, mt.Vintage, 2020, false, err)
	}
	return
}

// A bottle of rum or wine, tiny view
// Identifier: application/vnd.generic.bottle+json
type GenericBottleTiny struct {
	// API href of bottle
	Href string `json:"href" xml:"href"`
	// ID of bottle
	ID int `json:"id" xml:"id"`
	// The kind of bottle
	Kind string `json:"kind" xml:"kind"`
	// Links to related resources
	Links *GenericBottleLinks `json:"links,omitempty" xml:"links,omitempty"`
	Name  string              `json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating *int `json:"rating,omitempty" xml:"rating,omitempty"`
}

// Validate validates the media type instance.
func (mt *GenericBottleTiny) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Kind == "" {
		err = goa.MissingAttributeError(`response`, "kind", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}

	if !(mt.Kind == "wine" || mt.Kind == "rum") {
		err = goa.InvalidEnumValueError(`response.kind`, mt.Kind, []interface{}{"wine", "rum"}, err)
	}
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 2, true, err)
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false, err)
		}
	}
	return
}

// GenericBottleLinks contains links to related resources of GenericBottle.
type GenericBottleLinks struct {
	Account *app.AccountLink `json:"account,omitempty" xml:"account,omitempty"`
}

// , default view
// Identifier: application/vnd.generic.bottle+json; type=collection
type GenericBottleCollection []*GenericBottle

// Validate validates the media type instance.
func (mt GenericBottleCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Account != nil {
			if e.Account.CreatedAt != nil {
				if err2 := goa.ValidateFormat(goa.FormatDateTime, *e.Account.CreatedAt); err2 != nil {
					err = goa.InvalidFormatError(`response[*].account.created_at`, *e.Account.CreatedAt, goa.FormatDateTime, err2, err)
				}
			}
			if e.Account.CreatedBy != nil {
				if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Account.CreatedBy); err2 != nil {
					err = goa.InvalidFormatError(`response[*].account.created_by`, *e.Account.CreatedBy, goa.FormatEmail, err2, err)
				}
			}
		}
		if !(e.Kind == "wine" || e.Kind == "rum") {
			err = goa.InvalidEnumValueError(`response[*].kind`, e.Kind, []interface{}{"wine", "rum"}, err)
		}
		if len(e.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].name`, e.Name, len(e.Name), 2, true, err)
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 5, false, err)
			}
		}
		if len(e.Varietal) < 4 {
			err = goa.InvalidLengthError(`response[*].varietal`, e.Varietal, len(e.Varietal), 4, true, err)
		}
		if len(e.Vineyard) < 2 {
			err = goa.InvalidLengthError(`response[*].vineyard`, e.Vineyard, len(e.Vineyard), 2, true, err)
		}
		if e.Vintage < 1900 {
			err = goa.InvalidRangeError(`response[*].vintage`, e.Vintage, 1900, true, err)
		}
		if e.Vintage > 2020 {
			err = goa.InvalidRangeError(`response[*].vintage`, e.Vintage, 2020, false, err)
		}
	}
	return
}

// , tiny view
// Identifier: application/vnd.generic.bottle+json; type=collection
type GenericBottleTinyCollection []*GenericBottleTiny

// Validate validates the media type instance.
func (mt GenericBottleTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if !(e.Kind == "wine" || e.Kind == "rum") {
			err = goa.InvalidEnumValueError(`response[*].kind`, e.Kind, []interface{}{"wine", "rum"}, err)
		}
		if len(e.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].name`, e.Name, len(e.Name), 2, true, err)
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 5, false, err)
			}
		}
	}
	return
}

// GenericBottleLinksArray contains links to related resources of GenericBottleCollection.
type GenericBottleLinksArray []*GenericBottleLinks
