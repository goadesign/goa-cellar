package controllers

import (
	"fmt"
	"io"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar/app"
	"golang.org/x/net/websocket"
)

// ToBottleMedia converts a bottle model into a bottle media type
func ToBottleMedia(b *BottleModel) *app.Bottle {
	account := ToAccountMedia(b.Account)
	link := ToAccountLink(b.Account)
	return &app.Bottle{
		Account:  account,
		Href:     b.Href,
		ID:       b.ID,
		Links:    &app.BottleLinks{Account: link},
		Name:     b.Name,
		Rating:   b.Rating,
		Varietal: b.Varietal,
		Vineyard: b.Vineyard,
		Vintage:  b.Vintage,
	}
}

// BottleController implements the bottle resource.
type BottleController struct {
	*goa.Controller
	db *DB
}

// NewBottle creates a bottle controller.
func NewBottle(service *goa.Service) *BottleController {
	return &BottleController{
		Controller: service.NewController("Bottle"),
		db:         NewDB(),
	}
}

// List lists all the bottles in the account optionally filtering by year.
func (b *BottleController) List(ctx *app.ListBottleContext) error {
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
	bs := make([]*app.Bottle, len(bottles))
	for i, b := range bottles {
		bs[i] = ToBottleMedia(b)
	}
	return ctx.OK(bs)
}

// Show retrieves the bottle with the given id.
func (b *BottleController) Show(ctx *app.ShowBottleContext) error {
	bottle := b.db.GetBottle(ctx.AccountID, ctx.BottleID)
	if bottle == nil {
		return ctx.NotFound()
	}
	return ctx.OK(ToBottleMedia(bottle))
}

// Watch watches the bottle with the given id.
func (b *BottleController) Watch(ctx *app.WatchBottleContext) error {
	Watcher(ctx.AccountID, ctx.BottleID).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// Echo the data received on the WebSocket.
func Watcher(accountID, bottleID int) websocket.Handler {
	return func(ws *websocket.Conn) {
		watched := fmt.Sprintf("Account: %d, Bottle: %d", accountID, bottleID)
		ws.Write([]byte(watched))
		io.Copy(ws, ws)
	}
}

// Create records a new bottle.
func (b *BottleController) Create(ctx *app.CreateBottleContext) error {
	bottle := b.db.NewBottle(ctx.AccountID)
	payload := ctx.Payload
	bottle.Name = payload.Name
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
func (b *BottleController) Update(ctx *app.UpdateBottleContext) error {
	bottle := b.db.GetBottle(ctx.AccountID, ctx.BottleID)
	if bottle == nil {
		return ctx.NotFound()
	}
	payload := ctx.Payload
	if payload.Name != nil {
		bottle.Name = *payload.Name
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
func (b *BottleController) Delete(ctx *app.DeleteBottleContext) error {
	bottle := b.db.GetBottle(ctx.AccountID, ctx.BottleID)
	if bottle == nil {
		return ctx.NotFound()
	}
	b.db.DeleteBottle(bottle)
	return ctx.NoContent()
}

// Rate rates a bottle.
func (b *BottleController) Rate(ctx *app.RateBottleContext) error {
	bottle := b.db.GetBottle(ctx.AccountID, ctx.BottleID)
	if bottle == nil {
		return ctx.NotFound()
	}
	bottle.Rating = &ctx.Payload.Rating
	b.db.SaveBottle(bottle)
	return ctx.NoContent()
}
