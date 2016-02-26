package main

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa-cellar/client"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// PrettyPrint is true if the tool output should be formatted for human consumption.
var PrettyPrint bool

func main() {
	// Create command line parser
	app := &cobra.Command{
		Use:   "cellar-cli",
		Short: "CLI client for the cellar service (http://goa.design/getting-started.html)",
	}
	c := client.New()
	c.UserAgent = "cellar-cli/1.0"
	app.PersistentFlags().StringVarP(&c.Scheme, "scheme", "s", "http", "Set the requests scheme")
	app.PersistentFlags().StringVarP(&c.Host, "host", "H", "cellar.goa.design", "API hostname")
	app.PersistentFlags().DurationVarP(&c.Timeout, "timeout", "t", time.Duration(20)*time.Second, "Set the request timeout, defaults to 20s")
	app.PersistentFlags().BoolVar(&c.Dump, "dump", false, "Dump HTTP request and response.")
	app.PersistentFlags().BoolVar(&PrettyPrint, "pp", false, "Pretty print response body")
	RegisterCommands(app, c)
	if err := app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "request failed: %s", err)
		os.Exit(-1)
	}
}

// HandleResponse unmarshals the response body and analyzes the status code to print then exit.
func HandleResponse(c *client.Client, resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read body: %s", err)
		os.Exit(-1)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// Let user know if something went wrong
		var sbody string
		if len(body) > 0 {
			sbody = ": " + string(body)
		}
		fmt.Printf("error: %d%s", resp.StatusCode, sbody)
	} else if !c.Dump && len(body) > 0 {
		var out string
		if PrettyPrint {
			var jbody interface{}
			err = json.Unmarshal(body, &jbody)
			if err != nil {
				out = string(body)
			} else {
				var b []byte
				b, err = json.MarshalIndent(jbody, "", "    ")
				if err == nil {
					out = string(b)
				} else {
					out = string(body)
				}
			}
		} else {
			out = string(body)
		}
		fmt.Print(out)
	}

	// Figure out exit code
	exitStatus := 0
	switch {
	case resp.StatusCode == 401:
		exitStatus = 1
	case resp.StatusCode == 403:
		exitStatus = 3
	case resp.StatusCode == 404:
		exitStatus = 4
	case resp.StatusCode > 399 && resp.StatusCode < 500:
		exitStatus = 2
	case resp.StatusCode > 499:
		exitStatus = 5
	}
	os.Exit(exitStatus)
}

// RegisterCommands all the resource action subcommands to the application command line.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: "create action",
	}
	tmp1 := new(CreateAccountCommand)
	sub = &cobra.Command{
		Use:   "account",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp2 := new(CreateBottleCommand)
	sub = &cobra.Command{
		Use:   "bottle",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp3 := new(CreateGenericBottleCommand)
	sub = &cobra.Command{
		Use:   "generic_bottle",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: "delete action",
	}
	tmp4 := new(DeleteAccountCommand)
	sub = &cobra.Command{
		Use:   "account",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp5 := new(DeleteBottleCommand)
	sub = &cobra.Command{
		Use:   "bottle",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp6 := new(DeleteGenericBottleCommand)
	sub = &cobra.Command{
		Use:   "generic_bottle",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: "list action",
	}
	tmp7 := new(ListBottleCommand)
	sub = &cobra.Command{
		Use:   "bottle",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp8 := new(ListGenericBottleCommand)
	sub = &cobra.Command{
		Use:   "generic_bottle",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "rate",
		Short: "rate action",
	}
	tmp9 := new(RateBottleCommand)
	sub = &cobra.Command{
		Use:   "bottle",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp10 := new(RateGenericBottleCommand)
	sub = &cobra.Command{
		Use:   "generic_bottle",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: "show action",
	}
	tmp11 := new(ShowAccountCommand)
	sub = &cobra.Command{
		Use:   "account",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp12 := new(ShowBottleCommand)
	sub = &cobra.Command{
		Use:   "bottle",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp13 := new(ShowGenericBottleCommand)
	sub = &cobra.Command{
		Use:   "generic_bottle",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp13.Run(c, args) },
	}
	tmp13.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: "update action",
	}
	tmp14 := new(UpdateAccountCommand)
	sub = &cobra.Command{
		Use:   "account",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp14.Run(c, args) },
	}
	tmp14.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp15 := new(UpdateBottleCommand)
	sub = &cobra.Command{
		Use:   "bottle",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp15.Run(c, args) },
	}
	tmp15.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp16 := new(UpdateGenericBottleCommand)
	sub = &cobra.Command{
		Use:   "generic_bottle",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp16.Run(c, args) },
	}
	tmp16.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)

}
