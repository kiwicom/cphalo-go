package api

import (
	"fmt"
	"net/http"
)

type ServerGroup struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
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
	req, err := c.NewRequest(http.MethodGet, "groups", nil)
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
	req, err := c.NewRequest(http.MethodGet, "groups/"+ID, nil)
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
	_ = CreateServerGroupRequest{Group: group}
	return response, nil
}

func (c *Client) UpdateServerGroup(group ServerGroup) error {
	_ = GetServerGroupResponse{Group: group}
	return nil
}

func (c *Client) DeleteServerGroup(ID string) error {
	return nil
}
