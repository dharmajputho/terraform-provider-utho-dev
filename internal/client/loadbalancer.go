package client

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// ── Request structs ───────────────────────────────────────────────────────

type LoadBalancerCreateRequest struct {
	Type           string `json:"type"`
	Name           string `json:"name"`
	DCSlug         string `json:"dcslug"`
	VPC            string `json:"vpc,omitempty"`
	EnablePublicIP string `json:"enable_publicip"`
	Firewall       string `json:"firewall,omitempty"`
}

type LoadBalancerFrontendRequest struct {
	Name          string `json:"name"`
	Algorithm     string `json:"algorithm"`
	Proto         string `json:"proto"`
	Port          string `json:"port"`
	Cookie        string `json:"cookie"`
	RedirectHTTPS string `json:"redirecthttps"`
	CertificateID string `json:"certificate_id"`
	CookieName    string `json:"cookiename,omitempty"`
}

type LoadBalancerBackendRequest struct {
	FrontendID  string `json:"frontend_id"`
	BackendPort string `json:"backend_port"`
	Weight      string `json:"weight"`
	Type        string `json:"type"`
	CloudID     string `json:"cloudid,omitempty"`
	IP          string `json:"ip,omitempty"`
}

type LoadBalancerACLRequest struct {
	FrontendID    string `json:"frontend_id"`
	Name          string `json:"name"`
	ConditionType string `json:"conditionType"`
	Value         string `json:"value"`
}

type LoadBalancerSettingsRequest struct {
	TimeoutConnect       string `json:"timeout_connect"`
	TimeoutClient        string `json:"timeout_client"`
	TimeoutServer        string `json:"timeout_server"`
	TimeoutHTTPRequest   string `json:"timeout_http_request"`
	TimeoutHTTPKeepalive string `json:"timeout_http_keepalive"`
	TimeoutTunnel        string `json:"timeout_tunnel"`
	MaxConnections       string `json:"max_connections"`
	HTTP2                string `json:"http2"`
	Compression          string `json:"compression"`
	HSTS                 string `json:"hsts"`
}

// ── Response structs ──────────────────────────────────────────────────────

type LoadBalancerInstance struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	IP           string `json:"ip"`
	DNS          string `json:"dns"`
	DCSlug       string `json:"dcslug"`
	Status       string `json:"status"`
	BackendCount string `json:"backendcount"`
	CreatedAt    string `json:"created_at"`
}

type LoadBalancerListResponse struct {
	LoadBalancers []LoadBalancerInstance `json:"loadbalancers"`
}

// ── API methods ───────────────────────────────────────────────────────────

func (c *Client) CreateLoadBalancer(req *LoadBalancerCreateRequest) (string, error) {
	respBytes, err := c.Post("/loadbalancer", req)
	if err != nil {
		return "", fmt.Errorf("failed to create load balancer: %w", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", fmt.Errorf("failed to parse create load balancer response: %w", err)
	}
	if result["status"] != "success" {
		return "", fmt.Errorf("create load balancer failed: %s", result["message"])
	}
	id := fmt.Sprintf("%v", result["loadbalancerid"])
	return id, nil
}

func (c *Client) GetLoadBalancer(lbID string) (*LoadBalancerInstance, error) {
	respBytes, err := c.Get(fmt.Sprintf("/loadbalancer/%s", lbID))
	if err != nil {
		return nil, fmt.Errorf("failed to get load balancer: %w", err)
	}
	var resp LoadBalancerListResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse get load balancer response: %w", err)
	}
	if len(resp.LoadBalancers) == 0 {
		return nil, nil
	}
	return &resp.LoadBalancers[0], nil
}

func (c *Client) DeleteLoadBalancer(lbID string) error {
	respBytes, err := c.Delete(fmt.Sprintf("/loadbalancer/%s", lbID))
	if err != nil {
		return fmt.Errorf("failed to delete load balancer: %w", err)
	}
	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse delete load balancer response: %w", err)
	}
	if result["status"] != "success" {
		return fmt.Errorf("delete load balancer failed: %s", result["message"])
	}
	return nil
}

