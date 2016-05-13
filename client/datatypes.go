//************************************************************************//
// User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/goadesign/goa-cellar/design
// --out=$(GOPATH)/src/github.com/goadesign/goa-cellar
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"io"
	"time"
)

// A tenant account
type Account struct {
	// Date of creation
	CreatedAt *time.Time `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Email of account owner
	CreatedBy *string `json:"created_by,omitempty" xml:"created_by,omitempty"`
	// API href of account
	Href string `json:"href" xml:"href"`
	// ID of account
	ID int `json:"id" xml:"id"`
	// Name of account
	Name string `json:"name" xml:"name"`
}

// DecodeAccount decodes the Account instance encoded in r.
func DecodeAccount(r io.Reader, decoderFn goa.DecoderFunc) (*Account, error) {
	var decoded Account
	err := decoderFn(r).Decode(&decoded)
	return &decoded, err
}

// BottleCollection media type is a collection of Bottle.
type BottleCollection []*Bottle

// DecodeBottleCollection decodes the BottleCollection instance encoded in r.
func DecodeBottleCollection(r io.Reader, decoderFn goa.DecoderFunc) (BottleCollection, error) {
	var decoded BottleCollection
	err := decoderFn(r).Decode(&decoded)
	return decoded, err
}

// A bottle of wine
type Bottle struct {
	// Account that owns bottle
	Account *Account `json:"account,omitempty" xml:"account,omitempty"`
	Color   string   `json:"color" xml:"color"`
	Country *string  `json:"country,omitempty" xml:"country,omitempty"`
	// Date of creation
	CreatedAt *time.Time `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// API href of bottle
	Href string `json:"href" xml:"href"`
	// ID of bottle
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating    *int    `json:"rating,omitempty" xml:"rating,omitempty"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	// Date of last update
	UpdatedAt *time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	Varietal  string     `json:"varietal" xml:"varietal"`
	Vineyard  string     `json:"vineyard" xml:"vineyard"`
	Vintage   int        `json:"vintage" xml:"vintage"`
}

// DecodeBottle decodes the Bottle instance encoded in r.
func DecodeBottle(r io.Reader, decoderFn goa.DecoderFunc) (*Bottle, error) {
	var decoded Bottle
	err := decoderFn(r).Decode(&decoded)
	return &decoded, err
}
