//************************************************************************//
// API "cellar": Application Media Types
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

import "github.com/goadesign/goa"

// Account media type.
//
// Identifier: application/vnd.account+json
type Account struct {
	// Date of creation
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Email of account owner
	CreatedBy *string `json:"created_by,omitempty" xml:"created_by,omitempty"`
	// API href of account
	Href string `json:"href" xml:"href"`
	// ID of account
	ID int `json:"id" xml:"id"`
	// Name of account
	Name string `json:"name" xml:"name"`
}

// Validate validates the Account media type instance.
func (mt *Account) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}

	if mt.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.created_at`, *mt.CreatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if mt.CreatedBy != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.CreatedBy); err2 != nil {
			err = goa.InvalidFormatError(`response.created_by`, *mt.CreatedBy, goa.FormatEmail, err2, err)
		}
	}
	return
}

// AccountLink media type.
//
// Identifier: application/vnd.account+json
type AccountLink struct {
	// API href of account
	Href string `json:"href" xml:"href"`
	// ID of account
	ID int `json:"id" xml:"id"`
}

// Validate validates the AccountLink media type instance.
func (mt *AccountLink) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}

	return
}

// AccountTiny media type.
//
// Identifier: application/vnd.account+json
type AccountTiny struct {
	// API href of account
	Href string `json:"href" xml:"href"`
	// ID of account
	ID int `json:"id" xml:"id"`
	// Name of account
	Name string `json:"name" xml:"name"`
}

// Validate validates the AccountTiny media type instance.
func (mt *AccountTiny) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}

	return
}

// Bottle media type.
//
// Identifier: application/vnd.bottle+json
type Bottle struct {
	// Account that owns bottle
	Account *Account `json:"account,omitempty" xml:"account,omitempty"`
	// API href of bottle
	Href string `json:"href" xml:"href"`
	// ID of bottle
	ID int `json:"id" xml:"id"`
	// Links to related resources
	Links *BottleLinks `json:"links,omitempty" xml:"links,omitempty"`
	Name  string       `json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating   *int   `json:"rating,omitempty" xml:"rating,omitempty"`
	Varietal string `json:"varietal" xml:"varietal"`
	Vineyard string `json:"vineyard" xml:"vineyard"`
	Vintage  int    `json:"vintage" xml:"vintage"`
}

// Validate validates the Bottle media type instance.
func (mt *Bottle) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
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

// BottleFull media type.
//
// Identifier: application/vnd.bottle+json
type BottleFull struct {
	// Account that owns bottle
	Account *Account `json:"account,omitempty" xml:"account,omitempty"`
	Color   string   `json:"color" xml:"color"`
	Country *string  `json:"country,omitempty" xml:"country,omitempty"`
	// Date of creation
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// API href of bottle
	Href string `json:"href" xml:"href"`
	// ID of bottle
	ID int `json:"id" xml:"id"`
	// Links to related resources
	Links *BottleLinks `json:"links,omitempty" xml:"links,omitempty"`
	Name  string       `json:"name" xml:"name"`
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

// Validate validates the BottleFull media type instance.
func (mt *BottleFull) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
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

// BottleTiny media type.
//
// Identifier: application/vnd.bottle+json
type BottleTiny struct {
	// API href of bottle
	Href string `json:"href" xml:"href"`
	// ID of bottle
	ID int `json:"id" xml:"id"`
	// Links to related resources
	Links *BottleLinks `json:"links,omitempty" xml:"links,omitempty"`
	Name  string       `json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating *int `json:"rating,omitempty" xml:"rating,omitempty"`
}

// Validate validates the BottleTiny media type instance.
func (mt *BottleTiny) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
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

// BottleLinks contains links to related resources of Bottle.
type BottleLinks struct {
	Account *AccountLink `json:"account,omitempty" xml:"account,omitempty"`
}

// BottleCollection media type is a collection of Bottle.
//
// Identifier: application/vnd.bottle+json; type=collection
type BottleCollection []*Bottle

// Validate validates the BottleCollection media type instance.
func (mt BottleCollection) Validate() (err error) {
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

// BottleTinyCollection media type is a collection of BottleTiny.
//
// Identifier: application/vnd.bottle+json; type=collection
type BottleTinyCollection []*BottleTiny

// Validate validates the BottleTinyCollection media type instance.
func (mt BottleTinyCollection) Validate() (err error) {
	for _, e := range mt {
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

// BottleLinksArray contains links to related resources of BottleCollection.
type BottleLinksArray []*BottleLinks
