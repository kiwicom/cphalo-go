package cphalo

import (
	"fmt"
	"net/http"
)

type CreateCSPAccountAWSRequest struct {
	ExternalID         string `json:"aws_external_id,omitempty"`
	RoleArn            string `json:"aws_role_arn,omitempty"`
	SnsArn             string `json:"aws_sns_arn,omitempty"`
	GroupID            string `json:"group_id,omitempty"`
	CSPAccountType     string `json:"csp_account_type,omitempty"`
	CSPRegionType      string `json:"csp_region_type,omitempty"`
	AccountDisplayName string `json:"account_display_name,omitempty"`
}

type CSPAccount struct {
	ID                       string `json:"id"`
	CSPAccountType           string `json:"csp_account_type,omitempty"`
	CSPRegionType            string `json:"csp_region_type,omitempty"`
	CSPAccountID             string `json:"csp_account_id,omitempty"`
	CSPAccountAlias          string `json:"csp_account_alias,omitempty"`
	AccountDisplayName       string `json:"account_display_name,omitempty"`
	CreatedAt                string `json:"created_at,omitempty"`
	UpdatedAt                string `json:"updated_at,omitempty"`
	UserID                   string `json:"user_id,omitempty"`
	GroupID                  string `json:"group_id,omitempty"`
	MonitoringState          string `json:"monitoring_state,omitempty"`
	InitialScanCompleted     bool   `json:"initial_scan_completed,omitempty"`
	InitialRulesRunCompleted bool   `json:"initial_rules_run_completed,omitempty"`
	InitialScanSummary       struct {
		S3         string `json:"s3,omitempty"`
		Route53    string `json:"route53,omitempty"`
		Lambda     string `json:"lambda,omitempty"`
		Iam        string `json:"iam,omitempty"`
		Ec2        string `json:"ec2,omitempty"`
		Vpc        string `json:"vpc,omitempty"`
		CloudTrail string `json:"cloud_trail,omitempty"`
		APIGateway string `json:"api_gateway,omitempty"`
	} `json:"initial_scan_summary,omitempty"`
	ScanStatus          string `json:"scan_status,omitempty"`
	ErrorDetail         string `json:"error_detail,omitempty"`
	TimeOfLastScan      string `json:"time_of_last_scan,omitempty"`
	AzureDirectoryID    string `json:"azure_directory_id,omitempty"`
	AzureApplicationID  string `json:"azure_application_id,omitempty"`
	AzureApplicationKey string `json:"azure_application_key,omitempty"`
	AWSAccessKey        string `json:"aws_access_key,omitempty"`
	AWSSecret           string `json:"aws_secret,omitempty"`
	AWSRoleArn          string `json:"aws_role_arn,omitempty"`
	AWSExternalID       string `json:"aws_external_id,omitempty"`
	AWSSnsStatus        string `json:"aws_sns_status,omitempty"`
	AwsSnsArn           string `json:"aws_sns_arn,omitempty"`
	AwsSnsErrorDetail   string `json:"aws_sns_error_detail,omitempty"`
}

type ListCSPAccountsResponse struct {
	Count       int          `json:"count"`
	CSPAccounts []CSPAccount `json:"csp_accounts"`
}

type GetCSPAccountResponse struct {
	CSPAccount CSPAccount `json:"csp_account"`
}

type CreateCSPAccountResponse GetCSPAccountResponse

func (c *Client) ListCSPAccounts() (response ListCSPAccountsResponse, err error) {
	req, err := c.newRequest(http.MethodGet, "csp_accounts", nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) GetCSPAccount(ID string) (response GetCSPAccountResponse, err error) {
	req, err := c.newRequest(http.MethodGet, "csp_accounts/"+ID, nil, nil)
	if err != nil {
		return response, fmt.Errorf("cannot create new request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) CreateCSPAccount(account CreateCSPAccountAWSRequest) (response CreateCSPAccountResponse, err error) {
	req, err := c.newRequest(http.MethodPost, "csp_accounts", nil, account)
	if err != nil {
		return response, fmt.Errorf("cannot create new create request: %v", err)
	}

	_, err = c.Do(req, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) UpdateCSPAccount(account CSPAccount) error {
	aID := account.ID
	account.ID = ""

	req, err := c.newRequest(http.MethodPut, "csp_accounts/"+aID, nil, account)
	if err != nil {
		return fmt.Errorf("cannot create new update request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteCSPAccount(ID string) error {
	req, err := c.newRequest(http.MethodDelete, "csp_accounts/"+ID, nil, nil)
	if err != nil {
		return fmt.Errorf("cannot create new delete request: %v", err)
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
