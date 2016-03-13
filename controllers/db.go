package controllers

import (
	"fmt"
	"sync"
	"time"

	"github.com/goadesign/goa-cellar/app"
)

var (
	one   = 1
	three = 3
	four  = 4
	five  = 5
	usa   = "USA"
	ca    = "CA"
	gv    = "Great value"
	gbe   = "Good but expensive"
	ok    = "OK"
	fav   = "Favorite"
	solid = "Solid Pinot"
)

// DB emulates a database driver using in-memory data structures.
type DB struct {
	sync.Mutex
	maxAccountModelID int
	accounts          map[int]*AccountModel
	bottles           map[int][]*BottleModel
}

// BottleModel is the database "model" for bottles
type BottleModel struct {
	ID        int
	Href      string
	Account   *AccountModel
	Name      string
	Kind      string
	Color     string
	Country   *string
	CreatedAt *time.Time
	Rating    *int
	Region    *string
	Review    *string
	Sweetness *int
	UpdatedAt *time.Time
	Varietal  string
	Vineyard  string
	Vintage   int
}

// AccountModel is the database "model" for accounts
type AccountModel struct {
	ID        int
	Href      string
	Name      string
	CreatedAt *time.Time
	CreatedBy *string
}

// NewDB initializes a new "DB" with dummy data.
func NewDB() *DB {
	account := &AccountModel{ID: 1, Name: "account 1", Href: app.AccountHref(1)}
	account2 := &AccountModel{ID: 2, Name: "account 2", Href: app.AccountHref(2)}
	bottles := map[int][]*BottleModel{
		1: []*BottleModel{
			&BottleModel{
				ID:        100,
				Account:   account,
				Href:      app.BottleHref(1, 100),
				Name:      "Number 8",
				Kind:      "wine",
				Vineyard:  "Asti Winery",
				Varietal:  "Merlot",
				Vintage:   2012,
				Color:     "red",
				Sweetness: &one,
				Country:   &usa,
				Region:    &ca,
				Review:    &gv,
				Rating:    &four,
			},
			&BottleModel{
				ID:        101,
				Account:   account,
				Href:      app.BottleHref(1, 101),
				Name:      "Mourvedre",
				Kind:      "wine",
				Vineyard:  "Rideau",
				Varietal:  "Mourvedre",
				Vintage:   2012,
				Color:     "red",
				Sweetness: &one,
				Country:   &usa,
				Region:    &ca,
				Review:    &gbe,
				Rating:    &three,
			},
			&BottleModel{
				ID:        102,
				Account:   account,
				Href:      app.BottleHref(1, 102),
				Name:      "Blue's Cuvee",
				Kind:      "wine",
				Vineyard:  "Longoria",
				Varietal:  "Cabernet Franc with Merlot, Malbec, Cabernet Sauvignon and Syrah",
				Vintage:   2012,
				Color:     "red",
				Sweetness: &one,
				Country:   &usa,
				Region:    &ca,
				Review:    &fav,
				Rating:    &five,
			},
		},
		2: []*BottleModel{
			&BottleModel{
				ID:        200,
				Account:   account2,
				Href:      app.BottleHref(42, 200),
				Name:      "Blackstone Merlot",
				Kind:      "wine",
				Vineyard:  "Blackstone",
				Varietal:  "Merlot",
				Vintage:   2012,
				Color:     "red",
				Sweetness: &one,
				Country:   &usa,
				Region:    &ca,
				Review:    &ok,
				Rating:    &three,
			},
			&BottleModel{
				ID:        201,
				Account:   account2,
				Href:      app.BottleHref(42, 201),
				Name:      "Wild Horse",
				Kind:      "wine",
				Vineyard:  "Wild Horse",
				Varietal:  "Pinot Noir",
				Vintage:   2012,
				Color:     "red",
				Sweetness: &one,
				Country:   &usa,
				Region:    &ca,
				Review:    &solid,
				Rating:    &four,
			},
		},
	}
	return &DB{accounts: map[int]*AccountModel{1: account, 2: account2}, bottles: bottles, maxAccountModelID: 2}
}

// GetAccount returns the account with given id if any, nil otherwise.
func (db *DB) GetAccount(id int) *AccountModel {
	db.Lock()
	defer db.Unlock()
	return db.accounts[id]
}

// NewAccount creates a new blank account resource.
func (db *DB) NewAccount() *AccountModel {
	db.Lock()
	defer db.Unlock()
	db.maxAccountModelID++
	account := &AccountModel{ID: db.maxAccountModelID}
	db.accounts[db.maxAccountModelID] = account
	return account
}

// SaveAccount "persists" the account.
func (db *DB) SaveAccount(a *AccountModel) {
	db.Lock()
	defer db.Unlock()
	db.accounts[a.ID] = a
}

// DeleteAccount deletes the account.
func (db *DB) DeleteAccount(account *AccountModel) {
	db.Lock()
	defer db.Unlock()
	if account == nil {
		return
	}
	delete(db.bottles, account.ID)
	delete(db.accounts, account.ID)
}

// GetBottle returns the bottle with the given id from the given account or nil if not found.
func (db *DB) GetBottle(account, id int) *BottleModel {
	db.Lock()
	defer db.Unlock()
	bottles, ok := db.bottles[account]
	if !ok {
		return nil
	}
	for _, b := range bottles {
		if b.ID == id {
			return b
		}
	}
	return nil
}

// GetBottles return the bottles from the given account.
func (db *DB) GetBottles(account int) ([]*BottleModel, error) {
	db.Lock()
	defer db.Unlock()
	bottles, ok := db.bottles[account]
	if !ok {
		return nil, fmt.Errorf("unknown account %d", account)
	}
	return bottles, nil
}

// GetBottlesByYears returns the bottles with the vintage in the given array from the given account.
func (db *DB) GetBottlesByYears(account int, years []int) ([]*BottleModel, error) {
	db.Lock()
	defer db.Unlock()
	bottles, ok := db.bottles[account]
	if !ok {
		return nil, fmt.Errorf("unknown account %d", account)
	}
	var res []*BottleModel
	for _, b := range bottles {
		selected := false
		for _, y := range years {
			if y == b.Vintage {
				selected = true
				break
			}
		}
		if selected {
			res = append(res, b)
		}
	}
	return res, nil
}

// NewBottle creates a new bottle resource.
func (db *DB) NewBottle(account int) *BottleModel {
	db.Lock()
	defer db.Unlock()
	bottles, _ := db.bottles[account]
	newID := 1
	taken := true
	for ; taken; newID++ {
		taken = false
		for _, b := range bottles {
			if b.ID == newID {
				taken = true
				break
			}
		}
	}
	bottle := BottleModel{ID: newID}
	db.bottles[newID] = append(db.bottles[newID], &bottle)
	return &bottle
}

// SaveBottle persists bottle to bottlesbase.
func (db *DB) SaveBottle(b *BottleModel) {
	db.Lock()
	defer db.Unlock()
	db.bottles[b.Account.ID] = append(db.bottles[b.Account.ID], b)
}

// DeleteBottle deletes bottle from bottlesbase.
func (db *DB) DeleteBottle(bottle *BottleModel) {
	db.Lock()
	defer db.Unlock()
	if bottle == nil {
		return
	}
	account := bottle.Account
	id := bottle.ID
	if bs, ok := db.bottles[account.ID]; ok {
		idx := -1
		for i, b := range bs {
			if b.ID == id {
				idx = i
				break
			}
		}
		if idx > -1 {
			bs = append(bs[:idx], bs[idx+1:]...)
			db.bottles[account.ID] = bs
		}
	}
}
