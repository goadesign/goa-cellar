//************************************************************************//
// API "cellar" version 2.0: Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/goa-cellar
// --design=github.com/goadesign/goa-cellar/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package v2

import "fmt"

// GenericBottleHref returns the resource href.
func GenericBottleHref(accountID, bottleID interface{}) string {
	return fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID)
}
