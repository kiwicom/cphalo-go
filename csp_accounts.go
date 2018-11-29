package api

import (
	"fmt"
	"net/http"
)

type CreateCSPAccountRequest struct {
	ExternalID         string `json:"external_id"`
	RoleArn            string `json:"role_arn"`
	SnsArn             string `json:"sns_arn,omitempty"`
	GroupID            string `json:"group_id"`
	CSPAccountType     string `json:"csp_account_type"`
	AccountDisplayName string `json:"account_display_name"`
}

type CSPAccount struct {
	ID                 string `json:"id"`
	InitialScanSummary struct {
		S3         string `json:"s3,omitempty"`
		Route53    string `json:"route53,omitempty"`
		Lambda     string `json:"lambda,omitempty"`
		Iam        string `json:"iam,omitempty"`
		Ec2        string `json:"ec2,omitempty"`
		Vpc        string `json:"vpc,omitempty"`
		CloudTrail string `json:"cloud_trail,omitempty"`
		APIGateway string `json:"api_gateway,omitempty"`
	} `json:"initial_scan_summary,omitempty"`
	ScanStatus               string      `json:"scan_status,omitempty"`
	InitialRulesRunCompleted bool        `json:"initial_rules_run_completed,omitempty"`
	CSPAccountType           string      `json:"csp_account_type,omitempty"`
	AccountDisplayName       string      `json:"account_display_name,omitempty"`
	ErrorDetail              string      `json:"error_detail,omitempty"`
	CreatedAt                string      `json:"created_at,omitempty"`
	ExternalID               string      `json:"external_id,omitempty"`
	TimeOfLastScan           string      `json:"time_of_last_scan,omitempty"`
	InitialScanCompleted     bool        `json:"initial_scan_completed,omitempty"`
	SnsStatus                interface{} `json:"sns_status,omitempty"`
	RoleArn                  string      `json:"role_arn,omitempty"`
	UpdatedAt                string      `json:"updated_at,omitempty"`
	GroupID                  string      `json:"group_id,omitempty"`
	UserID                   string      `json:"user_id,omitempty"`
	CSPAccountAlias          string      `json:"csp_account_alias,omitempty"`
	CSPAccountID             string      `json:"csp_account_id,omitempty"`
	MonitoringState          string      `json:"monitoring_state,omitempty"`
}

type ListCSPAccountsResponse struct {
	Count       int          `json:"count"`
	CSPAccounts []CSPAccount `json:"csp_accounts"`
}

type GetCSPAccountResponse struct {
	CSPAccount CSPAccount `json:"csp_account"`
}

type CreateCSPAccountResponse string

func (c *Client) ListCSPAccounts() (response ListCSPAccountsResponse, err error) {
	req, err := c.NewRequest(http.MethodGet, "csp_accounts", nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute request: %v", err)
	}

	return response, nil
}

func (c *Client) GetCSPAccount(ID string) (response GetCSPAccountResponse, err error) {
	req, err := c.NewRequest(http.MethodGet, "csp_accounts/"+ID, nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute request: %v", err)
	}

	return response, nil
}

func (c *Client) CreateCSPAccount(account CreateCSPAccountRequest) (response CreateCSPAccountResponse, err error) {
	// only AWS is supported at the moment
	account.CSPAccountType = "AWS"

	req, err := c.NewRequest(http.MethodPost, "csp_accounts", nil, account)
	if err != nil {
		return response, fmt.Errorf("cannot create new create request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, fmt.Errorf("cannot execute create request: %v", err)
	}

	return response, nil
}

func (c *Client) UpdateCSPAccount(account CSPAccount) error {
	aID := account.ID
	account.ID = ""

	req, err := c.NewRequest(http.MethodPut, "csp_accounts/"+aID, nil, account)
	if err != nil {
		return fmt.Errorf("cannot create new update request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute update request: %v", err)
	}

	return nil
}

func (c *Client) DeleteCSPAccount(ID string) error {
	req, err := c.NewRequest(http.MethodDelete, "csp_accounts/"+ID, nil, nil)
	if err != nil {
		return fmt.Errorf("cannot create new delete request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("cannot execute delete request: %v", err)
	}

	return nil
}
