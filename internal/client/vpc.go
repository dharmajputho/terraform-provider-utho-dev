package client

import (
	"encoding/json"
	"fmt"
)

// ── Request structs ───────────────────────────────────────────────────────

type VPCCreateRequest struct {
	Name    string `json:"name"`
	Network string `json:"network"`
	Size    string `json:"size"`
	DCSlug  string `json:"dcslug"`
	PlanID  string `json:"planid"`
}

type SubnetCreateRequest struct {
	Name           string `json:"name"`
	VPCID          string `json:"vpcid"`
	AssignPublicIP int    `json:"assign_publicip"`
	IsDefault      int    `json:"is_default"`
	Type           string `json:"type"`
	Network        string `json:"network"`
	Size           int    `json:"size"`
}

type NATGatewayCreateRequest struct {
	Name     string `json:"name"`
	Subnet   string `json:"subnet"`
	PublicIP string `json:"publicip"`
	DCSlug   string `json:"dcslug"`
}

type RouteTableCreateRequest struct {
	Name   string `json:"name"`
	VPC    string `json:"vpc"`
	DCSlug string `json:"dcslug"`
}

type RouteCreateRequest struct {
	DestinationCIDR string `json:"destination_cidr_block"`
	RouteType       string `json:"route_type"`
	Target          string `json:"target"`
	RouteTableID    string `json:"route_table_id"`
}

type RouteUpdateRequest struct {
	DestinationCIDR string `json:"destination_cidr_block"`
	RouteType       string `json:"route_type"`
	Target          string `json:"target"`
	RouteTableID    string `json:"route_table_id"`
	RouteID         string `json:"route_id"`
}

type ElasticIPAllocateRequest struct {
	DCSlug       string `json:"dcslug"`
	BillingCycle string `json:"billingcycle"`
}

type VPCPeeringCreateRequest struct {
	Name           string `json:"name"`
	DCSlug         string `json:"dcslug"`
	RequesterVPCID string `json:"requester_vpc_id"`
	AccepterVPCID  string `json:"accepter_vpc_id"`
}

// ── Response structs ──────────────────────────────────────────────────────

type VPCInstance struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Network   string `json:"network"`
	Size      string `json:"size"`
	DCSlug    string `json:"dcslug"`
	IsDefault string `json:"is_default"`
	Total     int    `json:"total"`
	Available int    `json:"available"`
}

type VPCListResponse struct {
	Status string        `json:"status"`
	VPC    []VPCInstance `json:"vpc"`
}

type SubnetInstance struct {
	ID             string `json:"id"`
	VPCID          string `json:"vpcid"`
	Name           string `json:"name"`
	Network        string `json:"network"`
	Size           string `json:"size"`
	DCSlug         string `json:"dcslug"`
	Type           string `json:"type"`
	SubnetType     string `json:"subnet_type"`
	Gateway        string `json:"gateway"`
	AssignPublicIP string `json:"assign_publicip"`
	IsDefault      string `json:"is_default"`
	Status         string `json:"status"`
}

type SubnetListResponse struct {
	Status string           `json:"status"`
	Subnet []SubnetInstance `json:"subnet"`
}

type NATGatewayInstance struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Subnet   string `json:"subnet"`
	PublicIP string `json:"publicip"`
	DCSlug   string `json:"dcslug"`
	Status   string `json:"status"`
}

type RouteTableInstance struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	VPCID  string `json:"vpc_id"`
	IsMain string `json:"is_main"`
	Status string `json:"status"`
}

type RouteTableListResponse struct {
	Status string               `json:"status"`
	Data   []RouteTableInstance `json:"data"`
}

type RouteInstance struct {
	RouteID         string `json:"route_id"`
	DestinationCIDR string `json:"destination_cidr_block"`
	RouteType       string `json:"route_type"`
	Target          string `json:"target"`
	RouteTableID    string `json:"route_table_id"`
}

type ElasticIPInstance struct {
	ID           string `json:"id"`
	IP           string `json:"ip"`
	IPID         int    `json:"ipid"`
	DCSlug       string `json:"dcslug"`
	BillingCycle string `json:"billingcycle"`
}

type ElasticIPListResponse struct {
	Status string              `json:"status"`
	IPs    []ElasticIPInstance `json:"ips"`
}

type VPCPeeringInstance struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	RequesterVPCID string `json:"requester_vpc_id"`
	AccepterVPCID  string `json:"accepter_vpc_id"`
	Status         string `json:"status"`
}

type VPCPeeringListResponse struct {
	Status string               `json:"status"`
	Data   []VPCPeeringInstance `json:"data"`
}

type GenericResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	ID      string `json:"id"`
}

// ── VPC methods ───────────────────────────────────────────────────────────

