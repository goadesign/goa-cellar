//************************************************************************//
// API "cellar": Application Security
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/goa-cellar
// --design=github.com/goadesign/goa-cellar/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"errors"
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

type securitySchemeKey string
type key int

const securityScopesKey key = 1

func ConfigureAdminPassSecurity(service *goa.Service, f goa.BasicAuthSecurityConfigFunc) {
	def := &goa.BasicAuthSecurity{}
	def.Description = "Basic authentication method, for global admin authentication.\n\nHere are very secret credentials:\n* username: wine\n* password: lover\n"

	middleware := f(def)

	service.Context = context.WithValue(service.Context, securitySchemeKey("admin_pass"), middleware)
}

func handleSecurity(service *goa.Service, schemeName string, h goa.Handler, scopes ...string) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		scheme := service.Context.Value(securitySchemeKey(schemeName))
		middleware, ok := scheme.(goa.Middleware)
		if !ok {
			goa.RequestService(ctx).Error("security scheme not configured")
			return errors.New("security scheme not configured")
		}

		if len(scopes) != 0 {
			ctx = context.WithValue(ctx, securityScopesKey, scopes)
		}

		return middleware(h)(ctx, rw, req)
	}
}
