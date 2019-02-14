package api

import (
	"fmt"
	"net/http"
)

type FirewallPolicy struct {
	ID                    string         `json:"id,omitempty"`
	URL                   string         `json:"url,omitempty"`
	Name                  string         `json:"name,omitempty"`
	Platform              string         `json:"platform,omitempty"`
	Description           string         `json:"description,omitempty"`
	Shared                bool           `json:"shared,omitempty"`
	FirewallRules         []FirewallRule `json:"firewall_rules,omitempty"`
	IgnoreForwardingRules bool           `json:"ignore_forwarding_rules,omitempty"`
}

type ListFirewallPoliciesResponse struct {
	Count    int              `json:"count"`
	Policies []FirewallPolicy `json:"firewall_policies"`
}

type GetFirewallPolicyResponse struct {
	Policy FirewallPolicy `json:"firewall_policy"`
}

type CreateFirewallPolicyResponse = GetFirewallPolicyResponse
type CreateFirewallPolicyRequest = GetFirewallPolicyResponse
type UpdateFirewallPolicyRequest = GetFirewallPolicyResponse

func (c *client) ListFirewallPolicies() (response ListFirewallPoliciesResponse, err error) {
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

func (c *client) GetFirewallPolicy(ID string) (response GetFirewallPolicyResponse, err error) {
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

func (c *client) CreateFirewallPolicy(policy FirewallPolicy) (response CreateFirewallPolicyResponse, err error) {
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

func (c *client) UpdateFirewallPolicy(policy FirewallPolicy) error {
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

func (c *client) DeleteFirewallPolicy(ID string) error {
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
