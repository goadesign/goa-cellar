package main

import (
	"fmt"
	"github.com/goadesign/goa-cellar/client"
	"github.com/goadesign/goa-cellar/tool/cli"
	goaclient "github.com/goadesign/goa/client"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"time"
)

func main() {
	// Create command line parser
	app := &cobra.Command{
		Use:   "cellar-cli",
		Short: `CLI client for the cellar service (http://goa.design/getting-started.html)`,
	}

	// Create client struct
	c := client.New(newHTTPClient())

	// Register global flags
	app.PersistentFlags().StringVarP(&c.Scheme, "scheme", "s", "", "Set the requests scheme")
	app.PersistentFlags().StringVarP(&c.Host, "host", "H", "cellar.goa.design", "API hostname")
	app.PersistentFlags().DurationVarP(&c.Timeout, "timeout", "t", time.Duration(20)*time.Second, "Set the request timeout")
	app.PersistentFlags().BoolVar(&c.Dump, "dump", false, "Dump HTTP request and response.")

	// Register signer flags
	var user, pass string
	app.PersistentFlags().StringVar(&user, "user", "", "Username used for authentication")
	app.PersistentFlags().StringVar(&pass, "pass", "", "Password used for authentication")

	// Parse flags and setup signers
	app.ParseFlags(os.Args)
	adminPassSigner := newAdminPassSigner(user, pass)

	// Initialize API client
	c.SetAdminPassSigner(adminPassSigner)
	c.UserAgent = "cellar-cli/0"

	// Register API commands
	cli.RegisterCommands(app, c)

	// Execute!
	if err := app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
}

// newHTTPClient returns the HTTP client used by the API client to make requests to the service.
func newHTTPClient() *http.Client {
	// TBD: Change as needed (e.g. to use a different transport to control redirection policy or
	// disable cert validation or...)
	return http.DefaultClient
}

// newAdminPassSigner returns the request signer used for authenticating
// against the admin_pass security scheme.
func newAdminPassSigner(user, pass string) goaclient.Signer {
	return &goaclient.BasicSigner{
		Username: user,
		Password: pass,
	}

}
