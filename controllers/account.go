package controllers

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar/app"
)

// ToAccountMedia builds an account media type from an account model.
func ToAccountMedia(account *AccountModel) *app.Account {
	return &app.Account{
		CreatedAt: account.CreatedAt,
		CreatedBy: account.CreatedBy,
		Href:      account.Href,
		ID:        account.ID,
		Name:      account.Name,
	}
}

// ToAccountLink builds an account link from an account model.
func ToAccountLink(account *AccountModel) *app.AccountLink {
	return &app.AccountLink{
		Href: account.Href,
		ID:   account.ID,
	}
}

// AccountController implements the account resource.
type AccountController struct {
	goa.Controller
	db *DB
}

// NewAccount creates a account controller.
func NewAccount(service goa.Service) *AccountController {
	return &AccountController{
		Controller: service.NewController("Account"),
		db:         NewDB(),
	}
}

// Show retrieves the account with the given id.
func (b *AccountController) Show(c *app.ShowAccountContext) error {
	account := b.db.GetAccount(c.AccountID)
	if account == nil {
		return c.NotFound()
	}
	a := &app.Account{
		CreatedAt: account.CreatedAt,
		CreatedBy: account.CreatedBy,
		Href:      account.Href,
		ID:        account.ID,
		Name:      account.Name,
	}
	return c.OK(a)
}

// Create records a new account.
func (b *AccountController) Create(c *app.CreateAccountContext) error {
	account := b.db.NewAccount()
	payload := c.Payload
	account.Name = payload.Name
	c.ResponseData.Header().Set("Location", app.AccountHref(account.ID))
	return c.Created()
}

// Update updates a account field(s).
func (b *AccountController) Update(c *app.UpdateAccountContext) error {
	account := b.db.GetAccount(c.AccountID)
	if account == nil {
		return c.NotFound()
	}
	payload := c.Payload
	if payload.Name != "" {
		account.Name = payload.Name
	}
	b.db.SaveAccount(account)
	return c.NoContent()
}

// Delete removes a account from the database.
func (b *AccountController) Delete(c *app.DeleteAccountContext) error {
	account := b.db.GetAccount(c.AccountID)
	if account == nil {
		return c.NotFound()
	}
	b.db.DeleteAccount(account)
	return c.NoContent()
}
