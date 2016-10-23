package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// The "account" service exposes the account resource endpoints.
var _ = Service("account", func() {

	// Responses that do not explicitly define a type use Account.  Also
	// request parameters named after an attribute of Account automatically
	// inherit the attribute properties (type, description, validations).
	DefaultType(Account)

	//  HTTP transport specific expressions
	HTTP(func() {

		// Common path prefix to all the service endpoints. Individual
		// enpoints may define absolute paths by prefixing them with
		// "//".
		BasePath("/accounts")
	})

	// "list" account service endpoint
	Endpoint("list", func() {

		Description("Retrieve all accounts.")

		// Request type lists the request parameters via its attributes.
		Request(ListAccount)

		// Response type is a collection of Account.
		Response(CollectionOf(Account))

		// HTTP transport specific expressions
		HTTP(func() {

			// GET requests with paths "/accounts" invoke this
			// endpoint.
			GET("")

			// The "view" query string value defines the view used
			// to render the collection. The parameter is defined by
			// the "view" attribute of the ListAccount type.
			Param("view")

			// Response with status code 200 uses the endpoint
			// response type to define the response body and
			// content-type header.
			Response(OK)
		})
	})

	// "show" account service endpoint
	Endpoint("show", func() {

		Description("Retrieve account with given id. IDs 1 and 2 pre-exist in the system.")

		// Request type lists the request parameters via its attributes.
		Request(ShowAccount)

		// HTTP transport specific expressions
		HTTP(func() {

			// GET requests with paths of the form
			// "/accounts/{accountID}" where accountID is a path
			// parameter defined via the "accountID" attribute of
			// the ShowAccount type.
			GET("/{accountID}")

			// The "view" query string value defines the view used
			// to render the account. The parameter is defined by
			// the "view" attribute of the ShowAccount type.
			Param("view")

			// The response with HTTP status code 200 uses the
			// default service type to describe the response body.
			// Since the type Account is a media type it also
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

	// "create" account service endpoint
	Endpoint("create", func() {

		Description("Create new account")

		// Define request type inline
		Request(func() {

			// "name" attribute inherits type, description and
			// validations from Account media type "name" attribute.
			Attribute("name")

			// Make "name" attribute required. Requests whose body
			// do not specify a value for "name" are considered bad
			// requests.
			Required("name")
		})

		// Define response type inline as well.
		Response(func() {

			// "ID" attribute specifies the ID of the newly created
			// account. It inherits the type, description and
			// validations from the Account media type "ID"
			// attribute.
			Attribute("ID")

			// "Location" attribute specifies a URI for the newly
			// created account.
			Attribute("Location", String, "URI to newly created account", func() {

				// Validation regular expression
				Pattern("/accounts/[0-9]+")
			})
		})

		// HTTP transport specific expressions
		// Note that the request body is implicitly defined by the
		// Request type described above.
		HTTP(func() {

			// POST requests with paths "/accounts" invoke this
			// endpoint.
			POST("")

			// Response with status code 201. The Location response
			// attribute is transferred via a header.
			Response(Created, func() {

				// Refers to response type "Location" attribute.
				Header("Location")

				// Refers to response type "ID" attribute.
				Attribute("ID")
			})

			// Response with status code 400.
			Response(BadRequest, ErrorMedia)
		})
	})

	// "update" account service endpoint
	Endpoint("update", func() {

		Description("Change account name")

		Request(UpdateAccount)

		// Response defined using the built-in EmptyMedia media type.
		Response(EmptyMedia)

		// HTTP transport specific expressions
		HTTP(func() {

			// PUT requests with paths of the form
			// "/accounts/{accountID}" invoke this endpoint.
			PUT("/{accountID}")

			// Payload defines the request body. It uses the
			// attributes defined on the request type.
			// Defining the payload explicitly is required here
			// because it does not match the content of the request
			// type (the "accountID" attribute defines a parameter).
			Payload(func() {
				Attribute("name")
				Required("name")
			})

			// Response with status code 204, empty body.
			Response(NoContent)

			// Response with status code 404, empty body.
			Response(NotFound)

			// Response with status code 400.
			Response(BadRequest, ErrorMedia)
		})
	})

	// "delete" account service endpoint
	Endpoint("delete", func() {
		Description("Delete account")
		Request(DeleteAccount)
		Response(EmptyMedia)

		HTTP(func() {
			DELETE("/{accountID}")
			Response(NoContent)
			Response(NotFound)
			Response(BadRequest, ErrorMedia)
		})
	})
})
