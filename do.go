package cphalo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	return c.doTries(req, v, 0)
}

func (c *Client) doTries(req *http.Request, v interface{}, tries int) (*http.Response, error) {
	if tries >= c.maxAuthTries {
		return nil, fmt.Errorf("max tries exceeded")
	}

	if len(c.accessToken) == 0 {
		tries = tries + 1
		if err := c.renewAccessToken(); err != nil {
			return nil, fmt.Errorf("cannot set access token: %v", err)
		}
	}

	req.Header.Set("Authorization", "Bearer "+c.accessToken)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %v", err)
	}
	defer resp.Body.Close()

	// check if access token issue expired
	// https://library.cloudpassage.com/help/cloudpassage-api-documentation#token-management
	// the docs say 402, but in reality only 401 is used
	if resp.StatusCode == http.StatusPaymentRequired || resp.StatusCode == http.StatusUnauthorized {
		if err := c.renewAccessToken(); err != nil {
			return nil, fmt.Errorf("cannot renew access token: %v", err)
		}

		return c.doTries(req, v, tries+1)
	}

	if err := validateResponse(resp); err != nil {
		return nil, err
	}

	// no need to unmarshal body
	if v == nil {
		return resp, nil
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

	var customErr CPHaloResponseError

	switch r.StatusCode / 100 {
	case 4:
		switch r.StatusCode {
		case http.StatusNotFound:
			customErr = &ResponseError404{}
		case http.StatusUnprocessableEntity:
			customErr = &ResponseError422{}
		case http.StatusTooManyRequests:
			customErr = &ResponseError429{}
		default:
			customErr = &ResponseError400{StatusCode: r.StatusCode}
		}
	case 5:
		customErr = &ResponseError500{}
	default:
		customErr = &ResponseErrorGeneral{}
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return fmt.Errorf("cannot read error body: %v", err)
	}

	if len(bodyBytes) == 0 {
		return customErr
	}

	if err := json.Unmarshal(bodyBytes, &customErr); err != nil {
		return fmt.Errorf("cannot unmarshall error response: %v", err)
	}

	return customErr
}
