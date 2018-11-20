package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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

type Client struct {
	AppKey       string
	AppSecret    string
	AccessToken  string
	BaseUrl      *url.URL
	Timeout      time.Duration
	MaxAuthTries int

	client *http.Client
}

func NewClient(appKey string, appSecret string) *Client {
	baseUrl, _ := url.Parse(DefaultBaseUrl)
	c := &Client{
		AppKey:       appKey,
		AppSecret:    appSecret,
		BaseUrl:      baseUrl,
		Timeout:      DefaultTimeout,
		MaxAuthTries: DefaultMaxAuthTries,
	}
	c.client = &http.Client{Timeout: c.Timeout}

	return c
}

func (c *Client) NewRequest(method string, rsc string, params map[string]string, body interface{}) (*http.Request, error) {
	rawURL := c.BaseUrl.String() + "/" + DefaultApiVersion + "/" + rsc
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

		log.Println("request body: ", string(requestBody))
	}

	req, err := http.NewRequest(method, baseUrl.String(), bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, fmt.Errorf("cannot create request: %v", err)
	}

	return req, nil
}
