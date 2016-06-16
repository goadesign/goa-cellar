package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateAccountPayload is the account create action payload.
type CreateAccountPayload struct {
	// Name of account
	Name string `json:"name" xml:"name" form:"name"`
}

// CreateAccountPath computes a request path to the create action of account.
func CreateAccountPath() string {
	return fmt.Sprintf("/cellar/accounts")
}

// Create new account
func (c *Client) CreateAccount(ctx context.Context, path string, payload *CreateAccountPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateAccountRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateAccountRequest create the request corresponding to the create action endpoint of the account resource.
func (c *Client) NewCreateAccountRequest(ctx context.Context, path string, payload *CreateAccountPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteAccountPath computes a request path to the delete action of account.
func DeleteAccountPath(accountID int) string {
	return fmt.Sprintf("/cellar/accounts/%v", accountID)
}

// DeleteAccount makes a request to the delete action endpoint of the account resource
func (c *Client) DeleteAccount(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteAccountRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteAccountRequest create the request corresponding to the delete action endpoint of the account resource.
func (c *Client) NewDeleteAccountRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowAccountPath computes a request path to the show action of account.
func ShowAccountPath(accountID int) string {
	return fmt.Sprintf("/cellar/accounts/%v", accountID)
}

// Retrieve account with given id. IDs 1 and 2 pre-exist in the system.
func (c *Client) ShowAccount(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowAccountRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowAccountRequest create the request corresponding to the show action endpoint of the account resource.
func (c *Client) NewShowAccountRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateAccountPayload is the account update action payload.
type UpdateAccountPayload struct {
	// Name of account
	Name string `json:"name" xml:"name" form:"name"`
}

// UpdateAccountPath computes a request path to the update action of account.
func UpdateAccountPath(accountID int) string {
	return fmt.Sprintf("/cellar/accounts/%v", accountID)
}

// Change account name
func (c *Client) UpdateAccount(ctx context.Context, path string, payload *UpdateAccountPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateAccountRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateAccountRequest create the request corresponding to the update action endpoint of the account resource.
func (c *Client) NewUpdateAccountRequest(ctx context.Context, path string, payload *UpdateAccountPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
