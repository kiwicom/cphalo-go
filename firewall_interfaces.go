package cphalo

import (
	"fmt"
	"net/http"
)

// FirewallInterface represent a CPHalo firewall interface.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#object-representation-8
type FirewallInterface struct {
	ID     string `json:"id,omitempty"`
	URL    string `json:"url,omitempty"`
	Name   string `json:"name,omitempty"`
	System bool   `json:"system,omitempty"`
}

// ListFirewallInterfacesResponse represent a list of firewall interfaces response.
type ListFirewallInterfacesResponse struct {
	Count      int                 `json:"count"`
	Interfaces []FirewallInterface `json:"firewall_interfaces"`
}

// GetFirewallInterfaceResponse represent a get firewall interface response.
type GetFirewallInterfaceResponse struct {
	Interface FirewallInterface `json:"firewall_interface"`
}

// CreateFirewallInterfaceResponse represent a create firewall interface response.
type CreateFirewallInterfaceResponse = GetFirewallInterfaceResponse

// CreateFirewallInterfaceRequest represent a create firewall interface request.
type CreateFirewallInterfaceRequest = GetFirewallInterfaceResponse

// UpdateFirewallInterfaceRequest represent a update firewall interface request.
type UpdateFirewallInterfaceRequest = GetFirewallInterfaceResponse

// ListFirewallInterfaces lists all firewall interfaces.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#list-firewall-interfaces
func (c *Client) ListFirewallInterfaces() (response ListFirewallInterfacesResponse, err error) {
	req, err := c.newRequest(http.MethodGet, "firewall_interfaces", nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute request: %v", err)
	}

	return response, nil
}

// GetFirewallInterface returns details of the firewall interface.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#get-firewall-interface-details
func (c *Client) GetFirewallInterface(ID string) (response GetFirewallInterfaceResponse, err error) {
	req, err := c.newRequest(http.MethodGet, "firewall_interfaces/"+ID, nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// CreateFirewallInterface creates a new firewall interface.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#create-a-new-firewall-interface
func (c *Client) CreateFirewallInterface(fwInterface FirewallInterface) (response CreateFirewallInterfaceResponse, err error) {
	req, err := c.newRequest(http.MethodPost, "firewall_interfaces", nil, CreateFirewallInterfaceRequest{Interface: fwInterface})
	if err != nil {
		return response, fmt.Errorf("cannot create new create request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute create request: %v", err)
	}

	return response, nil

}

// UpdateFirewallInterface updates firewall interface.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/cloudpassage-api-documentation#firewall-interfaces
func (c *Client) UpdateFirewallInterface(fwInterface FirewallInterface) error {
	req, err := c.newRequest(http.MethodPut, "firewall_interfaces/"+fwInterface.ID, nil, UpdateFirewallInterfaceRequest{Interface: fwInterface})
	if err != nil {
		return fmt.Errorf("cannot create new update request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute update request: %v", err)
	}

	return nil
}

// DeleteFirewallInterface deletes a firewall interface.
//
// CPHalo API Docs: https://library.cloudpassage.com/help/article/link/cloudpassage-api-documentation#delete-firewall-interface
func (c *Client) DeleteFirewallInterface(ID string) error {
	req, err := c.newRequest(http.MethodDelete, "firewall_interfaces/"+ID, nil, nil)
	if err != nil {
		return fmt.Errorf("cannot create new delete request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute delete request: %v", err)
	}

	return nil
}
