package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// The "files" service exposes file download endpoints.
// This service endpooints are intended for browsers, it lets users view the
// endpoints exposed by the API via the generated OpenAPI spec and Swagger UI.
//
// TBD: define CORS.
var _ = Service("files", func() {

	// Files defines a file server. The first argument is the URL to the
	// file and hte second the relative path on disk to the process current
	// directory.
	//
	// This expression causes the server to serve the file
	// "public/html/index.html" when clients make requests to "/ui".
	Files("/ui", "public/html/index.html")

	// The URL may use a final wildcard to signify that all files under the
	// directory can be downloaded. The file path is computed by appending
	// the value of the wildcard parameter to the directory base path.
	//
	// This expression causes the server to serve the file
	// "public/js/some/path.js" when clients make requests to
	// "/js/some/path.js".
	Files("/js/*filepath", "public/js")

	Files("/swagger.json", "public/swagger/swagger.json")
})
