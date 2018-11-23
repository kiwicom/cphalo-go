package api

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	CreatedAt              time.Time `json:"created_at"`
	ID                     string    `json:"id"`
	URL                    string    `json:"url"`
	Hostname               string    `json:"hostname"`
	ServerLabel            string    `json:"server_label"`
	ReportedFQDN           string    `json:"reported_fqdn"`
	PrimaryIPAddress       string    `json:"primary_ip_address"`
	ConnectingIPAddress    string    `json:"connecting_ip_address"`
	State                  string    `json:"state"`
	DaemonVersion          string    `json:"daemon_version"`
	ReadOnly               bool      `json:"read_only"`
	Platform               string    `json:"platform"`
	PlatformVersion        string    `json:"platform_version"`
	OSVersion              string    `json:"os_version"`
	KernelName             string    `json:"kernel_name"`
	KernelMachine          string    `json:"kernel_machine"`
	SelfVerificationFailed bool      `json:"self_verification_failed"`
	ConnectingIPFQDN       string    `json:"connecting_ip_fqdn"`
	LastStateChange        time.Time `json:"last_state_change"`
	DockerInspection       string    `json:"docker_inspection"`
	GroupID                string    `json:"group_id"`
	GroupName              string    `json:"group_name"`
	GroupPath              string    `json:"group_path"`
}

type ListServersResponse struct {
	Count   int      `json:"count"`
	Servers []Server `json:"servers"`
}

type GetServersResponse struct {
	Server Server `json:"server"`
}

func (c *Client) ListServers() (response ListServersResponse, err error) {
	req, err := c.NewRequest(http.MethodGet, "servers", nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute request: %v", err)
	}

	return response, nil
}

func (c *Client) GetServer(ID string) (response GetServersResponse, err error) {
	req, err := c.NewRequest(http.MethodGet, "servers/"+ID, nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) DeleteServer(ID string) error {
	return nil
}
