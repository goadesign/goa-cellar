//************************************************************************//
// API "cellar": Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen.exe
// --out=$(GOPATH)\src\github.com\goadesign\goa-cellar
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
		err = goa.MissingAttributeError(`response`, "href")
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name")
	}

	if mt.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.created_at`, *mt.CreatedAt, goa.FormatDateTime, err2)
		}
	}
	if mt.CreatedBy != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.CreatedBy); err2 != nil {
			err = goa.InvalidFormatError(`response.created_by`, *mt.CreatedBy, goa.FormatEmail, err2)
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
		err = goa.MissingAttributeError(`response`, "href")
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
		err = goa.MissingAttributeError(`response`, "href")
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name")
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
		err = goa.MissingAttributeError(`response`, "href")
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name")
	}
	if mt.Vineyard == "" {
		err = goa.MissingAttributeError(`response`, "vineyard")
	}
	if mt.Varietal == "" {
		err = goa.MissingAttributeError(`response`, "varietal")
	}

	if mt.Account != nil {

		if mt.Account.Href == "" {
			err = goa.MissingAttributeError(`response.account`, "href")
		}
		if mt.Account.Name == "" {
			err = goa.MissingAttributeError(`response.account`, "name")
		}

		if mt.Account.CreatedAt != nil {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.Account.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response.account.created_at`, *mt.Account.CreatedAt, goa.FormatDateTime, err2)
			}
		}
		if mt.Account.CreatedBy != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Account.CreatedBy); err2 != nil {
				err = goa.InvalidFormatError(`response.account.created_by`, *mt.Account.CreatedBy, goa.FormatEmail, err2)
			}
		}
	}
	if mt.Links != nil {
		if mt.Links.Account != nil {

			if mt.Links.Account.Href == "" {
				err = goa.MissingAttributeError(`response.links.account`, "href")
			}

		}
	}
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 2, true)
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false)
		}
	}
	if len(mt.Varietal) < 4 {
		err = goa.InvalidLengthError(`response.varietal`, mt.Varietal, len(mt.Varietal), 4, true)
	}
	if len(mt.Vineyard) < 2 {
		err = goa.InvalidLengthError(`response.vineyard`, mt.Vineyard, len(mt.Vineyard), 2, true)
	}
	if mt.Vintage < 1900 {
		err = goa.InvalidRangeError(`response.vintage`, mt.Vintage, 1900, true)
	}
	if mt.Vintage > 2020 {
		err = goa.InvalidRangeError(`response.vintage`, mt.Vintage, 2020, false)
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
		err = goa.MissingAttributeError(`response`, "href")
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name")
	}
	if mt.Vineyard == "" {
		err = goa.MissingAttributeError(`response`, "vineyard")
	}
	if mt.Varietal == "" {
		err = goa.MissingAttributeError(`response`, "varietal")
	}

	if mt.Color == "" {
		err = goa.MissingAttributeError(`response`, "color")
	}

	if mt.Account != nil {

		if mt.Account.Href == "" {
			err = goa.MissingAttributeError(`response.account`, "href")
		}
		if mt.Account.Name == "" {
			err = goa.MissingAttributeError(`response.account`, "name")
		}

		if mt.Account.CreatedAt != nil {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.Account.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response.account.created_at`, *mt.Account.CreatedAt, goa.FormatDateTime, err2)
			}
		}
		if mt.Account.CreatedBy != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Account.CreatedBy); err2 != nil {
				err = goa.InvalidFormatError(`response.account.created_by`, *mt.Account.CreatedBy, goa.FormatEmail, err2)
			}
		}
	}
	if !(mt.Color == "red" || mt.Color == "white" || mt.Color == "rose" || mt.Color == "yellow" || mt.Color == "sparkling") {
		err = goa.InvalidEnumValueError(`response.color`, mt.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"})
	}
	if mt.Country != nil {
		if len(*mt.Country) < 2 {
			err = goa.InvalidLengthError(`response.country`, *mt.Country, len(*mt.Country), 2, true)
		}
	}
	if mt.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.created_at`, *mt.CreatedAt, goa.FormatDateTime, err2)
		}
	}
	if mt.Links != nil {
		if mt.Links.Account != nil {

			if mt.Links.Account.Href == "" {
				err = goa.MissingAttributeError(`response.links.account`, "href")
			}

		}
	}
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 2, true)
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false)
		}
	}
	if mt.Review != nil {
		if len(*mt.Review) < 3 {
			err = goa.InvalidLengthError(`response.review`, *mt.Review, len(*mt.Review), 3, true)
		}
	}
	if mt.Review != nil {
		if len(*mt.Review) > 300 {
			err = goa.InvalidLengthError(`response.review`, *mt.Review, len(*mt.Review), 300, false)
		}
	}
	if mt.Sweetness != nil {
		if *mt.Sweetness < 1 {
			err = goa.InvalidRangeError(`response.sweetness`, *mt.Sweetness, 1, true)
		}
	}
	if mt.Sweetness != nil {
		if *mt.Sweetness > 5 {
			err = goa.InvalidRangeError(`response.sweetness`, *mt.Sweetness, 5, false)
		}
	}
	if mt.UpdatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.UpdatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.updated_at`, *mt.UpdatedAt, goa.FormatDateTime, err2)
		}
	}
	if len(mt.Varietal) < 4 {
		err = goa.InvalidLengthError(`response.varietal`, mt.Varietal, len(mt.Varietal), 4, true)
	}
	if len(mt.Vineyard) < 2 {
		err = goa.InvalidLengthError(`response.vineyard`, mt.Vineyard, len(mt.Vineyard), 2, true)
	}
	if mt.Vintage < 1900 {
		err = goa.InvalidRangeError(`response.vintage`, mt.Vintage, 1900, true)
	}
	if mt.Vintage > 2020 {
		err = goa.InvalidRangeError(`response.vintage`, mt.Vintage, 2020, false)
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
		err = goa.MissingAttributeError(`response`, "href")
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name")
	}

	if mt.Links != nil {
		if mt.Links.Account != nil {

			if mt.Links.Account.Href == "" {
				err = goa.MissingAttributeError(`response.links.account`, "href")
			}

		}
	}
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 2, true)
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false)
		}
	}
	return
}

// BottleLinks contains links to related resources of Bottle.
type BottleLinks struct {
	Account *AccountLink `json:"account,omitempty" xml:"account,omitempty"`
}

// Validate validates the BottleLinks type instance.
func (ut *BottleLinks) Validate() (err error) {
	if ut.Account != nil {

		if ut.Account.Href == "" {
			err = goa.MissingAttributeError(`response.account`, "href")
		}

	}
	return
}

// BottleCollection media type is a collection of Bottle.
//
// Identifier: application/vnd.bottle+json; type=collection
type BottleCollection []*Bottle

// Validate validates the BottleCollection media type instance.
func (mt BottleCollection) Validate() (err error) {
	for _, e := range mt {

		if e.Href == "" {
			err = goa.MissingAttributeError(`response[*]`, "href")
		}
		if e.Name == "" {
			err = goa.MissingAttributeError(`response[*]`, "name")
		}
		if e.Vineyard == "" {
			err = goa.MissingAttributeError(`response[*]`, "vineyard")
		}
		if e.Varietal == "" {
			err = goa.MissingAttributeError(`response[*]`, "varietal")
		}

		if e.Account != nil {

			if e.Account.Href == "" {
				err = goa.MissingAttributeError(`response[*].account`, "href")
			}
			if e.Account.Name == "" {
				err = goa.MissingAttributeError(`response[*].account`, "name")
			}

			if e.Account.CreatedAt != nil {
				if err2 := goa.ValidateFormat(goa.FormatDateTime, *e.Account.CreatedAt); err2 != nil {
					err = goa.InvalidFormatError(`response[*].account.created_at`, *e.Account.CreatedAt, goa.FormatDateTime, err2)
				}
			}
			if e.Account.CreatedBy != nil {
				if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Account.CreatedBy); err2 != nil {
					err = goa.InvalidFormatError(`response[*].account.created_by`, *e.Account.CreatedBy, goa.FormatEmail, err2)
				}
			}
		}
		if e.Links != nil {
			if e.Links.Account != nil {

				if e.Links.Account.Href == "" {
					err = goa.MissingAttributeError(`response[*].links.account`, "href")
				}

			}
		}
		if len(e.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].name`, e.Name, len(e.Name), 2, true)
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 1, true)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 5, false)
			}
		}
		if len(e.Varietal) < 4 {
			err = goa.InvalidLengthError(`response[*].varietal`, e.Varietal, len(e.Varietal), 4, true)
		}
		if len(e.Vineyard) < 2 {
			err = goa.InvalidLengthError(`response[*].vineyard`, e.Vineyard, len(e.Vineyard), 2, true)
		}
		if e.Vintage < 1900 {
			err = goa.InvalidRangeError(`response[*].vintage`, e.Vintage, 1900, true)
		}
		if e.Vintage > 2020 {
			err = goa.InvalidRangeError(`response[*].vintage`, e.Vintage, 2020, false)
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

		if e.Href == "" {
			err = goa.MissingAttributeError(`response[*]`, "href")
		}
		if e.Name == "" {
			err = goa.MissingAttributeError(`response[*]`, "name")
		}

		if e.Links != nil {
			if e.Links.Account != nil {

				if e.Links.Account.Href == "" {
					err = goa.MissingAttributeError(`response[*].links.account`, "href")
				}

			}
		}
		if len(e.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].name`, e.Name, len(e.Name), 2, true)
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 1, true)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 5, false)
			}
		}
	}
	return
}

// BottleLinksArray contains links to related resources of BottleCollection.
type BottleLinksArray []*BottleLinks

// Validate validates the BottleLinksArray type instance.
func (ut BottleLinksArray) Validate() (err error) {
	for _, e := range ut {
		if e.Account != nil {

			if e.Account.Href == "" {
				err = goa.MissingAttributeError(`response[*].account`, "href")
			}

		}
	}
	return
}
