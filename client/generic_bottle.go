package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// CreateGenericBottlePayload is the data structure used to initialize the generic_bottle create request body.
type CreateGenericBottlePayload struct {
	Color   string  `json:"color" xml:"color"`
	Country *string `json:"country,omitempty" xml:"country,omitempty"`
	// Bottle kind
	Kind      string  `json:"kind" xml:"kind"`
	Name      string  `json:"name" xml:"name"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	Varietal  string  `json:"varietal" xml:"varietal"`
	Vineyard  string  `json:"vineyard" xml:"vineyard"`
	Vintage   int     `json:"vintage" xml:"vintage"`
}

// Record new bottle
func (c *Client) CreateGenericBottle(path string, payload *CreateGenericBottlePayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// DeleteGenericBottle makes a request to the delete action endpoint of the generic_bottle resource
func (c *Client) DeleteGenericBottle(path string) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// List all bottles in account optionally filtering by year
func (c *Client) ListGenericBottle(path string, years []int) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	values := u.Query()
	tmp23 := make([]string, len(years))
	for i, e := range years {
		tmp24 := strconv.Itoa(e)
		tmp23[i] = tmp24
	}
	tmp22 := strings.Join(tmp23, ",")
	values.Set("years", tmp22)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// RateGenericBottlePayload is the data structure used to initialize the generic_bottle rate request body.
type RateGenericBottlePayload struct {
	// Rating of bottle between 1 and 5
	Rating int `json:"rating" xml:"rating"`
}

// RateGenericBottle makes a request to the rate action endpoint of the generic_bottle resource
func (c *Client) RateGenericBottle(path string, payload *RateGenericBottlePayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// Retrieve bottle with given id
func (c *Client) ShowGenericBottle(path string) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// UpdateGenericBottlePayload is the data structure used to initialize the generic_bottle update request body.
type UpdateGenericBottlePayload struct {
	Color   *string `json:"color,omitempty" xml:"color,omitempty"`
	Country *string `json:"country,omitempty" xml:"country,omitempty"`
	// Bottle kind
	Kind      *string `json:"kind,omitempty" xml:"kind,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	Varietal  *string `json:"varietal,omitempty" xml:"varietal,omitempty"`
	Vineyard  *string `json:"vineyard,omitempty" xml:"vineyard,omitempty"`
	Vintage   *int    `json:"vintage,omitempty" xml:"vintage,omitempty"`
}

// UpdateGenericBottle makes a request to the update action endpoint of the generic_bottle resource
func (c *Client) UpdateGenericBottle(path string, payload *UpdateGenericBottlePayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}
