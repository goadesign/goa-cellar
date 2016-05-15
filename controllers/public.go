package controllers

import "github.com/goadesign/goa"

// Public implements the public resource.
type Public struct {
	*goa.Controller
}

// NewPublic creates a public controller.
func NewPublic(service *goa.Service) *Public {
	return &Public{Controller: service.NewController("Public")}
}
