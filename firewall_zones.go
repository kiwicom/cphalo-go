package cphalo

import (
	"fmt"
	"net/http"
)

// FirewallZone represent a CPHalo firewall zone.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#object-representation-10
type FirewallZone struct {
	ID          string `json:"id,omitempty"`
	URL         string `json:"url,omitempty"`
	Name        string `json:"name,omitempty"`
	IPAddress   string `json:"ip_address,omitempty"`
	Description string `json:"description,omitempty"`
	System      bool   `json:"system,omitempty"`
}

// ListFirewallZonesResponse represent a list of firewall zones response.
type ListFirewallZonesResponse struct {
	Count int            `json:"count"`
	Zones []FirewallZone `json:"firewall_zones"`
}

// GetFirewallZoneResponse represent a get firewall zone response.
type GetFirewallZoneResponse struct {
	Zone FirewallZone `json:"firewall_zone"`
}

// CreateFirewallZoneResponse represent a create firewall zone response.
type CreateFirewallZoneResponse = GetFirewallZoneResponse

// CreateFirewallZoneRequest represent a create firewall zone request.
type CreateFirewallZoneRequest = GetFirewallZoneResponse

// UpdateFirewallZoneRequest represent a update firewall zone request.
type UpdateFirewallZoneRequest = GetFirewallZoneResponse

// ListFirewallZones lists all firewall zones.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#list-firewall-zones
func (c *Client) ListFirewallZones() (response ListFirewallZonesResponse, err error) {
	req, err := c.newRequest(http.MethodGet, "firewall_zones", nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute request: %v", err)
	}

	return response, nil
}

// GetFirewallZone returns details of the firewall zone.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#list-firewall-zones
func (c *Client) GetFirewallZone(ID string) (response GetFirewallZoneResponse, err error) {
	req, err := c.newRequest(http.MethodGet, "firewall_zones/"+ID, nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// CreateFirewallZone creates a new firewall zone.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#create-a-new-firewall-zone
func (c *Client) CreateFirewallZone(zone FirewallZone) (response CreateFirewallZoneResponse, err error) {
	req, err := c.newRequest(http.MethodPost, "firewall_zones", nil, CreateFirewallZoneRequest{Zone: zone})
	if err != nil {
		return response, fmt.Errorf("cannot create new create request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute create request: %v", err)
	}

	return response, nil
}

// UpdateFirewallZone updates firewall zone.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#update-firewall-zone
func (c *Client) UpdateFirewallZone(zone FirewallZone) error {
	req, err := c.newRequest(http.MethodPut, "firewall_zones/"+zone.ID, nil, UpdateFirewallZoneRequest{Zone: zone})
	if err != nil {
		return fmt.Errorf("cannot create new update request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute update request: %v", err)
	}

	return nil
}

// DeleteFirewallZone deletes a firewall zone.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#delete-firewall-zone
func (c *Client) DeleteFirewallZone(ID string) error {
	req, err := c.newRequest(http.MethodDelete, "firewall_zones/"+ID, nil, nil)
	if err != nil {
		return fmt.Errorf("cannot create new delete request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute delete request: %v", err)
	}

	return nil
}
