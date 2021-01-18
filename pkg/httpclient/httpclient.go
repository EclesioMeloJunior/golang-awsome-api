package httpclient

import (
	"net/http"
	"time"
)

// HTTPClient interface will define an abstraction
// to HTTP calls
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type client struct {
	*http.Client
}

// NewHTTPClient returns the implementation of HTTPClient
// using the golang http library
func NewHTTPClient() HTTPClient {
	return &client{
		Client: &http.Client{
			Timeout: 240 * time.Second,
		},
	}
}

func (c *client) Do(r *http.Request) (*http.Response, error) {
	return c.Client.Do(r)
}
