package client

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"net/http"
)

// Client is the cellar service client.
type Client struct {
	*goaclient.Client
	AdminPassSigner goaclient.Signer
	Encoder         *goa.HTTPEncoder
	Decoder         *goa.HTTPDecoder
}

// New instantiates the client.
func New(c *http.Client) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	client.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")
	client.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	client.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}

// SetAdminPassSigner sets the request signer for the admin_pass security scheme.
func (c *Client) SetAdminPassSigner(signer goaclient.Signer) {
	c.AdminPassSigner = signer
}
