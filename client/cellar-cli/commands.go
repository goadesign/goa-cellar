package main

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa-cellar/client"
	"github.com/spf13/cobra"
)

type (
	// CreateAccountCommand is the command line data structure for the create action of account
	CreateAccountCommand struct {
		Payload string
	}

	// DeleteAccountCommand is the command line data structure for the delete action of account
	DeleteAccountCommand struct {
	}

	// ShowAccountCommand is the command line data structure for the show action of account
	ShowAccountCommand struct {
	}

	// UpdateAccountCommand is the command line data structure for the update action of account
	UpdateAccountCommand struct {
		Payload string
	}

	// CreateBottleCommand is the command line data structure for the create action of bottle
	CreateBottleCommand struct {
		Payload string
	}

	// DeleteBottleCommand is the command line data structure for the delete action of bottle
	DeleteBottleCommand struct {
	}

	// ListBottleCommand is the command line data structure for the list action of bottle
	ListBottleCommand struct {
		// Filter by years
		Years []int
	}

	// RateBottleCommand is the command line data structure for the rate action of bottle
	RateBottleCommand struct {
		Payload string
	}

	// ShowBottleCommand is the command line data structure for the show action of bottle
	ShowBottleCommand struct {
	}

	// UpdateBottleCommand is the command line data structure for the update action of bottle
	UpdateBottleCommand struct {
		Payload string
	}

	// CreateGenericBottleCommand is the command line data structure for the create action of generic_bottle
	CreateGenericBottleCommand struct {
		Payload string
	}

	// DeleteGenericBottleCommand is the command line data structure for the delete action of generic_bottle
	DeleteGenericBottleCommand struct {
	}

	// ListGenericBottleCommand is the command line data structure for the list action of generic_bottle
	ListGenericBottleCommand struct {
		// Filter by years
		Years []int
	}

	// RateGenericBottleCommand is the command line data structure for the rate action of generic_bottle
	RateGenericBottleCommand struct {
		Payload string
	}

	// ShowGenericBottleCommand is the command line data structure for the show action of generic_bottle
	ShowGenericBottleCommand struct {
	}

	// UpdateGenericBottleCommand is the command line data structure for the update action of generic_bottle
	UpdateGenericBottleCommand struct {
		Payload string
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
	resp, err := c.CreateAccount(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateAccountCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the DeleteAccountCommand command.
func (cmd *DeleteAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.DeleteAccount(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteAccountCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the ShowAccountCommand command.
func (cmd *ShowAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.ShowAccount(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowAccountCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the UpdateAccountCommand command.
func (cmd *UpdateAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload client.UpdateAccountPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.UpdateAccount(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateAccountCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the CreateBottleCommand command.
func (cmd *CreateBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload client.CreateBottlePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.CreateBottle(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateBottleCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the DeleteBottleCommand command.
func (cmd *DeleteBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.DeleteBottle(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteBottleCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the ListBottleCommand command.
func (cmd *ListBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.ListBottle(path, cmd.Years)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListBottleCommand) RegisterFlags(cc *cobra.Command) {
	var tmp17 []int
	cc.Flags().IntSliceVar(&cmd.Years, "years", tmp17, "Filter by years")
}

// Run makes the HTTP request corresponding to the RateBottleCommand command.
func (cmd *RateBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload client.RateBottlePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.RateBottle(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RateBottleCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the ShowBottleCommand command.
func (cmd *ShowBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.ShowBottle(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowBottleCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the UpdateBottleCommand command.
func (cmd *UpdateBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload client.UpdateBottlePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.UpdateBottle(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateBottleCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the CreateGenericBottleCommand command.
func (cmd *CreateGenericBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload client.CreateGenericBottlePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.CreateGenericBottle(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateGenericBottleCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the DeleteGenericBottleCommand command.
func (cmd *DeleteGenericBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.DeleteGenericBottle(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteGenericBottleCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the ListGenericBottleCommand command.
func (cmd *ListGenericBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.ListGenericBottle(path, cmd.Years)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListGenericBottleCommand) RegisterFlags(cc *cobra.Command) {
	var tmp18 []int
	cc.Flags().IntSliceVar(&cmd.Years, "years", tmp18, "Filter by years")
}

// Run makes the HTTP request corresponding to the RateGenericBottleCommand command.
func (cmd *RateGenericBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload client.RateGenericBottlePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.RateGenericBottle(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RateGenericBottleCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the ShowGenericBottleCommand command.
func (cmd *ShowGenericBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.ShowGenericBottle(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowGenericBottleCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the UpdateGenericBottleCommand command.
func (cmd *UpdateGenericBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload client.UpdateGenericBottlePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.UpdateGenericBottle(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateGenericBottleCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}
