//************************************************************************//
// cellar JavaScript Client Example
//
// Generated with goagen v0.0.1, command line:
// $ goagen.exe
// --out=$(GOPATH)\src\github.com\goadesign\goa-cellar
// --design=github.com/goadesign/goa-cellar/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package js

import "github.com/goadesign/goa"

// MountController mounts the JavaScript example controller under "/js".
func MountController(service *goa.Service) {
	// Serve static files under js
	service.ServeFiles("/js/*filepath", "C:\\Workspace\\src\\github.com\\goadesign\\goa-cellar\\js")
	service.Info("mount", "ctrl", "JS", "action", "ServeFiles", "route", "GET /js/*")
}
