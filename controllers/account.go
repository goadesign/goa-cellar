package controllers

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar/app"
)

// ToAccountMedia builds an account media type from an account model.
func ToAccountMedia(account *AccountModel) *app.Account {
	return &app.Account{
		ID:        account.ID,
		Href:      app.AccountHref(account.ID),
		Name:      account.Name,
		CreatedAt: account.CreatedAt,
		CreatedBy: account.CreatedBy,
	}
}

// ToAccountMediaTiny builds an account media type with tiny view from an account model.
func ToAccountMediaTiny(account *AccountModel) *app.AccountTiny {
	return &app.AccountTiny{
		ID:   account.ID,
		Href: app.AccountHref(account.ID),
		Name: account.Name,
	}
}

// ToAccountLink builds an account link from an account model.
func ToAccountLink(account *AccountModel) *app.AccountLink {
	return &app.AccountLink{
		ID:   account.ID,
		Href: app.AccountHref(account.ID),
	}
}

// AccountController implements the account resource.
type AccountController struct {
	*goa.Controller
	db *DB
}

// NewAccount creates a account controller.
func NewAccount(service *goa.Service) *AccountController {
	return &AccountController{
		Controller: service.NewController("Account"),
		db:         NewDB(),
	}
}

// List retrieves all the accounts.
func (b *AccountController) List(c *app.ListAccountContext) error {
	accounts := b.db.GetAccounts()
	res := make(app.AccountCollection, len(accounts))
	for i, account := range accounts {
		a := &app.Account{
			CreatedAt: account.CreatedAt,
			CreatedBy: account.CreatedBy,
			Href:      app.AccountHref(account.ID),
			ID:        account.ID,
			Name:      account.Name,
		}
		res[i] = a
	}
	return c.OK(res)
}

// Show retrieves the account with the given id.
func (b *AccountController) Show(c *app.ShowAccountContext) error {
	account, ok := b.db.GetAccount(c.AccountID)
	if !ok {
		return c.NotFound()
	}
	a := &app.Account{
		CreatedAt: account.CreatedAt,
		CreatedBy: account.CreatedBy,
		Href:      app.AccountHref(account.ID),
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
	account, ok := b.db.GetAccount(c.AccountID)
	if !ok {
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
	account, ok := b.db.GetAccount(c.AccountID)
	if !ok {
		return c.NotFound()
	}
	b.db.DeleteAccount(account)
	return c.NoContent()
}
