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

// PowerAction performs a power action on a cloud instance
func (c *Client) PowerAction(cloudID string, action string) error {
	endpoint := fmt.Sprintf("/cloud/%s/%s", cloudID, action)
	respBytes, err := c.Post(endpoint, nil)
	if err != nil {
		return fmt.Errorf("power action '%s' failed: %w", action, err)
	}

	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse power action response: %w", err)
	}

	if result["status"] != "success" {
		return fmt.Errorf("power action failed: %s", result["message"])
	}

	return nil
}

// ListClouds returns all cloud instances in the account
func (c *Client) ListClouds() ([]CloudInstance, error) {
	var allInstances []CloudInstance
	page := 1

	for {
		endpoint := fmt.Sprintf("/cloud?perpage=100&page=%d", page)
		respBytes, err := c.Get(endpoint)
		if err != nil {
			return nil, fmt.Errorf("failed to list cloud instances: %w", err)
		}

		var listResp CloudListResponse
		if err := json.Unmarshal(respBytes, &listResp); err != nil {
			return nil, fmt.Errorf("failed to parse list response: %w", err)
		}

		allInstances = append(allInstances, listResp.Cloud...)

		if page >= listResp.Meta.TotalPages || listResp.Meta.TotalPages == 0 {
			break
		}
		page++
	}

	return allInstances, nil
}

// AttachFirewall attaches a security group to a cloud instance
func (c *Client) AttachFirewall(firewallID string, cloudID string) error {
	endpoint := fmt.Sprintf("/firewall/%s/server/add", firewallID)
	payload := map[string]string{
		"cloudid": cloudID,
	}

	respBytes, err := c.Post(endpoint, payload)
	if err != nil {
		return fmt.Errorf("failed to attach firewall: %w", err)
	}

	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse attach firewall response: %w", err)
	}

	if result["status"] != "success" {
		return fmt.Errorf("attach firewall failed: %s", result["message"])
	}

	return nil
}

// DetachFirewall detaches a security group from a cloud instance
func (c *Client) DetachFirewall(firewallID string, cloudID string) error {
	endpoint := fmt.Sprintf("/firewall/%s/server/%s/delete", firewallID, cloudID)
	respBytes, err := c.Delete(endpoint)
	if err != nil {
		return fmt.Errorf("failed to detach firewall: %w", err)
	}

	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse detach firewall response: %w", err)
	}

	if result["status"] != "success" {
		return fmt.Errorf("detach firewall failed: %s", result["message"])
	}

	return nil
}

// ── General Storage ───────────────────────────────────────────────────────

// AddGeneralStorage adds a general disk to a cloud instance
func (c *Client) AddGeneralStorage(cloudID string, sizeGB int) error {
	endpoint := fmt.Sprintf("/cloud/%s/storage/add", cloudID)
	payload := map[string]int{
		"size": sizeGB,
	}

	respBytes, err := c.Post(endpoint, payload)
	if err != nil {
		return fmt.Errorf("failed to add storage: %w", err)
	}

	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse add storage response: %w", err)
	}

	if result["status"] != "success" {
		return fmt.Errorf("add storage failed: %s", result["message"])
	}

	return nil
}

// DeleteGeneralStorage deletes a general disk from a cloud instance
func (c *Client) DeleteGeneralStorage(cloudID string, diskID string) error {
	endpoint := fmt.Sprintf("/cloud/%s/storage/%s/delete", cloudID, diskID)
	respBytes, err := c.Delete(endpoint)
	if err != nil {
		return fmt.Errorf("failed to delete storage: %w", err)
	}

	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse delete storage response: %w", err)
	}

	if result["status"] != "success" {
		return fmt.Errorf("delete storage failed: %s", result["message"])
	}

	return nil
}

// UpdateGeneralStorage updates a general disk on a cloud instance
func (c *Client) UpdateGeneralStorage(cloudID string, diskID string, bus string, diskType string) error {
	endpoint := fmt.Sprintf("/cloud/%s/storage/%s/update", cloudID, diskID)
	payload := map[string]string{
		"bus":  bus,
		"type": diskType,
	}

	respBytes, err := c.Post(endpoint, payload)
	if err != nil {
		return fmt.Errorf("failed to update storage: %w", err)
	}

	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse update storage response: %w", err)
	}

	if result["status"] != "success" {
		return fmt.Errorf("update storage failed: %s", result["message"])
	}

	return nil
}

