package cphalo

import (
	"fmt"
	"net/http"
)

// ServerGroupFirewallPolicy represents firewall policies for a CPHalo server group.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#object-representation-1
type ServerGroupFirewallPolicy struct {
	GroupID               string
	LinuxFirewallPolicyID NullableString `json:"linux_firewall_policy_id"`
}

// GetServerGroupFirewallPolicyResponse represents a CPHalo server group firewall policies get response.
type GetServerGroupFirewallPolicyResponse struct {
	Group ServerGroupFirewallPolicy `json:"group"`
}

// UpdateServerGroupFirewallPolicyRequest represents a CPHalo server group firewall policies update request.
type UpdateServerGroupFirewallPolicyRequest = GetServerGroupFirewallPolicyResponse

// GetServerGroupFirewallPolicy return information describing firewall policies for a single group.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#get-a-single-server-group
func (c *Client) GetServerGroupFirewallPolicy(ID string) (response GetServerGroupFirewallPolicyResponse, err error) {
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

// UpdateServerGroupFirewallPolicy updates firewall policies for a server group.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#assign-a-firewall-policy-to-the-server-group
func (c *Client) UpdateServerGroupFirewallPolicy(group ServerGroupFirewallPolicy) error {
	req, err := c.newRequest(http.MethodPut, "groups/"+group.GroupID, nil, UpdateServerGroupFirewallPolicyRequest{Group: group})
	if err != nil {
		return fmt.Errorf("cannot create new update request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute update request: %v", err)
	}

	return nil
}
