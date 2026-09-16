package client

import (
	"encoding/json"
	"fmt"
)

// ── Request structs ───────────────────────────────────────────────────────

type IPSecCreateRequest struct {
	Name         string `json:"name"`
	DCSlug       string `json:"dcslug"`
	VPC          string `json:"vpc"`
	CPUModel     string `json:"cpumodel,omitempty"`
	BillingCycle string `json:"billingcycle"`
}

type IPSecConnectionRequest struct {
	IPSecID          string `json:"ipsecid"`
	ID               string `json:"id,omitempty"`
	Name             string `json:"name"`
	RemoteIP         string `json:"remote_ip"`
	RemoteLocalIP    string `json:"remote_local_ip"`
	LocalIP          string `json:"local_ip"`
	PSK              string `json:"psk"`
	Phase1Encryption string `json:"phase1-encryption"`
	Phase2Encryption string `json:"phase2-encryption"`
	Phase1Integrity  string `json:"phase1-integrity"`
	Phase2Integrity  string `json:"phase2-integrity"`
	Phase1DHGroup    string `json:"phase1-dh-group"`
	Phase2DHGroup    string `json:"phase2-dh-group"`
	IKEVersion       string `json:"ike_version"`
	Phase1Lifetime   string `json:"phase1_lifetime"`
	Phase2Lifetime   string `json:"phase2_lifetime"`
	RekeyMargin      string `json:"rekey_margin"`
	RekeyFuzz        string `json:"rekey_fuzz"`
	ReplayWindow     string `json:"replay_window"`
	DPDTimeout       string `json:"dpd_timeout"`
	DPDAction        string `json:"dpd_action"`
	StartupAction    string `json:"startup_action"`
}

// ── Response structs ──────────────────────────────────────────────────────

type IPSecInstance struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	DCSlug       string `json:"dcslug"`
	PSK          string `json:"psk"`
	Status       string `json:"status"`
	BillingCycle string `json:"billingcycle"`
	CreatedAt    string `json:"created_at"`
}

type IPSecListResponse struct {
	Status string          `json:"status"`
	Data   []IPSecInstance `json:"data"`
}

type IPSecDetailResponse struct {
	Status string        `json:"status"`
	Data   IPSecInstance `json:"data"`
}

type IPSecConnection struct {
	ID               string      `json:"id"`
	IPSecID          string      `json:"ipsecid"`
	Name             string      `json:"name"`
	RemoteIP         string      `json:"remote_ip"`
	RemoteLocalIP    string      `json:"remote_local_ip"`
	LocalIP          string      `json:"local_ip"`
	PSK              string      `json:"psk"`
	Phase1Encryption string      `json:"phase1-encryption"`
	Phase2Encryption string      `json:"phase2-encryption"`
	Phase1Integrity  string      `json:"phase1-integrity"`
	Phase2Integrity  string      `json:"phase2-integrity"`
	Phase1DHGroup    interface{} `json:"phase1-dh-group"`
	Phase2DHGroup    interface{} `json:"phase2-dh-group"`
	IKEVersion       string      `json:"ike_version"`
	Phase1Lifetime   string      `json:"phase1_lifetime"`
	Phase2Lifetime   string      `json:"phase2_lifetime"`
	RekeyMargin      string      `json:"rekey_margin"`
	RekeyFuzz        string      `json:"rekey_fuzz"`
	ReplayWindow     string      `json:"replay_window"`
	DPDTimeout       string      `json:"dpd_timeout"`
	DPDAction        string      `json:"dpd_action"`
	StartupAction    string      `json:"startup_action"`
	Status           string      `json:"status"`
	CreatedAt        string      `json:"created_at"`
}

type IPSecConnectionListResponse struct {
	Status string            `json:"status"`
	Data   []IPSecConnection `json:"data"`
}

// ── IPSec Tunnel methods ──────────────────────────────────────────────────

func (c *Client) CreateIPSec(req *IPSecCreateRequest) (string, error) {
	respBytes, err := c.Post("/ipsec?action=create", req)
	if err != nil {
		return "", fmt.Errorf("failed to create IPSec tunnel: %w", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", fmt.Errorf("failed to parse create IPSec response: %w", err)
	}
	if result["status"] != "success" {
		return "", fmt.Errorf("create IPSec failed: %s", result["message"])
	}
	return fmt.Sprintf("%v", result["id"]), nil
}

func (c *Client) GetIPSec(ipsecID string) (*IPSecInstance, error) {
	respBytes, err := c.Get(fmt.Sprintf("/ipsec?action=info&ipsecid=%s", ipsecID))
	if err != nil {
		return nil, fmt.Errorf("failed to get IPSec tunnel: %w", err)
	}
	var resp IPSecDetailResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse get IPSec response: %w", err)
	}
	if resp.Data.ID == "" {
		return nil, nil
	}
	return &resp.Data, nil
}

func (c *Client) DeleteIPSec(ipsecID string) error {
	respBytes, err := c.Delete(fmt.Sprintf("/ipsec?action=delete&ipsecid=%s", ipsecID))
	if err != nil {
		return fmt.Errorf("failed to delete IPSec tunnel: %w", err)
	}
	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse delete IPSec response: %w", err)
	}
	if result["status"] != "success" {
		return fmt.Errorf("delete IPSec failed: %s", result["message"])
	}
	return nil
}

// ── IPSec Connection methods ──────────────────────────────────────────────

func (c *Client) CreateIPSecConnection(req *IPSecConnectionRequest) (string, error) {
	respBytes, err := c.Post("/ipsec?action=connection", req)
	if err != nil {
		return "", fmt.Errorf("failed to create IPSec connection: %w", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", fmt.Errorf("failed to parse create IPSec connection response: %w", err)
	}
	if result["status"] != "success" {
		return "", fmt.Errorf("create IPSec connection failed: %s", result["message"])
	}
	return fmt.Sprintf("%v", result["id"]), nil
}

func (c *Client) UpdateIPSecConnection(req *IPSecConnectionRequest) error {
	respBytes, err := c.Put("/ipsec?action=connection", req)
	if err != nil {
		return fmt.Errorf("failed to update IPSec connection: %w", err)
	}
	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse update IPSec connection response: %w", err)
	}
	if result["status"] != "success" {
		return fmt.Errorf("update IPSec connection failed: %s", result["message"])
	}
	return nil
}

func (c *Client) ListIPSecConnections(ipsecID string) ([]IPSecConnection, error) {
	respBytes, err := c.Get(fmt.Sprintf("/ipsec?action=connection&ipsecid=%s", ipsecID))
	if err != nil {
		return nil, fmt.Errorf("failed to list IPSec connections: %w", err)
	}
	var resp IPSecConnectionListResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse list IPSec connections response: %w", err)
	}
	return resp.Data, nil
}

func (c *Client) DeleteIPSecConnection(ipsecID string, connectionID string) error {
	endpoint := fmt.Sprintf("/ipsec?action=connection&ipsecid=%s&id=%s", ipsecID, connectionID)
	respBytes, err := c.Delete(endpoint)
	if err != nil {
		return fmt.Errorf("failed to delete IPSec connection: %w", err)
	}
	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse delete IPSec connection response: %w", err)
	}
	if result["status"] != "success" {
		return fmt.Errorf("delete IPSec connection failed: %s", result["message"])
	}
	return nil
}
