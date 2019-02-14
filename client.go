package cphalo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultTimeout      = 10 * time.Second
	DefaultMaxAuthTries = 3
	DefaultBaseURL      = "https://api.cloudpassage.com"
	DefaultAPIVersion   = "v1"
)

type Client struct {
	appKey       string
	appSecret    string
	accessToken  string
	baseURL      *url.URL
	timeout      time.Duration
	maxAuthTries int

	client *http.Client
}

func NewClient(appKey string, appSecret string) *Client {
	baseURL, _ := url.Parse(DefaultBaseURL)
	c := &Client{
		appKey:       appKey,
		appSecret:    appSecret,
		baseURL:      baseURL,
		timeout:      DefaultTimeout,
		maxAuthTries: DefaultMaxAuthTries,
	}
	c.client = &http.Client{Timeout: c.timeout}

	return c
}

func (c *Client) newRequest(method string, rsc string, params map[string]string, body interface{}) (*http.Request, error) {
	rawURL := c.baseURL.String() + "/" + DefaultAPIVersion + "/" + rsc
	baseURL, err := url.Parse(rawURL)

	if err != nil {
		return nil, fmt.Errorf("cannot parse url %s: %v", rawURL, err)
	}

	if params != nil {
		ps := url.Values{}
		for k, v := range params {
			ps.Set(k, v)
		}
		baseURL.RawQuery = ps.Encode()
	}

	var requestBody []byte
	if body != nil {
		requestBody, err = json.Marshal(body)

		if err != nil {
			return nil, fmt.Errorf("cannot marshall request body: %v", err)
		}
	}

	req, err := http.NewRequest(method, baseURL.String(), bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, fmt.Errorf("cannot create request: %v", err)
	}

	return req, nil
}
