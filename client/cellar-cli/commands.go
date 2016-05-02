package main

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar/client"
	goaclient "github.com/goadesign/goa/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
)

type (
	// CreateAccountCommand is the command line data structure for the create action of account
	CreateAccountCommand struct {
		Payload string
	}
	// DeleteAccountCommand is the command line data structure for the delete action of account
	DeleteAccountCommand struct {
		// Account ID
		AccountID int
	}
	// ShowAccountCommand is the command line data structure for the show action of account
	ShowAccountCommand struct {
		// Account ID
		AccountID int
	}
	// UpdateAccountCommand is the command line data structure for the update action of account
	UpdateAccountCommand struct {
		Payload string
		// Account ID
		AccountID int
	}
	// CreateBottleCommand is the command line data structure for the create action of bottle
	CreateBottleCommand struct {
		Payload string
		// Account ID
		AccountID int
	}
	// DeleteBottleCommand is the command line data structure for the delete action of bottle
	DeleteBottleCommand struct {
		// Account ID
		AccountID int
		BottleID  int
	}
	// ListBottleCommand is the command line data structure for the list action of bottle
	ListBottleCommand struct {
		// Account ID
		AccountID int
		// Filter by years
		Years []int
	}
	// RateBottleCommand is the command line data structure for the rate action of bottle
	RateBottleCommand struct {
		Payload string
		// Account ID
		AccountID int
		BottleID  int
	}
	// ShowBottleCommand is the command line data structure for the show action of bottle
	ShowBottleCommand struct {
		// Account ID
		AccountID int
		BottleID  int
	}
	// UpdateBottleCommand is the command line data structure for the update action of bottle
	UpdateBottleCommand struct {
		Payload string
		// Account ID
		AccountID int
		BottleID  int
	}
	// WatchBottleCommand is the command line data structure for the watch action of bottle
	WatchBottleCommand struct {
		// Account ID
		AccountID int
		BottleID  int
	}
)

// Run makes the HTTP request corresponding to the CreateAccountCommand command.
func (cmd *CreateAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/cellar/accounts"
	}
	var payload client.CreateAccountPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateAccount(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateAccountCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	c.SignerAdminPass.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the DeleteAccountCommand command.
func (cmd *DeleteAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v", cmd.AccountID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteAccount(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteAccountCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
}

// Run makes the HTTP request corresponding to the ShowAccountCommand command.
func (cmd *ShowAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v", cmd.AccountID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowAccount(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowAccountCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
}

// Run makes the HTTP request corresponding to the UpdateAccountCommand command.
func (cmd *UpdateAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v", cmd.AccountID)
	}
	var payload client.UpdateAccountPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateAccount(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateAccountCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
}

// Run makes the HTTP request corresponding to the CreateBottleCommand command.
func (cmd *CreateBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v/bottles", cmd.AccountID)
	}
	var payload client.CreateBottlePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateBottle(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
}

// Run makes the HTTP request corresponding to the DeleteBottleCommand command.
func (cmd *DeleteBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v/bottles/%v", cmd.AccountID, cmd.BottleID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteBottle(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
	var bottleID int
	cc.Flags().IntVar(&cmd.BottleID, "bottleID", bottleID, ``)
}

// Run makes the HTTP request corresponding to the ListBottleCommand command.
func (cmd *ListBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v/bottles", cmd.AccountID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListBottle(ctx, path, cmd.Years)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
	var years []int
	cc.Flags().IntSliceVar(&cmd.Years, "years", years, `Filter by years`)
}

// Run makes the HTTP request corresponding to the RateBottleCommand command.
func (cmd *RateBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v/bottles/%v/actions/rate", cmd.AccountID, cmd.BottleID)
	}
	var payload client.RateBottlePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RateBottle(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RateBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
	var bottleID int
	cc.Flags().IntVar(&cmd.BottleID, "bottleID", bottleID, ``)
}

// Run makes the HTTP request corresponding to the ShowBottleCommand command.
func (cmd *ShowBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v/bottles/%v", cmd.AccountID, cmd.BottleID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowBottle(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
	var bottleID int
	cc.Flags().IntVar(&cmd.BottleID, "bottleID", bottleID, ``)
}

// Run makes the HTTP request corresponding to the UpdateBottleCommand command.
func (cmd *UpdateBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v/bottles/%v", cmd.AccountID, cmd.BottleID)
	}
	var payload client.UpdateBottlePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateBottle(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
	var bottleID int
	cc.Flags().IntVar(&cmd.BottleID, "bottleID", bottleID, ``)
}

// Run establishes a websocket connection for the WatchBottleCommand command.
func (cmd *WatchBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v/bottles/%v/watch", cmd.AccountID, cmd.BottleID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	ws, err := c.WatchBottle(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}
	go goaclient.WSWrite(ws)
	goaclient.WSRead(ws)

	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *WatchBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
	var bottleID int
	cc.Flags().IntVar(&cmd.BottleID, "bottleID", bottleID, ``)
}
