package csclient

import (
	"context"
	"fmt"
)

// ============================================================
// Payload DTOs
// ============================================================

// ArtifactDto represents a server-side artifact (payload file)
type ArtifactDto struct {
	Key                string `json:"key"`
	SymbolicReference string `json:"symbolicReference"`
}

// PayloadGenerateDto represents a stageless payload generation request
// Use map[string]interface{} for flexible configuration
// Required fields:
// - listenerName (string): The listener name
// - useListenerGuardRails (bool): Use listener's guard rails or custom
// - architecture (string): "x86" or "x64"
// - exitFunction (string): "Process" or "Thread"
// - systemCallMethod (string): "None", "Direct", or "Indirect"
// - httpLibrary (string): "winhttp", "wininet", or "" (for HTTP/HTTPS beacons)
// - format (string): payload format (e.g., "exe", "dll", "shellcode", "powershell", etc.)
// - fileName (string): output filename for the artifact
// Optional fields:
// - guardRails (object): custom guard rails if useListenerGuardRails is false
// - proxyCredentials (object): proxy credentials configuration
// - amsiBypass (bool): enable AMSI bypass
// - etw (string): ETW patching method
// - obfuscateAllocationRoutines (bool): obfuscate memory allocation routines
// - etc. (see OpenAPI spec for full list)

// PayloadStagerDto represents a stager payload generation request
// Use map[string]interface{} for flexible configuration
// Required fields:
// - listenerName (string): The listener name
// - architecture (string): "x86" or "x64"
// - format (string): payload format
// - fileName (string): output filename

// HashesDto contains file hashes for a generated payload
type HashesDto struct {
	MD5    string `json:"md5,omitempty"`
	SHA1   string `json:"sha1,omitempty"`
	SHA256 string `json:"sha256,omitempty"`
}

// PayloadResultDto represents the result of stageless payload generation
type PayloadResultDto struct {
	Status              string                 `json:"status,omitempty"`
	Notes               string                 `json:"notes,omitempty"`
	InformationFileName string                 `json:"informationFileName,omitempty"`
	PayloadFileName     string                 `json:"payloadFileName,omitempty"`
	Size                int                    `json:"size,omitempty"`
	Created             string                 `json:"created,omitempty"`
	Hashes              *HashesDto             `json:"hashes,omitempty"`
	Inputs              map[string]interface{} `json:"inputs,omitempty"`
	Error               string                 `json:"error,omitempty"`
}

// PayloadStagerResultDto represents the result of stager payload generation
type PayloadStagerResultDto struct {
	Status              string                 `json:"status,omitempty"`
	Notes               string                 `json:"notes,omitempty"`
	InformationFileName string                 `json:"informationFileName,omitempty"`
	PayloadFileName     string                 `json:"payloadFileName,omitempty"`
	Size                int                    `json:"size,omitempty"`
	Created             string                 `json:"created,omitempty"`
	Hashes              *HashesDto             `json:"hashes,omitempty"`
	Inputs              map[string]interface{} `json:"inputs,omitempty"`
	Error               string                 `json:"error,omitempty"`
}

// ============================================================
// Payload Generation Methods
// ============================================================

// ListArtifacts retrieves all server-side artifacts (generated payloads)
func (c *Client) ListArtifacts(ctx context.Context) ([]ArtifactDto, error) {
	var artifacts []ArtifactDto
	if err := c.doRequest(ctx, "GET", "/api/v1/artifacts", nil, &artifacts, true); err != nil {
		return nil, fmt.Errorf("failed to list artifacts: %w", err)
	}
	return artifacts, nil
}

// GenerateStagelessPayload generates a stageless beacon payload
// config should contain PayloadDto fields as per the OpenAPI spec.
// Example:
//
//	config := map[string]interface{}{
//	    "listenerName": "http",
//	    "useListenerGuardRails": true,
//	    "architecture": "x64",
//	    "exitFunction": "Process",
//	    "systemCallMethod": "None",
//	    "httpLibrary": "winhttp",
//	    "format": "exe",
//	    "fileName": "beacon.exe",
//	}
func (c *Client) GenerateStagelessPayload(ctx context.Context, config map[string]interface{}) (*PayloadResultDto, error) {
	var resp PayloadResultDto
	if err := c.doRequest(ctx, "POST", "/api/v1/payloads/generate/stageless", config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to generate stageless payload: %w", err)
	}
	return &resp, nil
}

