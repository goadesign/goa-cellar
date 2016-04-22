//************************************************************************//
// Resource Hrefs
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/goa-cellar
// --design=github.com/goadesign/goa-cellar/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import "fmt"

// accountHref returns the account resource href.
func accountHref(accountID interface{}) string {
	return fmt.Sprintf("/cellar/accounts/%v", accountID)
}

// bottleHref returns the bottle resource href.
func bottleHref(accountID, bottleID interface{}) string {
	return fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID)
}
