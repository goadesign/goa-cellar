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
	// Private type used to store auth handler info in request context
	authMiddlewareKey string
)

// UseAdminPassMiddleware mounts the admin_pass auth middleware onto the service.
func UseAdminPassMiddleware(service *goa.Service, middleware goa.Middleware) {
	service.Context = context.WithValue(service.Context, authMiddlewareKey("admin_pass"), middleware)
}

// NewAdminPassSecurity creates a admin_pass security definition.
func NewAdminPassSecurity() *goa.BasicAuthSecurity {
	def := goa.BasicAuthSecurity{}
	def.Description = "Basic authentication method, for global admin authentication.\n\nHere are very secret credentials:\n* username: wine\n* password: lover\n"
	return &def
}

// handleSecurity creates a handler that runs the auth middleware for the security scheme.
func handleSecurity(schemeName string, h goa.Handler, scopes ...string) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		scheme := ctx.Value(authMiddlewareKey(schemeName))
		am, ok := scheme.(goa.Middleware)
		if !ok {
			return goa.NoAuthMiddleware(schemeName)
		}
		ctx = goa.WithRequiredScopes(ctx, scopes)
		return am(h)(ctx, rw, req)
	}
}
