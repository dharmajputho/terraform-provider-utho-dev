package client

import (
	"encoding/json"
	"fmt"
)

// ── Request structs ───────────────────────────────────────────────────────
type CloudDeployRequest struct {
	DCSlug         string        `json:"dcslug"`
	PlanID         string        `json:"planid"`
	BillingCycle   string        `json:"billingcycle"`
	Auth           string        `json:"auth"`
	EnablePublicIP string        `json:"enable_publicip"`
	SubnetRequired string        `json:"subnetRequired,omitempty"`
	CPUModel       string        `json:"cpumodel,omitempty"`
	EnableBackup   string        `json:"enablebackup,omitempty"`
	Support        string        `json:"support,omitempty"`
	Firewall       string        `json:"firewall,omitempty"`
	RootPassword   string        `json:"root_password,omitempty"`
	SSHKeys        string        `json:"sshkeys,omitempty"`
	Image          string        `json:"image,omitempty"`
	VPC            string        `json:"vpc,omitempty"`
	SnapshotID     string        `json:"snapshotid,omitempty"`
	BackupID       string        `json:"backupid,omitempty"`
	ISO            string        `json:"iso,omitempty"`
	Stack          string        `json:"stack,omitempty"`
	EBS            []EBSVolume   `json:"ebs,omitempty"`
	Cloud          []CloudServer `json:"cloud"`
}

type EBSVolume struct {
	ID   string `json:"id"`
	Disk int    `json:"disk"`
	Type string `json:"type"`
}

type CloudServer struct {
	Hostname string `json:"hostname"`
}

// ── Response structs ──────────────────────────────────────────────────────

type CloudListResponse struct {
	Cloud []CloudInstance `json:"cloud"`
	Meta  CloudMeta       `json:"meta"`
}

type CloudMeta struct {
	Total       int `json:"total"`
	TotalPages  int `json:"totalpages"`
	CurrentPage int `json:"currentpage"`
}

type CloudInstance struct {
	CloudID      string          `json:"cloudid"`
	Hostname     string          `json:"hostname"`
	Status       string          `json:"status"`
	PowerStatus  string          `json:"powerstatus"`
	IP           string          `json:"ip"`
	CPU          string          `json:"cpu"`
	RAM          string          `json:"ram"`
	DiskSize     int             `json:"disksize"`
	DCSlug       string          `json:"dcslug"`
	BillingCycle string          `json:"billingcycle"`
	PlanSlug     string          `json:"planslug"`
	Image        CloudImage      `json:"image"`
	DCLocation   CloudDCLocation `json:"dclocation"`
	Plan         CloudPlan       `json:"plan"`
	CreatedAt    string          `json:"created_at"`
}

type CloudImage struct {
	Name         string `json:"name"`
	Distribution string `json:"distribution"`
	Version      string `json:"version"`
	Image        string `json:"image"`
}

type CloudDCLocation struct {
	Location string `json:"location"`
	Country  string `json:"country"`
	DC       string `json:"dc"`
}

type CloudPlan struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	CPU  string `json:"cpu"`
	RAM  string `json:"ram"`
	Disk string `json:"disk"`
}

type CloudDeleteResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type CloudDeploySuccessResponse struct {
	Status  string `json:"status"`
	CloudID string `json:"cloudid"`
	Message string `json:"message"`
	IPv4    string `json:"ipv4"`
}

// ── API methods ───────────────────────────────────────────────────────────

// CreateCloud deploys a new cloud instance
func (c *Client) CreateCloud(req *CloudDeployRequest) (*CloudInstance, error) {
	respBytes, err := c.Post("/cloud/deploy", req)
	if err != nil {
		return nil, fmt.Errorf("failed to create cloud instance: %w", err)
	}

	// Parse the deploy response
	var deployResp CloudDeploySuccessResponse
	if err := json.Unmarshal(respBytes, &deployResp); err != nil {
		return nil, fmt.Errorf("failed to parse create response: %w", err)
	}

	if deployResp.Status != "success" {
		return nil, fmt.Errorf("deploy failed: %s", deployResp.Message)
	}

	if deployResp.CloudID == "" {
		return nil, fmt.Errorf("no cloudid returned in response")
	}

	// Return a CloudInstance with what we have from deploy response
	// IP and full details will be fetched by Read() after creation
	return &CloudInstance{
		CloudID: deployResp.CloudID,
		IP:      deployResp.IPv4,
		Status:  "Active",
	}, nil
}

// GetCloud fetches a specific cloud instance by ID
func (c *Client) GetCloud(cloudID string) (*CloudInstance, error) {
	endpoint := fmt.Sprintf("/cloud/%s", cloudID)
	respBytes, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to get cloud instance: %w", err)
	}

	var listResp CloudListResponse
	if err := json.Unmarshal(respBytes, &listResp); err != nil {
		return nil, fmt.Errorf("failed to parse get response: %w", err)
	}

	if len(listResp.Cloud) == 0 {
		return nil, nil // not found — deleted outside terraform
	}

	return &listResp.Cloud[0], nil
}

// DeleteCloud destroys a cloud instance
func (c *Client) DeleteCloud(cloudID string, deleteEBS bool) error {
	ebsParam := "no"
	if deleteEBS {
		ebsParam = "yes"
	}

	endpoint := fmt.Sprintf("/cloud/%s/destroy?ebs_delete=%s", cloudID, ebsParam)
	respBytes, err := c.Delete(endpoint)
	if err != nil {
		return fmt.Errorf("failed to delete cloud instance: %w", err)
	}

	var deleteResp CloudDeleteResponse
	if err := json.Unmarshal(respBytes, &deleteResp); err != nil {
		return fmt.Errorf("failed to parse delete response: %w", err)
	}

	if deleteResp.Status != "success" {
		return fmt.Errorf("delete failed: %s", deleteResp.Message)
	}

	return nil
}
