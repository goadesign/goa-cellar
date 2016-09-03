package controllers

import (
	"fmt"
	"time"

	"github.com/goadesign/goa-cellar/store"
)

const (
	createdBy         = "test@goa.design"
	accountNameFormat = "Account #%d"
	bottleNameFormat  = "Bottle #%d"
	vineyard          = "vineyard"
	vintage           = 2012
	color             = "red"
)

var (
	createdAt = time.Now()
	kind      = "wine"
	sweetness = 1
	country   = "usa"
	region    = "ca"
	review    = "review"
	rating    = 4
	varietal  = "pinot noir"
)

func createAccount(db *store.DB, idx int) *store.AccountModel {
	a := db.NewAccount()
	a.Name = fmt.Sprintf(accountNameFormat, a.ID)
	a.CreatedAt = createdAt
	a.CreatedBy = createdBy
	db.SaveAccount(a)
	return &a
}
