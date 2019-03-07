package cphalo

import (
	"fmt"
	"net/http"
)

// ServerGroup represent a CPHalo server group.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#object-representation-1
type ServerGroup struct {
	ID                    string         `json:"id,omitempty"`
	URL                   string         `json:"url,omitempty"`
	Name                  string         `json:"name,omitempty"`
	Description           string         `json:"description,omitempty"`
	ParentID              string         `json:"parent_id,omitempty"`
	HasChildren           bool           `json:"has_children,omitempty"`
	Tag                   string         `json:"tag,omitempty"`
	LinuxFirewallPolicyID NullableString `json:"linux_firewall_policy_id"`
	AlertProfileIDs       []string       `json:"alert_profile_ids,omitempty"`
}

// ListServerGroupsResponse represent a CPHalo server group list response.
type ListServerGroupsResponse struct {
	Count  int           `json:"count"`
	Groups []ServerGroup `json:"groups"`
}

// GetServerGroupResponse represent a CPHalo server group get response.
type GetServerGroupResponse struct {
	Group ServerGroup `json:"group"`
}

// CreateServerGroupResponse represent a CPHalo server group create response.
type CreateServerGroupResponse = GetServerGroupResponse

// CreateServerGroupRequest represent a CPHalo server group create request.
type CreateServerGroupRequest = GetServerGroupResponse

// UpdateServerGroupRequest represent a CPHalo server group update request.
type UpdateServerGroupRequest = GetServerGroupResponse

// ListServerGroups lists all server groups.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#list-server-groups
func (c *Client) ListServerGroups() (response ListServerGroupsResponse, err error) {
	req, err := c.newRequest(http.MethodGet, "groups", nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute request: %v", err)
	}

	return response, nil
}

// GetServerGroup return information describing a single group.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#get-a-single-server-group
func (c *Client) GetServerGroup(ID string) (response GetServerGroupResponse, err error) {
	req, err := c.newRequest(http.MethodGet, "groups/"+ID, nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// CreateServerGroup creates new server group.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#create-a-new-server-group
func (c *Client) CreateServerGroup(group ServerGroup) (response CreateServerGroupResponse, err error) {
	req, err := c.newRequest(http.MethodPost, "groups", nil, CreateServerGroupRequest{Group: group})
	if err != nil {
		return response, fmt.Errorf("cannot create new create request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute create request: %v", err)
	}

	return response, nil
}

// UpdateServerGroup updates server group.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#update-server-group-attributes
func (c *Client) UpdateServerGroup(group ServerGroup) error {
	gID := group.ID
	group.ID = ""

	req, err := c.newRequest(http.MethodPut, "groups/"+gID, nil, UpdateServerGroupRequest{Group: group})
	if err != nil {
		return fmt.Errorf("cannot create new update request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute update request: %v", err)
	}

	return nil
}

// DeleteServerGroup deletes server group.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#delete-server-group-without-any-servers
func (c *Client) DeleteServerGroup(ID string) error {
	req, err := c.newRequest(http.MethodDelete, "groups/"+ID, nil, nil)
	if err != nil {
		return fmt.Errorf("cannot create new delete request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute delete request: %v", err)
	}

	return nil
}
