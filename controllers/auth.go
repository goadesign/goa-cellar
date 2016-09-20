package controllers

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar-ep/app"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
}

// NewAuthController creates a auth controller.
func NewAuth(service *goa.Service) *AuthController {
	return &AuthController{Controller: service.NewController("AuthController")}
}

// Basic runs the basic action.
func (c *AuthController) Basic(ctx *app.BasicAuthContext) error {
	res := &app.Auth{Info: ctx.Value("userinfo").(map[string]interface{})}
	return ctx.OK(res)
}

// JWT runs the jwt action.
func (c *AuthController) JWT(ctx *app.JWTAuthContext) error {
	res := &app.Auth{Info: ctx.Value("userinfo").(map[string]interface{})}
	return ctx.OK(res)
}
