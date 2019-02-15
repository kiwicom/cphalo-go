package cphalo

import (
	"fmt"
	"net/http"
)

// FirewallRuleSourceTarget represent a CPHalo firewall source and target.
type FirewallRuleSourceTarget struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	IPAddress string `json:"ip_address,omitempty"`
	Kind      string `json:"type,omitempty"`
}

// GetID returns ID if exists, otherwise Name
func (f *FirewallRuleSourceTarget) GetID() string {
	if f.ID == "" {
		return f.Name
	}

	return f.ID
}

// FirewallRule represent a CPHalo firewall rule.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#object-representation-7
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

// ListFirewallRulesResponse represent a list of firewall rules response.
type ListFirewallRulesResponse struct {
	Count int            `json:"count"`
	Rules []FirewallRule `json:"firewall_rules"`
}

// GetFirewallRuleResponse represent a get firewall rule response.
type GetFirewallRuleResponse struct {
	Rule FirewallRule `json:"firewall_rule"`
}

// CreateFirewallRuleResponse represent a create firewall rule response.
type CreateFirewallRuleResponse = GetFirewallRuleResponse

// CreateFirewallRuleRequest represent a create firewall rule request.
type CreateFirewallRuleRequest = GetFirewallRuleResponse

// UpdateFirewallRuleRequest represent a create firewall rule request.
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

// ListFirewallRules lists all firewall rules.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#list-firewall-rules-in-firewall-policy
func (c *Client) ListFirewallRules(policyID string) (response ListFirewallRulesResponse, err error) {
	url := fmt.Sprintf("firewall_policies/%s/firewall_rules", policyID)
	req, err := c.newRequest(http.MethodGet, url, nil, nil)
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

// GetFirewallRule returns details of the firewall rule.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#get-firewall-rule-details
func (c *Client) GetFirewallRule(policyID, ruleID string) (response GetFirewallRuleResponse, err error) {
	url := fmt.Sprintf("firewall_policies/%s/firewall_rules/%s", policyID, ruleID)
	req, err := c.newRequest(http.MethodGet, url, nil, nil)
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

// CreateFirewallRule creates a new firewall rule.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#add-new-firewall-rule-to-the-firewall-policy
func (c *Client) CreateFirewallRule(policyID string, rule FirewallRule) (response CreateFirewallRuleResponse, err error) {
	url := fmt.Sprintf("firewall_policies/%s/firewall_rules", policyID)
	rule.applyCorrections()
	req, err := c.newRequest(http.MethodPost, url, nil, CreateFirewallRuleRequest{Rule: rule})
	if err != nil {
		return response, fmt.Errorf("cannot create new create request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute create request: %v", err)
	}

	return response, nil
}

// UpdateFirewallRule updates firewall rule.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#update-firewall-rule
func (c *Client) UpdateFirewallRule(policyID string, rule FirewallRule) error {
	url := fmt.Sprintf("firewall_policies/%s/firewall_rules/%s", policyID, rule.ID)
	rule.applyCorrections()
	req, err := c.newRequest(http.MethodPut, url, nil, UpdateFirewallRuleRequest{Rule: rule})
	if err != nil {
		return fmt.Errorf("cannot create new update request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute update request: %v", err)
	}

	return nil
}

// DeleteFirewallRule deletes a firewall rule.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#delete-firewall-rule
func (c *Client) DeleteFirewallRule(policyID, ruleID string) error {
	url := fmt.Sprintf("firewall_policies/%s/firewall_rules/%s", policyID, ruleID)
	req, err := c.newRequest(http.MethodDelete, url, nil, nil)
	if err != nil {
		return fmt.Errorf("cannot create new delete request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute delete request: %v", err)
	}

	return nil
}
