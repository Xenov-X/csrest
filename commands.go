package csclient

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

// ExecuteShell executes a shell command on the beacon
func (c *Client) ExecuteShell(ctx context.Context, bid string, command string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	path := fmt.Sprintf("/api/v1/beacons/%s/spawn/command/shell", bid)
	req := map[string]string{"command": command}
	if err := c.doRequest(ctx, "POST", path, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute shell command: %w", err)
	}
	return &resp, nil
}

// ExecutePowerShell executes a PowerShell command on the beacon using managed PowerShell
// The command should be the full PowerShell command/script to execute
func (c *Client) ExecutePowerShell(ctx context.Context, bid string, command string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	path := fmt.Sprintf("/api/v1/beacons/%s/spawn/powershell", bid)
	req := PowerShellDto{
		Commandlet: command,
		Arguments:  "",
	}
	if err := c.doRequest(ctx, "POST", path, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute powershell command: %w", err)
	}
	return &resp, nil
}

// Upload uploads a file to the beacon's current working directory
func (c *Client) Upload(ctx context.Context, bid string, localPath string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	path := fmt.Sprintf("/api/v1/beacons/%s/execute/upload", bid)

	// Read file and base64 encode
	fileData, err := readAndEncodeFile(localPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Extract filename from local path
	filename := filepath.Base(localPath)

	req := UploadDto{
		File:  "@files/" + filename,  // Reference to files map
		Files: map[string]string{filename: fileData},
	}

	if err := c.doRequest(ctx, "POST", path, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to upload file: %w", err)
	}
	return &resp, nil
}

// Download downloads a file from the beacon
func (c *Client) Download(ctx context.Context, bid string, remotePath string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	path := fmt.Sprintf("/api/v1/beacons/%s/execute/download", bid)
	req := map[string]string{"path": remotePath}
	if err := c.doRequest(ctx, "POST", path, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to download file: %w", err)
	}
	return &resp, nil
}

// Screenshot captures a screenshot from the beacon by injecting into a process
// pid: Process ID to inject into (use 0 for automatic selection)
// arch: Architecture ("x86" or "x64")
func (c *Client) Screenshot(ctx context.Context, bid string, pid int, arch string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	path := fmt.Sprintf("/api/v1/beacons/%s/inject/screenshot", bid)
	req := map[string]interface{}{
		"pid":  pid,
		"arch": arch,
	}
	if err := c.doRequest(ctx, "POST", path, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to capture screenshot: %w", err)
	}
	return &resp, nil
}

// ScreenshotSpawn captures a screenshot by spawning a new process
func (c *Client) ScreenshotSpawn(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	path := fmt.Sprintf("/api/v1/beacons/%s/spawn/screenshot", bid)
	if err := c.doRequest(ctx, "POST", path, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to capture screenshot: %w", err)
	}
	return &resp, nil
}

// readAndEncodeFile reads a file and returns its base64 encoded content
func readAndEncodeFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}
