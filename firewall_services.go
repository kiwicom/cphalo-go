package cphalo

import (
	"fmt"
	"net/http"
)

// FirewallService represent a CPHalo firewall service.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#object-representation-9
type FirewallService struct {
	ID       string `json:"id,omitempty"`
	URL      string `json:"url,omitempty"`
	Name     string `json:"name,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Port     string `json:"port,omitempty"`
	System   bool   `json:"system,omitempty"`
}

// ListFirewallServicesResponse represent a list of firewall services response.
type ListFirewallServicesResponse struct {
	Count    int               `json:"count"`
	Services []FirewallService `json:"firewall_services"`
}

// GetFirewallServiceResponse represent a get firewall service response.
type GetFirewallServiceResponse struct {
	Service FirewallService `json:"firewall_service"`
}

// CreateFirewallServiceResponse represent a create firewall service response.
type CreateFirewallServiceResponse = GetFirewallServiceResponse

// CreateFirewallServiceRequest represent a create firewall service request.
type CreateFirewallServiceRequest = GetFirewallServiceResponse

// UpdateFirewallServiceRequest represent a update firewall service request.
type UpdateFirewallServiceRequest = GetFirewallServiceResponse

// ListFirewallServices lists all firewall service.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#list-firewall-services
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

// GetFirewallService returns details of the firewall service.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#get-firewall-service-details
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

// CreateFirewallService creates a new firewall service.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#create-a-new-firewall-service
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

// UpdateFirewallService updates firewall service.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/cloudpassage-api-documentation#firewall-services
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

// DeleteFirewallService deletes a firewall service.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#delete-firewall-service
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
