package main

import (
	"fmt"
	"github.com/goadesign/goa-cellar/client"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// PrettyPrint is true if the tool output should be formatted for human consumption.
var PrettyPrint bool

func main() {
	// Create command line parser
	app := &cobra.Command{
		Use:   "cellar-cli",
		Short: `CLI client for the cellar service (http://goa.design/getting-started.html)`,
	}
	c := client.New(nil)
	c.UserAgent = "cellar-cli/0"
	app.PersistentFlags().StringVarP(&c.Scheme, "scheme", "s", "", "Set the requests scheme")
	app.PersistentFlags().StringVarP(&c.Host, "host", "H", "cellar.goa.design", "API hostname")
	app.PersistentFlags().DurationVarP(&c.Timeout, "timeout", "t", time.Duration(20)*time.Second, "Set the request timeout")
	app.PersistentFlags().BoolVar(&c.Dump, "dump", false, "Dump HTTP request and response.")
	app.PersistentFlags().BoolVar(&PrettyPrint, "pp", false, "Pretty print response body")
	RegisterCommands(app, c)
	if err := app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "request failed: %s", err)
		os.Exit(-1)
	}
}

// RegisterCommands all the resource action subcommands to the application command line.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `create action`,
	}
	tmp1 := new(CreateAccountCommand)
	sub = &cobra.Command{
		Use:   `account [/cellar/accounts] or`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp2 := new(CreateBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle [/cellar/accounts/ACCOUNTID/bottles] or`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `delete action`,
	}
	tmp3 := new(DeleteAccountCommand)
	sub = &cobra.Command{
		Use:   `account [/cellar/accounts/ACCOUNTID] or`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp4 := new(DeleteBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle [/cellar/accounts/ACCOUNTID/bottles/BOTTLEID] or`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `List all bottles in account optionally filtering by year`,
	}
	tmp5 := new(ListBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle [/cellar/accounts/ACCOUNTID/bottles] or`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "rate",
		Short: ``,
	}
	tmp6 := new(RateBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle [/cellar/accounts/ACCOUNTID/bottles/BOTTLEID/actions/rate] or`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp7 := new(ShowAccountCommand)
	sub = &cobra.Command{
		Use:   `account [/cellar/accounts/ACCOUNTID] or`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp8 := new(ShowBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle [/cellar/accounts/ACCOUNTID/bottles/BOTTLEID] or`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `update action`,
	}
	tmp9 := new(UpdateAccountCommand)
	sub = &cobra.Command{
		Use:   `account [/cellar/accounts/ACCOUNTID] or`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp10 := new(UpdateBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle [/cellar/accounts/ACCOUNTID/bottles/BOTTLEID] or`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "watch",
		Short: `Retrieve bottle with given id`,
	}
	tmp11 := new(WatchBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle [/cellar/accounts/ACCOUNTID/bottles/BOTTLEID/watch] or`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)

	dl := new(DownloadCommand)
	dlc := &cobra.Command{
		Use:   "download [PATH]",
		Short: "Download file with given path",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dl.Run(c, args)
		},
	}
	dlc.Flags().StringVar(&dl.OutFile, "out", "", "Output file")
	app.AddCommand(dlc)
}
