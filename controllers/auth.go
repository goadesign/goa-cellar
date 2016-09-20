package controllers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

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
	res := &app.Auth{Info: authInfo(ctx.Request)}
	return ctx.OK(res)
}

// JWT runs the jwt action.
func (c *AuthController) JWT(ctx *app.JWTAuthContext) error {
	res := &app.Auth{Info: authInfo(ctx.Request)}
	return ctx.OK(res)
}

// Extract auth info from security header created by Google Cloud Endpoint service.
func authInfo(req *http.Request) map[string]interface{} {
	info := req.Header.Get("X-Endpoint-API-UserInfo")
	if info == "" {
		return map[string]interface{}{"id": "anonymous"}
	}
	js, err := base64.StdEncoding.DecodeString(info)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}
	}
	var res map[string]interface{}
	err = json.Unmarshal(js, &res)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}
	}
	return res
}
