package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa-cellar/app"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Record new bottle
func (c *Client) CreateBottle(path string, payload *app.CreateBottlePayload) (*http.Response, error) {
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

// DeleteBottle makes a request to the delete action endpoint of the bottle resource
func (c *Client) DeleteBottle(path string) (*http.Response, error) {
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
func (c *Client) ListBottle(path string, years []int) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	values := u.Query()
	tmp13 := make([]string, len(years))
	for i, e := range years {
		tmp14 := strconv.Itoa(e)
		tmp13[i] = tmp14
	}
	tmp12 := strings.Join(tmp13, ",")
	values.Set("years", tmp12)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// RateBottle makes a request to the rate action endpoint of the bottle resource
func (c *Client) RateBottle(path string, payload *app.RateBottlePayload) (*http.Response, error) {
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
func (c *Client) ShowBottle(path string) (*http.Response, error) {
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

// UpdateBottle makes a request to the update action endpoint of the bottle resource
func (c *Client) UpdateBottle(path string, payload *app.UpdateBottlePayload) (*http.Response, error) {
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
