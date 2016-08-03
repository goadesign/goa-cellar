//************************************************************************//
// API "cellar": Application Resource Href Factories
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/goadesign/goa-cellar/design
// --out=$(GOPATH)/src/github.com/goadesign/goa-cellar
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// AccountHref returns the resource href.
func AccountHref(accountID interface{}) string {
	return fmt.Sprintf("/cellar/accounts/%v", accountID)
}

// BottleHref returns the resource href.
func BottleHref(accountID, bottleID interface{}) string {
	return fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID)
}
