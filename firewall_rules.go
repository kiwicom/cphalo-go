package api

import (
	"fmt"
	"net/http"
)

type FirewallRuleSourceTarget struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	IpAddress string `json:"ip_address,omitempty"`
	Kind      string `json:"type,omitempty"`
}

func (f *FirewallRuleSourceTarget) GetID() string {
	if f.ID == "" {
		return f.Name
	}

	return f.ID
}

type FirewallRule struct {
	ID                string                    `json:"id,omitempty"`
	URL               string                    `json:"url,omitempty"`
	Chain             string                    `json:"chain,omitempty"`
	Action            string                    `json:"action,omitempty"`
	Active            bool                      `json:"active,omitempty"`
	ConnectionStates  string                    `json:"connection_states,omitempty"`
	Position          int                       `json:"position,omitempty"`
	FirewallInterface *FirewallInterface        `json:"firewall_interface,omitempty"`
	FirewallService   *FirewallService          `json:"firewall_service,omitempty"`
	FirewallSource    *FirewallRuleSourceTarget `json:"firewall_source,omitempty"`
	FirewallTarget    *FirewallRuleSourceTarget `json:"firewall_target,omitempty"`
}

type ListFirewallRulesResponse struct {
	Count int            `json:"count"`
	Rules []FirewallRule `json:"firewall_rules"`
}

type GetFirewallRuleResponse struct {
	Rule FirewallRule `json:"firewall_rule"`
}

type CreateFirewallRuleResponse = GetFirewallRuleResponse
type CreateFirewallRuleRequest = GetFirewallRuleResponse
type UpdateFirewallRuleRequest = GetFirewallRuleResponse

func (r *FirewallRule) applyCorrections() {
	const (
		allServers = "All Active Servers"
		allUsers   = "All GhostPorts users"
	)

	if r.FirewallSource != nil {
		switch r.FirewallSource.Kind {
		case "Group":
			if r.FirewallSource.ID == allServers {
				r.FirewallSource.Name = r.FirewallSource.ID
				r.FirewallSource.ID = ""
			}
		case "UserGroup":
			if r.FirewallSource.ID == allUsers {
				r.FirewallSource.Name = r.FirewallSource.ID
				r.FirewallSource.ID = ""
			}
		}
	}

	if r.FirewallTarget != nil {
		switch r.FirewallTarget.Kind {
		case "Group":
			if r.FirewallTarget.ID == allServers {
				r.FirewallTarget.Name = r.FirewallTarget.ID
				r.FirewallTarget.ID = ""
			}
		case "UserGroup":
			if r.FirewallTarget.ID == allUsers {
				r.FirewallTarget.Name = r.FirewallTarget.ID
				r.FirewallTarget.ID = ""
			}
		}
	}

	if r.ConnectionStates == "ANY" {
		r.ConnectionStates = ""
	}
}

func (c *Client) ListFirewallRules(policyID string) (response ListFirewallRulesResponse, err error) {
	url := fmt.Sprintf("firewall_policies/%s/firewall_rules", policyID)
	req, err := c.NewRequest(http.MethodGet, url, nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute request: %v", err)
	}

	for _, rule := range response.Rules {
		if rule.ConnectionStates == "" {
			rule.ConnectionStates = "ANY"
		}
	}

	return response, nil
}

func (c *Client) GetFirewallRule(policyID, ruleID string) (response GetFirewallRuleResponse, err error) {
	url := fmt.Sprintf("firewall_policies/%s/firewall_rules/%s", policyID, ruleID)
	req, err := c.NewRequest(http.MethodGet, url, nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, err
	}

	if response.Rule.ConnectionStates == "" {
		response.Rule.ConnectionStates = "ANY"
	}

	return response, nil
}

func (c *Client) CreateFirewallRule(policyID string, rule FirewallRule) (response CreateFirewallRuleResponse, err error) {
	url := fmt.Sprintf("firewall_policies/%s/firewall_rules", policyID)
	rule.applyCorrections()
	req, err := c.NewRequest(http.MethodPost, url, nil, CreateFirewallRuleRequest{Rule: rule})
	if err != nil {
		return response, fmt.Errorf("cannot create new create request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute create request: %v", err)
	}

	return response, nil
}

func (c *Client) UpdateFirewallRule(policyID string, rule FirewallRule) error {
	url := fmt.Sprintf("firewall_policies/%s/firewall_rules/%s", policyID, rule.ID)
	rule.applyCorrections()
	req, err := c.NewRequest(http.MethodPut, url, nil, UpdateFirewallRuleRequest{Rule: rule})
	if err != nil {
		return fmt.Errorf("cannot create new update request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute update request: %v", err)
	}

	return nil
}

func (c *Client) DeleteFirewallRule(policyID, ruleID string) error {
	url := fmt.Sprintf("firewall_policies/%s/firewall_rules/%s", policyID, ruleID)
	req, err := c.NewRequest(http.MethodDelete, url, nil, nil)
	if err != nil {
		return fmt.Errorf("cannot create new delete request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute delete request: %v", err)
	}

	return nil
}
