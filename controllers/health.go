package controllers

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar-ep/app"
	"github.com/goadesign/goa-cellar-ep/store"
)

// HealthController implements the account resource.
type HealthController struct {
	*goa.Controller
	db *store.DB
}

// NewHealth creates a account controller.
func NewHealth(service *goa.Service, db *store.DB) *HealthController {
	return &HealthController{
		Controller: service.NewController("Health"),
		db:         db,
	}
}

// Health performs the health-check.
func (b *HealthController) Health(c *app.HealthHealthContext) error {
	return c.OK([]byte("ok"))
}
