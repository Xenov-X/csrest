package csclient

import (
	"context"
	"fmt"
)

// --- DTO Types ---

// NetViewDto represents a net view request to list domain hosts
type NetViewDto struct {
	Domain string `json:"domain,omitempty"`
}

// NetUserDto represents a net user request to list users
type NetUserDto struct {
	Target string `json:"target,omitempty"`
}

// NetUserDetailDto represents a net user detail request to get information about a specific user
type NetUserDetailDto struct {
	Target string `json:"target,omitempty"`
	User   string `json:"user,omitempty"`
}

// NetTimeDto represents a net time request to show time for a target
type NetTimeDto struct {
	Target string `json:"target,omitempty"`
}

// NetShareDto represents a net share request to list shares on a target
type NetShareDto struct {
	Target string `json:"target,omitempty"`
}

// NetSessionsDto represents a net sessions request to list sessions on a target
type NetSessionsDto struct {
	Target string `json:"target,omitempty"`
}

// NetLogonsDto represents a net logons request to list logged in users on a target
type NetLogonsDto struct {
	Target string `json:"target,omitempty"`
}

// NetLocalGroupDto represents a net localgroup request to enumerate local groups
type NetLocalGroupDto struct {
	Target    string `json:"target,omitempty"`
	GroupName string `json:"groupName,omitempty"`
}

// NetGroupDto represents a net group request to enumerate domain groups
type NetGroupDto struct {
	Target    string `json:"target,omitempty"`
	GroupName string `json:"groupName,omitempty"`
}

// NetDomainTrustsDto represents a net domain trusts request
type NetDomainTrustsDto struct {
	Domain string `json:"domain,omitempty"`
}

// NetDomainControllersDto represents a net domain controllers request
type NetDomainControllersDto struct {
	Domain string `json:"domain,omitempty"`
}

// NetDcListDto represents a net dclist request to list domain controllers
type NetDcListDto struct {
	Domain string `json:"domain,omitempty"`
}

// NetComputersDto represents a net computers request to list domain computers
type NetComputersDto struct {
	Domain string `json:"domain,omitempty"`
}

// PortScanDto represents a portscan request
type PortScanDto struct {
	Targets        []string `json:"targets"`
	Ports          []string `json:"ports"`
	Method         string   `json:"method,omitempty"`
	MaxConnections int      `json:"maxConnections,omitempty"`
}

// --- Client Methods ---

// NetDomain gets the current domain
func (c *Client) NetDomain(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/net/domain", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute net domain: %w", err)
	}
	return &resp, nil
}

// NetView lists domain hosts
func (c *Client) NetView(ctx context.Context, bid string, domain string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/net/view", bid)
	req := NetViewDto{Domain: domain}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute net view: %w", err)
	}
	return &resp, nil
}

// NetUser lists users on a system
func (c *Client) NetUser(ctx context.Context, bid string, target string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/net/user", bid)
	req := NetUserDto{Target: target}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute net user: %w", err)
	}
	return &resp, nil
}

// NetUserDetail gets information about a specific user
func (c *Client) NetUserDetail(ctx context.Context, bid string, target string, user string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/net/user/detail", bid)
	req := NetUserDetailDto{Target: target, User: user}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute net user detail: %w", err)
	}
	return &resp, nil
}

// NetTime shows time for a target
func (c *Client) NetTime(ctx context.Context, bid string, target string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/net/time", bid)
	req := NetTimeDto{Target: target}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute net time: %w", err)
	}
	return &resp, nil
}

// NetShare lists shares on a target
func (c *Client) NetShare(ctx context.Context, bid string, target string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/net/share", bid)
	req := NetShareDto{Target: target}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute net share: %w", err)
	}
	return &resp, nil
}

// NetSessions lists sessions on a target
func (c *Client) NetSessions(ctx context.Context, bid string, target string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/net/sessions", bid)
	req := NetSessionsDto{Target: target}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute net sessions: %w", err)
	}
	return &resp, nil
}

// NetLogons lists logged in users on a target
func (c *Client) NetLogons(ctx context.Context, bid string, target string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/net/logons", bid)
	req := NetLogonsDto{Target: target}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute net logons: %w", err)
	}
	return &resp, nil
}

// NetLocalGroup enumerates local groups on a specific system
func (c *Client) NetLocalGroup(ctx context.Context, bid string, target string, groupName string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/net/localGroup", bid)
	req := NetLocalGroupDto{Target: target, GroupName: groupName}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute net localgroup: %w", err)
	}
	return &resp, nil
}

// NetGroup enumerates groups on a domain controller
func (c *Client) NetGroup(ctx context.Context, bid string, target string, groupName string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/net/group", bid)
	req := NetGroupDto{Target: target, GroupName: groupName}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute net group: %w", err)
	}
	return &resp, nil
}

// NetDomainTrusts lists domain trusts for the specified domain
func (c *Client) NetDomainTrusts(ctx context.Context, bid string, domain string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/net/domainTrusts", bid)
	req := NetDomainTrustsDto{Domain: domain}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute net domain trusts: %w", err)
	}
	return &resp, nil
}

// NetDomainControllers lists hosts from the Domain Controllers group on the specified domain
func (c *Client) NetDomainControllers(ctx context.Context, bid string, domain string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/net/domainControllers", bid)
	req := NetDomainControllersDto{Domain: domain}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute net domain controllers: %w", err)
	}
	return &resp, nil
}

// NetDcList lists domain controllers for the specified domain
func (c *Client) NetDcList(ctx context.Context, bid string, domain string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/net/dclist", bid)
	req := NetDcListDto{Domain: domain}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute net dclist: %w", err)
	}
	return &resp, nil
}

// NetComputers lists hosts from the Domain Computers and Domain Controllers groups on the specified domain
func (c *Client) NetComputers(ctx context.Context, bid string, domain string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/net/computers", bid)
	req := NetComputersDto{Domain: domain}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute net computers: %w", err)
	}
	return &resp, nil
}

// PortScan runs a portscan against the specified hosts
func (c *Client) PortScan(ctx context.Context, bid string, targets []string, ports []string, method string, maxConnections int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/portscan", bid)
	req := PortScanDto{
		Targets:        targets,
		Ports:          ports,
		Method:         method,
		MaxConnections: maxConnections,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute portscan: %w", err)
	}
	return &resp, nil
}
