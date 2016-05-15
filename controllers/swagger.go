package controllers

import "github.com/goadesign/goa"

// Swagger implements the swagger resource.
type Swagger struct {
	*goa.Controller
}

// NewSwagger creates a swagger controller.
func NewSwagger(service *goa.Service) *Swagger {
	return &Swagger{Controller: service.NewController("Swagger")}
}
