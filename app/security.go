//************************************************************************//
// API "cellar": Application Security
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/goadesign/goa-cellar/design
// --out=$(GOPATH)/src/github.com/goadesign/goa-cellar
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

type (
	// Private type used to store scheme info in request context
	securitySchemeKey string
	// Private type used to store scopes in request context
	key int
)

// Scopes context key
const securityScopesKey key = 1

// ConfigureAdminPassSecurity configures the admin_pass security scheme.
// It accepts a BasicAuthSecurityConfigFunc and calls it giving it the scheme definition as defined in
// the DSL as well as a fetchScopes function object that allows retrieving the request scopes. The
// function should return a middleware that the generated code invokes on every request.
func ConfigureAdminPassSecurity(service *goa.Service, f goa.BasicAuthSecurityConfigFunc) {
	def := &goa.BasicAuthSecurity{}
	def.Description = "Basic authentication method, for global admin authentication.\n\nHere are very secret credentials:\n* username: wine\n* password: lover\n"

	middleware := f(def)
	service.Context = context.WithValue(service.Context, securitySchemeKey("admin_pass"), middleware)
}

// handleSecurity creates a goa request handler that takes care of executing the security middleware
// registered via the ConfigureXXXSecurity functions.
func handleSecurity(schemeName string, h goa.Handler, scopes ...string) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		scheme := ctx.Value(securitySchemeKey(schemeName))
		middleware, ok := scheme.(goa.Middleware)
		if !ok {
			return goa.NoSecurityScheme(schemeName)
		}

		if len(scopes) != 0 {
			ctx = context.WithValue(ctx, securityScopesKey, scopes)
		}

		return middleware(h)(ctx, rw, req)
	}
}
