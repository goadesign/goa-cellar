package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// ShowAccount is the type used to define the shape of messages sent to retrieve an account.
var ShowAccount = Type("ShowAccount", func() {
	Attribute("ID", Integer, "Account ID", func() {
		Minimum(1) // ID parameter must be an integer greater than 1
	})
	Attribute("View", String, "The account view to retrieve", func() {
		Enum("default", "extended") // View parameter must be "default" or "extended"
	})
})

// BottlePayload defines the data structure used in the create bottle request body.
// It is also the base type for the bottle media type used to render bottles.
var BottlePayload = Type("BottlePayload", func() {
	Attribute("name", func() {
		MinLength(2)
		Example("Number 8")
	})
	Attribute("vineyard", func() {
		MinLength(2)
		Example("Asti")
	})
	Attribute("varietal", func() {
		MinLength(4)
		Example("Merlot")
	})
	Attribute("vintage", Integer, func() {
		Minimum(1900)
		Maximum(2020)
		Example(2012)
	})
	Attribute("color", func() {
		Enum("red", "white", "rose", "yellow", "sparkling")
	})
	Attribute("sweetness", Integer, func() {
		Minimum(1)
		Maximum(5)
	})
	Attribute("country", func() {
		MinLength(2)
		Example("USA")
	})
	Attribute("region", func() {
		Example("Napa Valley")
	})
	Attribute("review", func() {
		MinLength(3)
		MaxLength(300)
		Example("Great and inexpensive")
	})
})
