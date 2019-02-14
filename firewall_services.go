package cphalo

import (
	"fmt"
	"net/http"
)

type FirewallService struct {
	ID       string `json:"id,omitempty"`
	URL      string `json:"url,omitempty"`
	Name     string `json:"name,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Port     string `json:"port,omitempty"`
	System   bool   `json:"system,omitempty"`
}

type ListFirewallServicesResponse struct {
	Count    int               `json:"count"`
	Services []FirewallService `json:"firewall_services"`
}

type GetFirewallServiceResponse struct {
	Service FirewallService `json:"firewall_service"`
}

type CreateFirewallServiceResponse = GetFirewallServiceResponse
type CreateFirewallServiceRequest = GetFirewallServiceResponse
type UpdateFirewallServiceRequest = GetFirewallServiceResponse

func (c *Client) ListFirewallServices() (response ListFirewallServicesResponse, err error) {
	req, err := c.newRequest(http.MethodGet, "firewall_services", nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute request: %v", err)
	}

	return response, nil
}

func (c *Client) GetFirewallService(ID string) (response GetFirewallServiceResponse, err error) {
	req, err := c.newRequest(http.MethodGet, "firewall_services/"+ID, nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) CreateFirewallService(service FirewallService) (response CreateFirewallServiceResponse, err error) {
	req, err := c.newRequest(http.MethodPost, "firewall_services", nil, CreateFirewallServiceRequest{Service: service})
	if err != nil {
		return response, fmt.Errorf("cannot create new create request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute create request: %v", err)
	}

	return response, nil
}

func (c *Client) UpdateFirewallService(service FirewallService) error {
	req, err := c.newRequest(http.MethodPut, "firewall_services/"+service.ID, nil, UpdateFirewallServiceRequest{Service: service})
	if err != nil {
		return fmt.Errorf("cannot create new update request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute update request: %v", err)
	}

	return nil
}

func (c *Client) DeleteFirewallService(ID string) error {
	req, err := c.newRequest(http.MethodDelete, "firewall_services/"+ID, nil, nil)
	if err != nil {
		return fmt.Errorf("cannot create new delete request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute delete request: %v", err)
	}

	return nil
}
