//************************************************************************//
// cellar JSON Hyper-schema
//
// Generated with goagen v0.0.1, command line:
// $ goagen.exe
// --out=$(GOPATH)\src\github.com\goadesign\goa-cellar
// --design=github.com/goadesign/goa-cellar/design
// --url=http://localhost
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package schema

import "github.com/goadesign/goa"

// MountController mounts the API JSON schema controller under "/schema.json".
func MountController(service *goa.Service) {
	service.ServeFiles("/schema.json", "schema/schema.json")
}
