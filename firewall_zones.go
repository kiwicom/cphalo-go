package api

import (
	"fmt"
	"net/http"
)

type FirewallZone struct {
	ID        string `json:"id,omitempty"`
	URL       string `json:"url,omitempty"`
	Name      string `json:"name,omitempty"`
	IpAddress string `json:"ip_address,omitempty"`
}

type ListFirewallZonesResponse struct {
	Count int            `json:"count"`
	Zones []FirewallZone `json:"firewall_zones"`
}

type GetFirewallZoneResponse struct {
	Zone FirewallZone `json:"firewall_zone"`
}

type CreateFirewallZoneResponse = GetFirewallZoneResponse
type CreateFirewallZoneRequest = GetFirewallZoneResponse
type UpdateFirewallZoneRequest = GetFirewallZoneResponse

func (c *Client) ListFirewallZones() (response ListFirewallZonesResponse, err error) {
	req, err := c.NewRequest(http.MethodGet, "firewall_zones", nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute request: %v", err)
	}

	return response, nil
}

func (c *Client) GetFirewallZone(ID string) (response GetFirewallZoneResponse, err error) {
	req, err := c.NewRequest(http.MethodGet, "firewall_zones/"+ID, nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) CreateFirewallZone(zone FirewallZone) (response CreateFirewallZoneResponse, err error) {
	req, err := c.NewRequest(http.MethodPost, "firewall_zones", nil, CreateFirewallZoneRequest{Zone: zone})
	if err != nil {
		return response, fmt.Errorf("cannot create new create request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute create request: %v", err)
	}

	return response, nil
}

func (c *Client) UpdateFirewallZone(zone FirewallZone) error {
	req, err := c.NewRequest(http.MethodPut, "firewall_zones/"+zone.ID, nil, UpdateFirewallZoneRequest{Zone: zone})
	if err != nil {
		return fmt.Errorf("cannot create new update request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute update request: %v", err)
	}

	return nil
}

func (c *Client) DeleteFirewallZone(ID string) error {
	req, err := c.NewRequest(http.MethodDelete, "firewall_zones/"+ID, nil, nil)
	if err != nil {
		return fmt.Errorf("cannot create new delete request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute delete request: %v", err)
	}

	return nil
}
