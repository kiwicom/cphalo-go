package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	return c.doTries(req, v, 0)
}

func (c *Client) doTries(req *http.Request, v interface{}, tries int) (*http.Response, error) {
	log.Println("making request to: " + req.URL.String())
	if tries >= c.MaxAuthTries {
		return nil, fmt.Errorf("max tries exceeded")
	}

	if len(c.AccessToken) == 0 {
		tries = tries + 1
		log.Println("access token not set")
		if err := c.RenewAccessToken(); err != nil {
			return nil, fmt.Errorf("cannot set access token: %v", err)
		}
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %v", err)
	}
	defer resp.Body.Close()

	// check if access token issue expired
	// https://library.cloudpassage.com/help/cloudpassage-api-documentation#token-management
	// the docs say 402, but in reality only 401 is used
	if resp.StatusCode == http.StatusPaymentRequired || resp.StatusCode == http.StatusUnauthorized {
		log.Println("access token expired")
		if err := c.RenewAccessToken(); err != nil {
			return nil, fmt.Errorf("cannot renew access token: %v", err)
		}

		return c.doTries(req, v, tries+1)
	}

	if err := validateResponse(resp); err != nil {
		return nil, fmt.Errorf("response validation failed: %v", err)
	}

	err = parseResponse(resp, v)

	if err != nil {
		return nil, fmt.Errorf("cannon parse response: %v", err)
	}

	return resp, err
}

func parseResponse(r *http.Response, v interface{}) error {
	if v == nil {
		return fmt.Errorf("nil interface provided")
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return fmt.Errorf("cannot read response body: %v", err)
	}

	bodyString := string(bodyBytes)
	//fmt.Println(bodyString)
	err = json.Unmarshal([]byte(bodyString), &v)

	if err != nil {
		return fmt.Errorf("cannot unmarshall body: %v", err)
	}

	return nil
}

func validateResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	log.Println("processing error response")
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return fmt.Errorf("cannot read error body: %v", err)
	}

	bodyString := string(bodyBytes)
	log.Println(bodyString)

	m := &ResponseError{}

	if err := json.Unmarshal([]byte(bodyString), &m); err != nil {
		return fmt.Errorf("cannot unmarshall error response: %v", err)
	}

	return fmt.Errorf("error response: %s", m.Error)
}
