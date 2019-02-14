package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c *Client) renewAccessToken() error {
	rsc := "/oauth/access_token?grant_type=client_credentials"
	rawURL := c.baseURL.String() + rsc
	baseURL, err := url.Parse(rawURL)

	if err != nil {
		return fmt.Errorf("cannot parse url %s: %v", rawURL, err)
	}

	authString := c.appKey + ":" + c.appSecret
	encodedAuthString := base64.StdEncoding.EncodeToString([]byte(authString))

	req, err := http.NewRequest(http.MethodPost, baseURL.String(), nil)

	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Authorization", "Basic "+encodedAuthString)

	resp, err := c.client.Do(req)

	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("cannot read body: %v", err)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("invalid credentials")
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with code %d", resp.StatusCode)
	}

	m := &accessTokenResponse{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return fmt.Errorf("unmarshalling failed: %v", err)
	}

	c.accessToken = m.AccessToken

	return nil
}
