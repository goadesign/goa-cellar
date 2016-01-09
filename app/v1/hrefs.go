//************************************************************************//
// API "cellar" version 1.0: Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/raphael/goa-cellar
// --design=github.com/raphael/testd/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package v1

import "fmt"

// BottleHref returns the resource href.
func BottleHref(accountID, bottleID interface{}) string {
	return fmt.Sprintf("/1.0/cellar/accounts/%v/bottles/%v", accountID, bottleID)
}
