package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Service("account", func() {

	HTTP(func() {
		DefaultMedia(Account)
		BasePath("/accounts")
	})

	Endpoint("list", func() {
		Description("Retrieve all accounts.")
		Response(CollectionOf(Account))

		HTTP(func() { GET("") })
	})

	Endpoint("show", func() {
		Description("Retrieve account with given id. IDs 1 and 2 pre-exist in the system.")
		Routing(
			GET("/:accountID"),
		)
		Params(func() {
			Param("accountID", Integer, "Account ID", func() {
				Minimum(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Endpoint("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new account")
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(Created, "/accounts/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Endpoint("update", func() {
		Routing(
			PUT("/:accountID"),
		)
		Description("Change account name")
		Params(func() {
			Param("accountID", Integer, "Account ID")
		})
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Endpoint("delete", func() {
		Routing(
			DELETE("/:accountID"),
		)
		Params(func() {
			Param("accountID", Integer, "Account ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Service("bottle", func() {

	DefaultMedia(Bottle)
	BasePath("bottles")
	Parent("account")

	Endpoint("list", func() {
		Routing(
			GET(""),
		)
		Description("List all bottles in account optionally filtering by year")
		Params(func() {
			Param("years", ArrayOf(Integer), "Filter by years")
		})
		Response(OK, func() {
			Media(CollectionOf(Bottle, func() {
				View("default")
				View("tiny")
			}))
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Endpoint("show", func() {
		Routing(
			GET("/:bottleID"),
		)
		Description("Retrieve bottle with given id")
		Params(func() {
			Param("bottleID", Integer)
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Endpoint("watch", func() {
		Routing(
			GET("/:bottleID/watch"),
		)
		Scheme("ws")
		Description("Retrieve bottle with given id")
		Params(func() {
			Param("bottleID", Integer)
		})
		Response(SwitchingProtocols)
		Response(BadRequest, ErrorMedia)
	})

	Endpoint("create", func() {
		Routing(
			POST(""),
		)
		Description("Record new bottle")
		Payload(BottlePayload, func() {
			Required("name", "vineyard", "varietal", "vintage", "color")
		})
		Response(Created, "^/accounts/[0-9]+/bottles/[0-9]+$")
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Endpoint("update", func() {
		Routing(
			PATCH("/:bottleID"),
		)
		Params(func() {
			Param("bottleID", Integer)
		})
		Payload(BottlePayload)
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Endpoint("rate", func() {
		Routing(
			PUT("/:bottleID/actions/rate"),
		)
		Params(func() {
			Param("bottleID", Integer)
		})
		Payload(func() {
			Member("rating")
			Required("rating")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Endpoint("delete", func() {
		Routing(
			DELETE("/:bottleID"),
		)
		Params(func() {
			Param("bottleID", Integer)
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Service("public", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/ui", "public/html/index.html")
})

var _ = Service("js", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/js/*filepath", "public/js")
})

var _ = Service("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/swagger.json", "public/swagger/swagger.json")
})
