package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// The "bottle" service exposes the bottle resource endpoints.
var _ = Service("bottle", func() {

	// Responses that do not explicitly define a type use Bottle.  Also
	// request parameters named after an attribute of Bottle automatically
	// inherit the attribute properties (type, description, validations).
	DefaultType(Bottle)

	//  HTTP transport specific expressions
	HTTP(func() {

		// Common path prefix to all the service endpoints. Individual
		// enpoints may define absolute paths by prefixing them with
		// "//".
		BasePath("/bottles")

		// Identifies the parent resource. This causes all
		// (non-absolute) request paths to be prefixed with the parent
		// resource canonical endpoint request path. The canonical
		// endpoint of a resource is "show". It can be overridden via
		// the "CanonicalEndpoint" DSL.
		Parent("account")
	})

	// "list" bottle service endpoint
	Endpoint("list", func() {

		Description("List all bottles in account optionally filtering by year")

		// Request type lists the request parameters via its attributes.
		Request(ListBottle)

		// Response type is a collection of Bottle.
		Response(CollectionOf(Bottle))

		// HTTP transport specific expressions
		HTTP(func() {

			// GET requests with paths of the form
			// "/accounts/{accountID}/bottles" invoke this endpoint.
			GET("")

			// Define "years" query string parameter. The parameter
			// type, description and validations are inherited from
			// the corresponding ListBottle attribute.
			Param("years")

			// Response with status code 200 uses the endpoint
			// response type to define the response body and
			// content-type header.
			Response(OK)

			// The response with HTTP status code 404. No type means
			// empty body (only the OK response uses the response
			// type by default). This response is returned when
			// there is no account whose ID matches the "accountID"
			// path parameter.
			Response(NotFound)

			// The response with HTTP status code 400. The body and
			// content type header are defined by the ErrorMedia
			// built-in media type. This response is returned when
			// the "accountID" path parameter is malformed.
			Response(BadRequest, ErrorMedia)
		})
	})

	// "show" bottle service endpoint
	Endpoint("show", func() {

		Description("Retrieve bottle with given id")

		// Request type lists the request parameters via its attributes.
		Request(ShowBottle)

		HTTP(func() {

			// GET requests with paths of the form
			// "/accounts/{accountID}/bottles/{bottleID}" where
			// accountID and bottleID are path parameters defined
			// via the "accountID" and "bottleID" attributes of the
			// ShowAccount and ShowBottle types.
			GET("/{bottleID}")

			// The response with HTTP status code 200 uses the
			// default service type to describe the response body.
			// Since the type Bottle is a media type it also
			// defines the content-type response header.
			Response(OK)

			// The response with HTTP status code 404. No type means
			// empty body (only the OK response uses the service
			// default type).
			Response(NotFound)

			// The response with HTTP status code 400. The body and
			// content type header are defined by the ErrorMedia
			// built-in media type.
			Response(BadRequest, ErrorMedia)
		})
	})

	Endpoint("watch", func() {

		Description("Subscribe to changes made to the bottle")

		// Request type lists the request parameters via its attributes.
		Request(ShowBottle)

		// Stream defines a streaming response implemented via websocket
		// when using the HTTP transport, gRPC streaming responses
		// when using gRPC.
		Response(Stream)

		HTTP(func() {

			// GET requests with paths of the form
			// "/accounts/{accountID}/bottles/{bottleID}/watch"
			// where accountID and bottleID are path parameters
			// defined via the "accountID" and "bottleID" attributes
			// of the ShowAccount and ShowBottle types.
			GET("/{bottleID}/watch")

			// The response with HTTP status code 101.
			Response(SwitchingProtocols)

			// The response with HTTP status code 404. No type means
			// empty body (only the OK response uses the service
			// default type).
			Response(NotFound)

			// The response with HTTP status code 400. The body and
			// content type header are defined by the ErrorMedia
			// built-in media type.
			Response(BadRequest, ErrorMedia)
		})
	})

	Endpoint("create", func() {

		Description("Record new bottle")

		Request(BottlePayload)

		// Define response type inline.
		Response(func() {

			// "ID" attribute specifies the ID of the newly created
			// bottle. It inherits the type, description and
			// validations from the Bottle media type "ID"
			// attribute.
			Attribute("ID")

			// "Location" attribute specifies a URI for the newly
			// created bottle.
			Attribute("Location", String, "URI to newly created bottle", func() {

				// Validation regular expression
				Pattern("/accounts/[0-9]+/bottles/[0-9]+")
			})
		})

		// HTTP transport specific expressions
		// Note that the request body is implicitly defined by the
		// Request type described above.
		HTTP(func() {

			// POST requests with paths
			// "/accounts/{accountID}/bottles" invoke this endpoint.
			POST("")

			// Response with status code 201. The Location response
			// attribute is transferred via a header.
			Response(Created, func() {

				// Refers to response type "Location" attribute.
				Header("Location")

				// Refers to response type "ID" attribute.
				Attribute("ID")
			})

			// Response with status code 404.
			Response(NotFound)

			// Response with status code 400.
			Response(BadRequest, ErrorMedia)
		})
	})

	Endpoint("update", func() {
		Description("Update bottle attributes with patch semantic")
		Request(BottlePayload)
		Response(EmptyMedia)

		HTTP(func() {
			PATCH("/{bottleID}")
			Response(NoContent)
			Response(NotFound)
			Response(BadRequest, ErrorMedia)
		})
	})

	Endpoint("rate", func() {

		Description("Rate bottle")

		// Define request type inline
		Request(func() {
			// Attribute "rating" as defined in service default
			// type.
			Attribute("rating")
			Required("rating")
		})

		Response(EmptyMedia)

		HTTP(func() {
			PUT("/{bottleID}/rating")
			Response(NoContent)
			Response(NotFound)
			Response(BadRequest, ErrorMedia)
		})
	})

	Endpoint("delete", func() {
		Description("Delete bottle")
		Request(DeleteBottle)
		Response(EmptyMedia)

		HTTP(func() {
			DELETE("/{bottleID}")
			Response(NoContent)
			Response(NotFound)
			Response(BadRequest, ErrorMedia)
		})
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
