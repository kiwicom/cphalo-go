package api

import (
	"fmt"
	"net/http"
)

type FirewallInterface struct {
	ID     string `json:"id,omitempty"`
	URL    string `json:"url,omitempty"`
	Name   string `json:"name,omitempty"`
	System bool   `json:"system,omitempty"`
}

type ListFirewallInterfacesResponse struct {
	Count      int                 `json:"count"`
	Interfaces []FirewallInterface `json:"firewall_interfaces"`
}

type GetFirewallInterfaceResponse struct {
	Interface FirewallInterface `json:"firewall_interface"`
}

type CreateFirewallInterfaceResponse = GetFirewallInterfaceResponse
type CreateFirewallInterfaceRequest = GetFirewallInterfaceResponse
type UpdateFirewallInterfaceRequest = GetFirewallInterfaceResponse

func (c *Client) ListFirewallInterfaces() (response ListFirewallInterfacesResponse, err error) {
	req, err := c.NewRequest(http.MethodGet, "firewall_interfaces", nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute request: %v", err)
	}

	return response, nil
}

func (c *Client) GetFirewallInterface(ID string) (response GetFirewallInterfaceResponse, err error) {
	req, err := c.NewRequest(http.MethodGet, "firewall_interfaces/"+ID, nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) CreateFirewallInterface(fwInterface FirewallInterface) (response CreateFirewallInterfaceResponse, err error) {
	req, err := c.NewRequest(http.MethodPost, "firewall_interfaces", nil, CreateFirewallInterfaceRequest{Interface: fwInterface})
	if err != nil {
		return response, fmt.Errorf("cannot create new create request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute create request: %v", err)
	}

	return response, nil
}

func (c *Client) UpdateFirewallInterface(fwInterface FirewallInterface) error {
	req, err := c.NewRequest(http.MethodPut, "firewall_interfaces/"+fwInterface.ID, nil, UpdateFirewallInterfaceRequest{Interface: fwInterface})
	if err != nil {
		return fmt.Errorf("cannot create new update request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute update request: %v", err)
	}

	return nil
}

func (c *Client) DeleteFirewallInterface(ID string) error {
	req, err := c.NewRequest(http.MethodDelete, "firewall_interfaces/"+ID, nil, nil)
	if err != nil {
		return fmt.Errorf("cannot create new delete request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute delete request: %v", err)
	}

	return nil
}
