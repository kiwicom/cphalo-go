package cphalo

import (
	"fmt"
	"net/http"
)

// FirewallPolicy represent a CPHalo firewall policy.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#object-representation-6
type FirewallPolicy struct {
	ID                    string         `json:"id,omitempty"`
	URL                   string         `json:"url,omitempty"`
	Name                  string         `json:"name,omitempty"`
	Platform              string         `json:"platform,omitempty"`
	Description           string         `json:"description,omitempty"`
	Shared                StringableBool `json:"shared"`
	FirewallRules         []FirewallRule `json:"firewall_rules,omitempty"`
	IgnoreForwardingRules bool           `json:"ignore_forwarding_rules,omitempty"`
}

// ListFirewallPoliciesResponse represent a list of firewall policies response.
type ListFirewallPoliciesResponse struct {
	Count    int              `json:"count"`
	Policies []FirewallPolicy `json:"firewall_policies"`
}

// GetFirewallPolicyResponse represent a get firewall policy response.
type GetFirewallPolicyResponse struct {
	Policy FirewallPolicy `json:"firewall_policy"`
}

// CreateFirewallPolicyResponse represent a create firewall policy response.
type CreateFirewallPolicyResponse = GetFirewallPolicyResponse

// CreateFirewallPolicyRequest represent a create firewall policy request.
type CreateFirewallPolicyRequest = GetFirewallPolicyResponse

// UpdateFirewallPolicyRequest represent a update firewall policy request.
type UpdateFirewallPolicyRequest = GetFirewallPolicyResponse

// ListFirewallPolicies lists all firewall policies.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#list-firewall-policies
func (c *Client) ListFirewallPolicies() (response ListFirewallPoliciesResponse, err error) {
	req, err := c.newRequest(http.MethodGet, "firewall_policies", nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute request: %v", err)
	}

	return response, nil
}

// GetFirewallPolicy returns details of the firewall policy.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#get-firewall-policy-details-including-firewall-rules
func (c *Client) GetFirewallPolicy(ID string) (response GetFirewallPolicyResponse, err error) {
	req, err := c.newRequest(http.MethodGet, "firewall_policies/"+ID, nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// CreateFirewallPolicy creates a new firewall policy.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#create-new-firewall-policy
func (c *Client) CreateFirewallPolicy(policy FirewallPolicy) (response CreateFirewallPolicyResponse, err error) {
	req, err := c.newRequest(http.MethodPost, "firewall_policies", nil, CreateFirewallPolicyRequest{Policy: policy})
	if err != nil {
		return response, fmt.Errorf("cannot create new create request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute create request: %v", err)
	}

	return response, nil
}

// UpdateFirewallPolicy updates firewall policy.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#update-name-or-description-for-the-firewall-policy
func (c *Client) UpdateFirewallPolicy(policy FirewallPolicy) error {
	req, err := c.newRequest(http.MethodPut, "firewall_policies/"+policy.ID, nil, UpdateFirewallPolicyRequest{Policy: policy})
	if err != nil {
		return fmt.Errorf("cannot create new update request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute update request: %v", err)
	}

	return nil
}

// DeleteFirewallPolicy deletes a firewall policy.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#delete-firewall-policy
func (c *Client) DeleteFirewallPolicy(ID string) error {
	req, err := c.newRequest(http.MethodDelete, "firewall_policies/"+ID, nil, nil)
	if err != nil {
		return fmt.Errorf("cannot create new delete request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute delete request: %v", err)
	}

	return nil
}
