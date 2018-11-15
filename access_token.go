package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c *Client) RenewAccessToken() error {
	rsc := "/oauth/access_token?grant_type=client_credentials"
	rawURL := c.BaseUrl.String() + rsc
	baseUrl, err := url.Parse(rawURL)
	log.Println("Going to authenticate and obtain access token.")
	log.Println("Auth URL is " + baseUrl.String())

	if err != nil {
		return fmt.Errorf("cannot parse url %s: %v", rawURL, err)
	}

	authString := c.AppKey + ":" + c.AppSecret
	encodedAuthString := base64.StdEncoding.EncodeToString([]byte(authString))

	req, err := http.NewRequest(http.MethodPost, baseUrl.String(), nil)
	req.Header.Add("Authorization", "Basic "+encodedAuthString)

	resp, err := c.client.Do(req)

	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if err != nil {
		return fmt.Errorf("cannot read body: %v", err)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("invalid credentials")
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with code %d: %s", resp.StatusCode, bodyString)
	}

	m := &accessTokenResponse{}
	err = json.Unmarshal([]byte(bodyString), &m)
	if err != nil {
		return fmt.Errorf("unmarshalling failed: %v", err)
	}

	c.AccessToken = m.AccessToken
	log.Printf("access token: %+v\n", m)

	return nil
}