func (c *Client) AddLBFrontend(lbID string, req *LoadBalancerFrontendRequest) (string, error) {
	respBytes, err := c.Post(fmt.Sprintf("/loadbalancer/%s/frontend", lbID), req)
	if err != nil {
		return "", fmt.Errorf("failed to add frontend: %w", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", fmt.Errorf("failed to parse add frontend response: %w", err)
	}
	if result["status"] != "success" {
		return "", fmt.Errorf("add frontend failed: %s", result["message"])
	}
	id := fmt.Sprintf("%v", result["id"])
	return id, nil
}

func (c *Client) DeleteLBFrontend(lbID string, frontendID string) error {
	endpoint := fmt.Sprintf("/loadbalancer/%s/frontend/%s", lbID, frontendID)

	// Retry up to 5 times to handle "pending action" errors
	for attempt := 0; attempt < 5; attempt++ {
		respBytes, err := c.Delete(endpoint)
		if err != nil {
			// 404 = already deleted
			if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") {
				return nil
			}
			return fmt.Errorf("failed to delete frontend: %w", err)
		}

		trimmed := strings.TrimSpace(string(respBytes))
		if trimmed == "" || trimmed == "null" {
			return nil
		}

		var result map[string]string
		if err := json.Unmarshal(respBytes, &result); err != nil {
			return nil
		}

		if result["status"] == "success" {
			return nil
		}

		msg := result["message"]
		// Not found = already deleted
		if strings.Contains(msg, "not found") || strings.Contains(msg, "Not Found") {
			return nil
		}
		// Pending action = retry after delay
		if strings.Contains(msg, "pending action") || strings.Contains(msg, "in process") {
			time.Sleep(10 * time.Second)
			continue
		}

		return fmt.Errorf("delete frontend failed: %s", msg)
	}

	return fmt.Errorf("delete frontend failed: LB still processing after retries")
}

func (c *Client) AddLBBackend(lbID string, req *LoadBalancerBackendRequest) (string, error) {
	respBytes, err := c.Post(fmt.Sprintf("/loadbalancer/%s/backend", lbID), req)
	if err != nil {
		return "", fmt.Errorf("failed to add backend: %w", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", fmt.Errorf("failed to parse add backend response: %w", err)
	}
	if result["status"] != "success" {
		return "", fmt.Errorf("add backend failed: %s", result["message"])
	}
	id := fmt.Sprintf("%v", result["id"])
	return id, nil
}

func (c *Client) DeleteLBBackend(lbID string, backendID string) error {
	respBytes, err := c.Delete(fmt.Sprintf("/loadbalancer/%s/backend/%s", lbID, backendID))
	if err != nil {
		return fmt.Errorf("failed to delete backend: %w", err)
	}
	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse delete backend response: %w", err)
	}
	if result["status"] != "success" {
		return fmt.Errorf("delete backend failed: %s", result["message"])
	}
	return nil
}

func (c *Client) AddLBACL(lbID string, req *LoadBalancerACLRequest) (string, error) {
	respBytes, err := c.Post(fmt.Sprintf("/loadbalancer/%s/acl", lbID), req)
	if err != nil {
		return "", fmt.Errorf("failed to add ACL rule: %w", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", fmt.Errorf("failed to parse add ACL response: %w", err)
	}
	if result["status"] != "success" {
		return "", fmt.Errorf("add ACL failed: %s", result["message"])
	}
	id := fmt.Sprintf("%v", result["id"])
	return id, nil
}

func (c *Client) UpdateLBACL(lbID string, aclID string, req *LoadBalancerACLRequest) error {
	respBytes, err := c.Put(fmt.Sprintf("/loadbalancer/%s/acl/%s", lbID, aclID), req)
	if err != nil {
		return fmt.Errorf("failed to update ACL rule: %w", err)
	}
	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse update ACL response: %w", err)
	}
	if result["status"] != "success" {
		return fmt.Errorf("update ACL failed: %s", result["message"])
	}
	return nil
}

func (c *Client) DeleteLBACL(lbID string, aclID string) error {
	respBytes, err := c.Delete(fmt.Sprintf("/loadbalancer/%s/acl/%s", lbID, aclID))
	if err != nil {
		return fmt.Errorf("failed to delete ACL rule: %w", err)
	}
	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse delete ACL response: %w", err)
	}
	if result["status"] != "success" {
		return fmt.Errorf("delete ACL failed: %s", result["message"])
	}
	return nil
}

func (c *Client) UpdateLBSettings(lbID string, req *LoadBalancerSettingsRequest) error {
	respBytes, err := c.Put(fmt.Sprintf("/loadbalancer/%s/settings", lbID), req)
	if err != nil {
		return fmt.Errorf("failed to update load balancer settings: %w", err)
	}
	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse update settings response: %w", err)
	}
	if result["status"] != "success" {
		return fmt.Errorf("update settings failed: %s", result["message"])
	}
	return nil
}
