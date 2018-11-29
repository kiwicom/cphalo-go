package api

import (
	"fmt"
	"net/http"
)

type CreateCSPAccountRequest struct {
	ExternalID     string `json:"external_id"`
	RoleArn        string `json:"role_arn"`
	SnsArn         string `json:"sns_arn,omitempty"`
	GroupID        string `json:"group_id"`
	CSPAccountType string `json:"csp_account_type"`
}

type CSPAccount struct {
	ID                 string `json:"id"`
	InitialScanSummary struct {
		S3         string `json:"s3"`
		Route53    string `json:"route53"`
		Lambda     string `json:"lambda"`
		Iam        string `json:"iam"`
		Ec2        string `json:"ec2"`
		Vpc        string `json:"vpc"`
		CloudTrail string `json:"cloud_trail"`
		APIGateway string `json:"api_gateway"`
	} `json:"initial_scan_summary"`
	ScanStatus               string      `json:"scan_status"`
	InitialRulesRunCompleted bool        `json:"initial_rules_run_completed"`
	CspAccountType           string      `json:"csp_account_type"`
	AccountDisplayName       string      `json:"account_display_name"`
	ErrorDetail              string      `json:"error_detail"`
	CreatedAt                string      `json:"created_at"`
	ExternalID               string      `json:"external_id"`
	TimeOfLastScan           string      `json:"time_of_last_scan"`
	InitialScanCompleted     bool        `json:"initial_scan_completed"`
	SnsStatus                interface{} `json:"sns_status"`
	RoleArn                  string      `json:"role_arn"`
	UpdatedAt                string      `json:"updated_at"`
	GroupID                  string      `json:"group_id"`
	UserID                   string      `json:"user_id"`
	CSPAccountAlias          string      `json:"csp_account_alias"`
	CSPAccountID             string      `json:"csp_account_id"`
	MonitoringState          string      `json:"monitoring_state"`
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
