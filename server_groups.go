package api

import (
	"fmt"
	"net/http"
)

type ServerGroup struct {
	ID          string `json:"id,omitempty"`
	URL         string `json:"url,omitempty"`
	Name        string `json:"name"`
	ParentID    string `json:"parent_id,omitempty"`
	HasChildren bool   `json:"has_children,omitempty"`
}

type ListServerGroupsResponse struct {
	Count  int           `json:"count"`
	Groups []ServerGroup `json:"groups"`
}

type GetServerGroupResponse struct {
	Group ServerGroup `json:"group"`
}

type CreateServerGroupResponse = GetServerGroupResponse
type CreateServerGroupRequest = GetServerGroupResponse

func (c *Client) ListServerGroups() (response ListServerGroupsResponse, err error) {
	req, err := c.NewRequest(http.MethodGet, "groups", nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute request: %v", err)
	}

	return response, nil
}

func (c *Client) GetServerGroup(ID string) (response GetServerGroupResponse, err error) {
	req, err := c.NewRequest(http.MethodGet, "groups/"+ID, nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute request: %v", err)
	}

	return response, nil
}

func (c *Client) CreateServerGroup(group ServerGroup) (response CreateServerGroupResponse, err error) {
	req, err := c.NewRequest(http.MethodPost, "groups", nil, CreateServerGroupRequest{Group: group})
	if err != nil {
		return response, fmt.Errorf("cannot create new create request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute create request: %v", err)
	}

	return response, nil
}

func (c *Client) UpdateServerGroup(group ServerGroup) error {
	_ = GetServerGroupResponse{Group: group}
	return nil
}

func (c *Client) DeleteServerGroup(ID string) error {
	req, err := c.NewRequest(http.MethodDelete, "groups/"+ID, nil, nil)
	if err != nil {
		return fmt.Errorf("cannot create new delete request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute delete request: %v", err)
	}

	return nil
}
