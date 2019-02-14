package api

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
	DefaultBaseUrl      = "https://api.cloudpassage.com"
	DefaultApiVersion   = "v1"
)

type client struct {
	appKey       string
	appSecret    string
	accessToken  string
	baseUrl      *url.URL
	timeout      time.Duration
	maxAuthTries int

	client *http.Client
}

func NewClient(appKey string, appSecret string) *client {
	baseUrl, _ := url.Parse(DefaultBaseUrl)
	c := &client{
		appKey:       appKey,
		appSecret:    appSecret,
		baseUrl:      baseUrl,
		timeout:      DefaultTimeout,
		maxAuthTries: DefaultMaxAuthTries,
	}
	c.client = &http.Client{Timeout: c.timeout}

	return c
}

func (c *client) newRequest(method string, rsc string, params map[string]string, body interface{}) (*http.Request, error) {
	rawURL := c.baseUrl.String() + "/" + DefaultApiVersion + "/" + rsc
	baseUrl, err := url.Parse(rawURL)

	if err != nil {
		return nil, fmt.Errorf("cannot parse url %s: %v", rawURL, err)
	}

	if params != nil {
		ps := url.Values{}
		for k, v := range params {
			ps.Set(k, v)
		}
		baseUrl.RawQuery = ps.Encode()
	}

	var requestBody []byte
	if body != nil {
		requestBody, err = json.Marshal(body)

		if err != nil {
			return nil, fmt.Errorf("cannot marshall request body: %v", err)
		}
	}

	req, err := http.NewRequest(method, baseUrl.String(), bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, fmt.Errorf("cannot create request: %v", err)
	}

	return req, nil
}
