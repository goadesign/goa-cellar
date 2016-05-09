package client

import (
	goaclient "github.com/goadesign/goa/client"
	"net/http"
)

// Client is the cellar service client.
type Client struct {
	*goaclient.Client
	AdminPassSigner *goaclient.BasicSigner
}

// New instantiates the client.
func New(c *http.Client) *Client {
	return &Client{
		Client:          goaclient.New(c),
		AdminPassSigner: &goaclient.BasicSigner{},
	}
}
