package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// BottlePayload defines the data structure used in the create bottle request body.
// It is also the base type for the bottle media type used to render bottles.
var BottlePayload = Type("BottlePayload", func() {
	Attribute("name", func() {
		MinLength(2)
	})
	Attribute("vineyard", func() {
		MinLength(2)
	})
	Attribute("varietal", func() {
		MinLength(4)
	})
	Attribute("vintage", Integer, func() {
		Minimum(1900)
		Maximum(2020)
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
	})
	Attribute("region")
	Attribute("review", func() {
		MinLength(3)
		MaxLength(300)
	})
})

// GenericBottlePayload defines the data structure used in the create bottle request body for API
// version 2.0.
var GenericBottlePayload = Type("GenericBottlePayload", func() {
	Reference(BottlePayload)
	APIVersion("2.0")
	Attribute("name")
	Attribute("kind", String, "Bottle kind", func() {
		Enum("wine", "rum")
	})
	Attribute("vineyard")
	Attribute("varietal")
	Attribute("vintage")
	Attribute("color")
	Attribute("sweetness")
	Attribute("country")
	Attribute("region")
	Attribute("review")
})