// GenerateStagerPayload generates a stager beacon payload
// config should contain PayloadStagerDto fields as per the OpenAPI spec.
// Example:
//
//	config := map[string]interface{}{
//	    "listenerName": "http",
//	    "architecture": "x64",
//	    "format": "exe",
//	    "fileName": "stager.exe",
//	}
func (c *Client) GenerateStagerPayload(ctx context.Context, config map[string]interface{}) (*PayloadStagerResultDto, error) {
	var resp PayloadStagerResultDto
	if err := c.doRequest(ctx, "POST", "/api/v1/payloads/generate/stager", config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to generate stager payload: %w", err)
	}
	return &resp, nil
}

// DownloadPayload downloads a previously generated payload from the server
// Returns the raw payload bytes
func (c *Client) DownloadPayload(ctx context.Context, fileName string) ([]byte, error) {
	var payloadBytes []byte
	endpoint := fmt.Sprintf("/api/v1/payloads/%s", fileName)

	// Note: The API returns binary data (byte format), so we use doRequest with []byte
	if err := c.doRequest(ctx, "GET", endpoint, nil, &payloadBytes, true); err != nil {
		return nil, fmt.Errorf("failed to download payload: %w", err)
	}
	return payloadBytes, nil
}

// ============================================================
// Server Configuration Methods
// ============================================================

// ApiRoot retrieves the API root information (entry point)
func (c *Client) ApiRoot(ctx context.Context) (map[string]interface{}, error) {
	var resp map[string]interface{}
	if err := c.doRequest(ctx, "GET", "/api/v1", nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to get api root: %w", err)
	}
	return resp, nil
}

// GetKillDate retrieves the beacon kill date configured on the team server.
// The API returns a plain text string.
func (c *Client) GetKillDate(ctx context.Context) (string, error) {
	var resp string
	if err := c.doRequest(ctx, "GET", "/api/v1/config/killdate", nil, &resp, true); err != nil {
		return "", fmt.Errorf("failed to get kill date: %w", err)
	}
	return resp, nil
}

// GetC2Profile retrieves the Malleable C2 profile from the team server.
// The API returns a plain text string.
func (c *Client) GetC2Profile(ctx context.Context) (string, error) {
	var resp string
	if err := c.doRequest(ctx, "GET", "/api/v1/config/profile", nil, &resp, true); err != nil {
		return "", fmt.Errorf("failed to get c2 profile: %w", err)
	}
	return resp, nil
}

// GetSystemInformation retrieves system information from the team server.
// The API returns a plain text string.
func (c *Client) GetSystemInformation(ctx context.Context) (string, error) {
	var resp string
	if err := c.doRequest(ctx, "GET", "/api/v1/config/systeminformation", nil, &resp, true); err != nil {
		return "", fmt.Errorf("failed to get system information: %w", err)
	}
	return resp, nil
}

// GetTeamserverIp retrieves the team server IP address.
// The API returns a plain text string.
func (c *Client) GetTeamserverIp(ctx context.Context) (string, error) {
	var resp string
	if err := c.doRequest(ctx, "GET", "/api/v1/config/teamserverIp", nil, &resp, true); err != nil {
		return "", fmt.Errorf("failed to get teamserver ip: %w", err)
	}
	return resp, nil
}

// ResetData resets the data model on the team server (WARNING: destructive operation)
// This removes all beacons, tasks, and other operational data
func (c *Client) ResetData(ctx context.Context) error {
	if err := c.doRequest(ctx, "DELETE", "/api/v1/config/resetData", nil, nil, true); err != nil {
		return fmt.Errorf("failed to reset data: %w", err)
	}
	return nil
}
