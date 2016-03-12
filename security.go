// +build !appengine

package main

import (
	"net/http"

	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// ErrAuthenticationFailed is an HTTP error that says you're not authenticated.
var ErrAuthenticationFailed = goa.NewErrorClass("auth_failed", "Invalid credentials, authentication failed", 401)

var adminPasswordMiddleware = goa.Middleware(func(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		//method := goa.SecurityMethod(ctx).(*goa.BasicAuthSecurity)

		username, password, ok := req.BasicAuth()
		if !ok || username != "wine" || password != "lover" {
			goa.Info(ctx, "login failed wine lover !")
			return ErrAuthenticationFailed.Errorf("Invalid credentials you wine lover !")
		}

		return h(context.WithValue(ctx, "logged_in", true), rw, req)
	}
})
