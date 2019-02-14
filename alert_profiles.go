package cphalo

import (
	"fmt"
	"net/http"
	"time"
)

type AlertProfile struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	GroupID     string    `json:"group_id"`
	GroupName   string    `json:"group_name"`
	Description string    `json:"description"`
	Frequency   string    `json:"frequency"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	Shared      bool      `json:"shared"`
	UsedBy      []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"used_by"`
}

type ListAlertProfilesResponse struct {
	Count         int            `json:"count"`
	AlertProfiles []AlertProfile `json:"alert_profiles"`
}

func (c *Client) ListAlertProfiles() (response ListAlertProfilesResponse, err error) {
	req, err := c.newRequest(http.MethodGet, "alert_profiles", nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
