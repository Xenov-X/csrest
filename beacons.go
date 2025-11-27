package csclient

import (
	"context"
	"fmt"
)

// ListBeacons retrieves all beacons
func (c *Client) ListBeacons(ctx context.Context) ([]BeaconDto, error) {
	var beacons []BeaconDto
	if err := c.doRequest(ctx, "GET", "/api/v1/beacons", nil, &beacons, true); err != nil {
		return nil, fmt.Errorf("failed to list beacons: %w", err)
	}
	return beacons, nil
}

// GetBeacon retrieves information about a specific beacon
func (c *Client) GetBeacon(ctx context.Context, bid string) (*BeaconDto, error) {
	var beacon BeaconDto
	if err := c.doRequest(ctx, "GET", fmt.Sprintf("/api/v1/beacons/%s", bid), nil, &beacon, true); err != nil {
		return nil, fmt.Errorf("failed to get beacon: %w", err)
	}
	return &beacon, nil
}

// ExecuteBOFString executes a BOF with string arguments
func (c *Client) ExecuteBOFString(ctx context.Context, bid string, req InlineExecuteStringDto) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	path := fmt.Sprintf("/api/v1/beacons/%s/execute/bof/string", bid)
	if err := c.doRequest(ctx, "POST", path, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute BOF: %w", err)
	}
	return &resp, nil
}

// ExecuteBOFPacked executes a BOF with packed arguments
func (c *Client) ExecuteBOFPacked(ctx context.Context, bid string, req InlineExecutePackedDto) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	path := fmt.Sprintf("/api/v1/beacons/%s/execute/bof/packed", bid)
	if err := c.doRequest(ctx, "POST", path, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute BOF: %w", err)
	}
	return &resp, nil
}

// ExecuteBOFPack executes a BOF with typed arguments
func (c *Client) ExecuteBOFPack(ctx context.Context, bid string, req InlineExecutePackDto) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	path := fmt.Sprintf("/api/v1/beacons/%s/execute/bof/pack", bid)
	if err := c.doRequest(ctx, "POST", path, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute BOF: %w", err)
	}
	return &resp, nil
}

// GetUID executes the getuid command (whoami equivalent)
func (c *Client) GetUID(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	path := fmt.Sprintf("/api/v1/beacons/%s/execute/getUid", bid)
	if err := c.doRequest(ctx, "POST", path, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute getuid: %w", err)
	}
	return &resp, nil
}

// GetSystem attempts to elevate to SYSTEM
func (c *Client) GetSystem(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	path := fmt.Sprintf("/api/v1/beacons/%s/execute/getSystem", bid)
	if err := c.doRequest(ctx, "POST", path, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute getsystem: %w", err)
	}
	return &resp, nil
}
