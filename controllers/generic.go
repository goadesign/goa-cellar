package controllers

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar/app"
	"github.com/goadesign/goa-cellar/app/v2"
)

// ToGenericBottleMedia converts a bottle model into a bottle media type
func ToGenericBottleMedia(b *BottleModel) *v2.GenericBottle {
	account := ToAccountMedia(b.Account)
	link := ToAccountLink(b.Account)
	return &v2.GenericBottle{
		Account:  account,
		Href:     b.Href,
		ID:       b.ID,
		Kind:     b.Kind,
		Links:    &v2.GenericBottleLinks{Account: link},
		Name:     b.Name,
		Rating:   b.Rating,
		Varietal: b.Varietal,
		Vineyard: b.Vineyard,
		Vintage:  b.Vintage,
	}
}

// GenericBottleController implements the bottle resource.
type GenericBottleController struct {
	*BottleController
	db *DB
}

// NewGeneric creates a generic bottle controller (API 2.0).
func NewGeneric(service *goa.Service) *GenericBottleController {
	return &GenericBottleController{
		BottleController: NewBottle(service),
		db:               NewDB(),
	}
}

// List lists all the bottles in the account optionally filtering by year.
func (b *GenericBottleController) List(ctx *v2.ListGenericBottleContext) error {
	var bottles []*BottleModel
	var err error
	if ctx.Years != nil {
		bottles, err = b.db.GetBottlesByYears(ctx.AccountID, ctx.Years)
	} else {
		bottles, err = b.db.GetBottles(ctx.AccountID)
	}
	if err != nil {
		return ctx.NotFound()
	}
	bs := make([]*v2.GenericBottle, len(bottles))
	for i, b := range bottles {
		bs[i] = ToGenericBottleMedia(b)
	}
	return ctx.OK(bs)
}

// Show retrieves the bottle with the given id.
func (b *GenericBottleController) Show(ctx *v2.ShowGenericBottleContext) error {
	bottle := b.db.GetBottle(ctx.AccountID, ctx.BottleID)
	if bottle == nil {
		return ctx.NotFound()
	}
	return ctx.OK(ToGenericBottleMedia(bottle))
}

// Create records a new bottle.
func (b *GenericBottleController) Create(ctx *v2.CreateGenericBottleContext) error {
	bottle := b.db.NewBottle(ctx.AccountID)
	payload := ctx.Payload
	bottle.Name = payload.Name
	bottle.Kind = payload.Kind
	bottle.Vintage = payload.Vintage
	bottle.Vineyard = payload.Vineyard
	if payload.Varietal != "" {
		bottle.Varietal = payload.Varietal
	}
	if payload.Color != "" {
		bottle.Color = payload.Color
	}
	if payload.Sweetness != nil {
		bottle.Sweetness = payload.Sweetness
	}
	if payload.Country != nil {
		bottle.Country = payload.Country
	}
	if payload.Region != nil {
		bottle.Region = payload.Region
	}
	if payload.Review != nil {
		bottle.Review = payload.Review
	}
	ctx.ResponseData.Header().Set("Location", app.BottleHref(ctx.AccountID, bottle.ID))
	return ctx.Created()
}

// Update updates a bottle field(s).
func (b *GenericBottleController) Update(ctx *v2.UpdateGenericBottleContext) error {
	bottle := b.db.GetBottle(ctx.AccountID, ctx.BottleID)
	if bottle == nil {
		return ctx.NotFound()
	}
	payload := ctx.Payload
	if payload.Name != nil {
		bottle.Name = *payload.Name
	}
	if payload.Kind != nil {
		bottle.Kind = *payload.Kind
	}
	if payload.Vintage != nil {
		bottle.Vintage = *payload.Vintage
	}
	if payload.Vineyard != nil {
		bottle.Vineyard = *payload.Vineyard
	}
	if payload.Varietal != nil {
		bottle.Varietal = *payload.Varietal
	}
	if payload.Color != nil {
		bottle.Color = *payload.Color
	}
	if payload.Sweetness != nil {
		bottle.Sweetness = payload.Sweetness
	}
	if payload.Country != nil {
		bottle.Country = payload.Country
	}
	if payload.Region != nil {
		bottle.Region = payload.Region
	}
	if payload.Review != nil {
		bottle.Review = payload.Review
	}
	b.db.SaveBottle(bottle)
	return ctx.NoContent()
}

// Delete removes a bottle from the database.
func (b *GenericBottleController) Delete(ctx *v2.DeleteGenericBottleContext) error {
	bottle := b.db.GetBottle(ctx.AccountID, ctx.BottleID)
	if bottle == nil {
		return ctx.NotFound()
	}
	b.db.DeleteBottle(bottle)
	return ctx.NoContent()
}

// Rate rates a bottle.
func (b *GenericBottleController) Rate(ctx *v2.RateGenericBottleContext) error {
	bottle := b.db.GetBottle(ctx.AccountID, ctx.BottleID)
	if bottle == nil {
		return ctx.NotFound()
	}
	bottle.Rating = &ctx.Payload.Rating
	b.db.SaveBottle(bottle)
	return ctx.NoContent()
}