// ── EBS Storage ───────────────────────────────────────────────────────────

// AttachEBS attaches an EBS volume to a cloud instance
func (c *Client) AttachEBS(ebsID string, cloudID string) error {
	endpoint := fmt.Sprintf("/ebs/%s/attach", ebsID)
	payload := map[string]string{
		"type":       "cloud",
		"resourceid": cloudID,
	}

	respBytes, err := c.Put(endpoint, payload)
	if err != nil {
		return fmt.Errorf("failed to attach EBS: %w", err)
	}

	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse attach EBS response: %w", err)
	}

	if result["status"] != "success" {
		return fmt.Errorf("attach EBS failed: %s", result["message"])
	}

	return nil
}

// DetachEBS detaches an EBS volume from a cloud instance
func (c *Client) DetachEBS(ebsID string, cloudID string) error {
	endpoint := fmt.Sprintf("/ebs/%s/dettach", ebsID)
	payload := map[string]string{
		"type":       "cloud",
		"resourceid": cloudID,
	}

	respBytes, err := c.Put(endpoint, payload)
	if err != nil {
		return fmt.Errorf("failed to detach EBS: %w", err)
	}

	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse detach EBS response: %w", err)
	}

	if result["status"] != "success" {
		return fmt.Errorf("detach EBS failed: %s", result["message"])
	}

	return nil
}

// UpdateEBS updates an attached EBS volume on a cloud instance
func (c *Client) UpdateEBS(cloudID string, ebsID string, bus string, diskType string) error {
	endpoint := fmt.Sprintf("/cloud/%s/storage/%s/update", cloudID, ebsID)
	payload := map[string]string{
		"bus":  bus,
		"type": diskType,
	}

	respBytes, err := c.Post(endpoint, payload)
	if err != nil {
		return fmt.Errorf("failed to update EBS: %w", err)
	}

	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse update EBS response: %w", err)
	}

	if result["status"] != "success" {
		return fmt.Errorf("update EBS failed: %s", result["message"])
	}

	return nil
}

// ── Snapshot ──────────────────────────────────────────────────────────────

// CreateSnapshot creates a snapshot of a cloud instance
func (c *Client) CreateSnapshot(cloudID string, name string) error {
	endpoint := fmt.Sprintf("/cloud/%s/snapshot/create", cloudID)
	payload := map[string]string{
		"name": name,
	}

	respBytes, err := c.Post(endpoint, payload)
	if err != nil {
		return fmt.Errorf("failed to create snapshot: %w", err)
	}

	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse create snapshot response: %w", err)
	}

	if result["status"] != "success" {
		return fmt.Errorf("create snapshot failed: %s", result["message"])
	}

	return nil
}

// DeleteSnapshot deletes a snapshot of a cloud instance
func (c *Client) DeleteSnapshot(cloudID string, snapshotID string) error {
	endpoint := fmt.Sprintf("/cloud/%s/snapshot/%s/delete", cloudID, snapshotID)
	respBytes, err := c.Delete(endpoint)
	if err != nil {
		return fmt.Errorf("failed to delete snapshot: %w", err)
	}

	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse delete snapshot response: %w", err)
	}

	if result["status"] != "success" {
		return fmt.Errorf("delete snapshot failed: %s", result["message"])
	}

	return nil
}

// RestoreSnapshot restores a cloud instance from a snapshot
func (c *Client) RestoreSnapshot(cloudID string, snapshotID string) error {
	endpoint := fmt.Sprintf("/cloud/%s/snapshot/%s/restore", cloudID, snapshotID)
	respBytes, err := c.Post(endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to restore snapshot: %w", err)
	}

	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse restore snapshot response: %w", err)
	}

	if result["status"] != "success" {
		return fmt.Errorf("restore snapshot failed: %s", result["message"])
	}

	return nil
}
