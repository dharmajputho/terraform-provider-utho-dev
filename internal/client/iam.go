package client

import (
	"encoding/json"
	"fmt"
)

// ── Request structs ───────────────────────────────────────────────────────

type IAMUserCreateRequest struct {
	FullName    string `json:"fullname"`
	Email       string `json:"email"`
	MobileCC    string `json:"mobilecc"`
	Mobile      string `json:"mobile"`
	Permissions string `json:"permissions"`
	Resources   string `json:"resources"`
	Status      string `json:"status"`
}

type IAMUserUpdateRequest struct {
	Permissions string `json:"permissions"`
	Resources   string `json:"resources"`
	Status      string `json:"status,omitempty"`
}

// ── Response structs ──────────────────────────────────────────────────────

type IAMUserInstance struct {
	ID          string      `json:"id"`
	SubUser     interface{} `json:"subuser"`
	FullName    string      `json:"fullname"`
	Email       string      `json:"email"`
	Permissions string      `json:"permissions"`
	Status      string      `json:"status"`
	Resources   string      `json:"resources"`
	DateAdded   string      `json:"dateadded"`
}

type IAMUserListResponse struct {
	Status        string            `json:"status"`
	AccountAccess []IAMUserInstance `json:"accountaccess"`
}

type IAMUserDetailResponse struct {
	Status string          `json:"status"`
	Info   IAMUserInstance `json:"info"`
}

// ── API methods ───────────────────────────────────────────────────────────

func (c *Client) CreateIAMUser(req *IAMUserCreateRequest) (string, error) {
	respBytes, err := c.Post("/user", req)
	if err != nil {
		return "", fmt.Errorf("failed to create IAM user: %w", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", fmt.Errorf("failed to parse create IAM user response: %w", err)
	}
	if result["status"] != "success" {
		return "", fmt.Errorf("create IAM user failed: %s", result["message"])
	}
	// API returns subuser ID — we need the account access ID from the list
	var subuserID string
	switch v := result["id"].(type) {
	case float64:
		subuserID = fmt.Sprintf("%.0f", v)
	default:
		subuserID = fmt.Sprintf("%v", v)
	}
	// Find account access ID by matching subuser ID
	users, err := c.ListIAMUsers()
	if err == nil {
		for _, u := range users {
			subID := fmt.Sprintf("%.0f", func() float64 {
				switch v := u.SubUser.(type) {
				case float64:
					return v
				}
				return 0
			}())
			if subID == subuserID {
				return u.ID, nil
			}
		}
	}
	// Fallback to subuser ID if list fails
	return subuserID, nil
}

func (c *Client) GetIAMUser(userID string) (*IAMUserInstance, error) {
	users, err := c.ListIAMUsers()
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		subID := fmt.Sprintf("%.0f", func() float64 {
			switch v := u.SubUser.(type) {
			case float64:
				return v
			}
			return 0
		}())
		if u.ID == userID || subID == userID {
			return &u, nil
		}
	}
	return nil, nil
}

func (c *Client) ListIAMUsers() ([]IAMUserInstance, error) {
	respBytes, err := c.Get("/user")
	if err != nil {
		return nil, fmt.Errorf("failed to list IAM users: %w", err)
	}
	var resp IAMUserListResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse list IAM users response: %w", err)
	}
	return resp.AccountAccess, nil
}

func (c *Client) UpdateIAMUser(userID string, req *IAMUserUpdateRequest) error {
	respBytes, err := c.Put(fmt.Sprintf("/user/%s", userID), req)
	if err != nil {
		return fmt.Errorf("failed to update IAM user: %w", err)
	}
	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse update IAM user response: %w", err)
	}
	if result["status"] != "success" {
		return fmt.Errorf("update IAM user failed: %s", result["message"])
	}
	return nil
}

func (c *Client) DeleteIAMUser(userID string) error {
	respBytes, err := c.Delete(fmt.Sprintf("/user/%s", userID))
	if err != nil {
		return fmt.Errorf("failed to delete IAM user: %w", err)
	}
	var result map[string]string
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("failed to parse delete IAM user response: %w", err)
	}
	if result["status"] != "success" {
		return fmt.Errorf("delete IAM user failed: %s", result["message"])
	}
	return nil
}