func (c *Client) CreateVPC(req *VPCCreateRequest) (string, error) {
	respBytes, err := c.Post("/vpc/create", req)
	if err != nil {
		return "", fmt.Errorf("failed to create VPC: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return "", fmt.Errorf("failed to parse create VPC response: %w", err)
	}
	if resp.Status != "success" {
		return "", fmt.Errorf("create VPC failed: %s", resp.Message)
	}
	return resp.ID, nil
}

func (c *Client) GetVPC(vpcID string) (*VPCInstance, error) {
	respBytes, err := c.Get(fmt.Sprintf("/vpc/%s", vpcID))
	if err != nil {
		return nil, fmt.Errorf("failed to get VPC: %w", err)
	}
	var vpc VPCInstance
	if err := json.Unmarshal(respBytes, &vpc); err != nil {
		return nil, fmt.Errorf("failed to parse get VPC response: %w", err)
	}
	if vpc.ID == "" {
		return nil, nil
	}
	return &vpc, nil
}

func (c *Client) DeleteVPC(vpcID string) error {
	respBytes, err := c.Delete(fmt.Sprintf("/vpc/%s", vpcID))
	if err != nil {
		return fmt.Errorf("failed to delete VPC: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return fmt.Errorf("failed to parse delete VPC response: %w", err)
	}
	if resp.Status != "success" {
		return fmt.Errorf("delete VPC failed: %s", resp.Message)
	}
	return nil
}

// ── Subnet methods ────────────────────────────────────────────────────────

func (c *Client) CreateSubnet(req *SubnetCreateRequest) (string, error) {
	respBytes, err := c.Post("/vpc/subnet/create", req)
	if err != nil {
		return "", fmt.Errorf("failed to create subnet: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return "", fmt.Errorf("failed to parse create subnet response: %w", err)
	}
	if resp.Status != "success" {
		return "", fmt.Errorf("create subnet failed: %s", resp.Message)
	}
	return resp.ID, nil
}

func (c *Client) GetSubnet(subnetID string) (*SubnetInstance, error) {
	respBytes, err := c.Get(fmt.Sprintf("/vpc/subnet/%s", subnetID))
	if err != nil {
		return nil, fmt.Errorf("failed to get subnet: %w", err)
	}
	var subnet SubnetInstance
	if err := json.Unmarshal(respBytes, &subnet); err != nil {
		return nil, fmt.Errorf("failed to parse get subnet response: %w", err)
	}
	if subnet.ID == "" {
		return nil, nil
	}
	return &subnet, nil
}

func (c *Client) DeleteSubnet(subnetID string) error {
	respBytes, err := c.Delete(fmt.Sprintf("/vpc/subnet/%s", subnetID))
	if err != nil {
		return fmt.Errorf("failed to delete subnet: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return fmt.Errorf("failed to parse delete subnet response: %w", err)
	}
	if resp.Status != "success" {
		return fmt.Errorf("delete subnet failed: %s", resp.Message)
	}
	return nil
}

// ── NAT Gateway methods ───────────────────────────────────────────────────

func (c *Client) CreateNATGateway(req *NATGatewayCreateRequest) (string, error) {
	respBytes, err := c.Post("/vpc/natgateway", req)
	if err != nil {
		return "", fmt.Errorf("failed to create NAT gateway: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return "", fmt.Errorf("failed to parse create NAT gateway response: %w", err)
	}
	if resp.Status != "success" {
		return "", fmt.Errorf("create NAT gateway failed: %s", resp.Message)
	}
	return resp.ID, nil
}

func (c *Client) DetachNATGateway(natGatewayID string, subnetID string) error {
	endpoint := fmt.Sprintf("/vpc/natgateway/%s/dettach", natGatewayID)
	payload := map[string]string{"subnet": subnetID}
	respBytes, err := c.Post(endpoint, payload)
	if err != nil {
		return fmt.Errorf("failed to detach NAT gateway: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return fmt.Errorf("failed to parse detach NAT gateway response: %w", err)
	}
	if resp.Status != "success" {
		return fmt.Errorf("detach NAT gateway failed: %s", resp.Message)
	}
	return nil
}

func (c *Client) DeleteNATGateway(natGatewayID string) error {
	respBytes, err := c.Delete(fmt.Sprintf("/vpc/natgateway/%s", natGatewayID))
	if err != nil {
		return fmt.Errorf("failed to delete NAT gateway: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return fmt.Errorf("failed to parse delete NAT gateway response: %w", err)
	}
	if resp.Status != "success" {
		return fmt.Errorf("delete NAT gateway failed: %s", resp.Message)
	}
	return nil
}

// ── Route Table methods ───────────────────────────────────────────────────

func (c *Client) CreateRouteTable(req *RouteTableCreateRequest) (string, error) {
	respBytes, err := c.Post("/vpc/routetable", req)
	if err != nil {
		return "", fmt.Errorf("failed to create route table: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return "", fmt.Errorf("failed to parse create route table response: %w", err)
	}
	if resp.Status != "success" {
		return "", fmt.Errorf("create route table failed: %s", resp.Message)
	}
	return resp.ID, nil
}

func (c *Client) AttachRouteTable(vpcID string, routeTableID string) error {
	endpoint := fmt.Sprintf("/vpc/%s/routetable/%s/attach", vpcID, routeTableID)
	respBytes, err := c.Post(endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to attach route table: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return fmt.Errorf("failed to parse attach route table response: %w", err)
	}
	if resp.Status != "success" {
		return fmt.Errorf("attach route table failed: %s", resp.Message)
	}
	return nil
}

func (c *Client) DetachRouteTable(vpcID string, routeTableID string) error {
	endpoint := fmt.Sprintf("/vpc/%s/routetable/%s/dettach", vpcID, routeTableID)
	respBytes, err := c.Post(endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to detach route table: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return fmt.Errorf("failed to parse detach route table response: %w", err)
	}
	if resp.Status != "success" {
		return fmt.Errorf("detach route table failed: %s", resp.Message)
	}
	return nil
}

func (c *Client) DeleteRouteTable(routeTableID string) error {
	respBytes, err := c.Delete(fmt.Sprintf("/vpc/routetable?id=%s", routeTableID))
	if err != nil {
		return fmt.Errorf("failed to delete route table: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return fmt.Errorf("failed to parse delete route table response: %w", err)
	}
	if resp.Status != "success" {
		return fmt.Errorf("delete route table failed: %s", resp.Message)
	}
	return nil
}

// ── Route methods ─────────────────────────────────────────────────────────

func (c *Client) CreateRoute(req *RouteCreateRequest) (string, error) {
	respBytes, err := c.Post("/vpc/route", req)
	if err != nil {
		return "", fmt.Errorf("failed to create route: %w", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", fmt.Errorf("failed to parse create route response: %w", err)
	}
	if result["status"] != "success" {
		return "", fmt.Errorf("create route failed: %s", result["message"])
	}
	data, _ := result["data"].(map[string]interface{})
	routeID, _ := data["route_id"].(string)
	return routeID, nil
}

func (c *Client) UpdateRoute(req *RouteUpdateRequest) error {
	respBytes, err := c.Put("/vpc/route", req)
	if err != nil {
		return fmt.Errorf("failed to update route: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return fmt.Errorf("failed to parse update route response: %w", err)
	}
	if resp.Status != "success" {
		return fmt.Errorf("update route failed: %s", resp.Message)
	}
	return nil
}

func (c *Client) DeleteRoute(routeID string) error {
	endpoint := fmt.Sprintf("/vpc/route?route_id=%s", routeID)
	payload := map[string]string{"route_id": routeID}
	respBytes, err := c.doRequest("DELETE", endpoint, payload)
	if err != nil {
		return fmt.Errorf("failed to delete route: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return fmt.Errorf("failed to parse delete route response: %w", err)
	}
	if resp.Status != "success" {
		return fmt.Errorf("delete route failed: %s", resp.Message)
	}
	return nil
}

// ── Elastic IP methods ────────────────────────────────────────────────────

func (c *Client) AllocateElasticIP(req *ElasticIPAllocateRequest) (*ElasticIPInstance, error) {
	respBytes, err := c.Post("/elasticip/allocate", req)
	if err != nil {
		return nil, fmt.Errorf("failed to allocate elastic IP: %w", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse allocate elastic IP response: %w", err)
	}
	if result["status"] != "success" {
		return nil, fmt.Errorf("allocate elastic IP failed: %s", result["message"])
	}
	ip, _ := result["ip"].(string)
	dcslug, _ := result["dcslug"].(string)
	billingcycle, _ := result["billingcycle"].(string)
	return &ElasticIPInstance{
		IP:           ip,
		DCSlug:       dcslug,
		BillingCycle: billingcycle,
	}, nil
}

func (c *Client) ReleaseElasticIP(ipID string) error {
	respBytes, err := c.Delete(fmt.Sprintf("/elasticip/%s", ipID))
	if err != nil {
		return fmt.Errorf("failed to release elastic IP: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return fmt.Errorf("failed to parse release elastic IP response: %w", err)
	}
	if resp.Status != "success" {
		return fmt.Errorf("release elastic IP failed: %s", resp.Message)
	}
	return nil
}

// ── VPC Peering methods ───────────────────────────────────────────────────

func (c *Client) CreateVPCPeering(req *VPCPeeringCreateRequest) (string, error) {
	respBytes, err := c.Post("/vpc/peering", req)
	if err != nil {
		return "", fmt.Errorf("failed to create VPC peering: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return "", fmt.Errorf("failed to parse create VPC peering response: %w", err)
	}
	if resp.Status != "success" {
		return "", fmt.Errorf("create VPC peering failed: %s", resp.Message)
	}
	return resp.ID, nil
}

func (c *Client) DeleteVPCPeering(peeringID string) error {
	respBytes, err := c.Delete(fmt.Sprintf("/vpc/peering?peering_id=%s", peeringID))
	if err != nil {
		return fmt.Errorf("failed to delete VPC peering: %w", err)
	}
	var resp GenericResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return fmt.Errorf("failed to parse delete VPC peering response: %w", err)
	}
	if resp.Status != "success" {
		return fmt.Errorf("delete VPC peering failed: %s", resp.Message)
	}
	return nil
}
