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

// --- File & Directory Operations ---

// Cd changes the current working directory on the beacon
func (c *Client) Cd(ctx context.Context, bid string, path string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/cd", bid)
	req := CdDto{Path: path}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute cd: %w", err)
	}
	return &resp, nil
}

// Ls lists directory contents on the beacon
func (c *Client) Ls(ctx context.Context, bid string, path string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/ls", bid)
	req := LsDto{Path: path}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute ls: %w", err)
	}
	return &resp, nil
}

// Pwd gets the current working directory on the beacon
func (c *Client) Pwd(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/pwd", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute pwd: %w", err)
	}
	return &resp, nil
}

// Mkdir creates a directory on the beacon
func (c *Client) Mkdir(ctx context.Context, bid string, folder string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/mkdir", bid)
	req := MkdirDto{Folder: folder}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute mkdir: %w", err)
	}
	return &resp, nil
}

// Cp copies a file on the beacon
func (c *Client) Cp(ctx context.Context, bid string, src, dst string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/cp", bid)
	req := CpDto{Src: src, Dst: dst}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute cp: %w", err)
	}
	return &resp, nil
}

// Mv moves/renames a file on the beacon
func (c *Client) Mv(ctx context.Context, bid string, source, destination string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/mv", bid)
	req := MoveDto{Source: source, Destination: destination}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute mv: %w", err)
	}
	return &resp, nil
}

// Rm removes a file or folder on the beacon
func (c *Client) Rm(ctx context.Context, bid string, path string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/rm", bid)
	req := RmDto{Path: path}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute rm: %w", err)
	}
	return &resp, nil
}

// Drives lists drives on the beacon
func (c *Client) Drives(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/drives", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute drives: %w", err)
	}
	return &resp, nil
}

// Timestomp copies file timestamps from source to destination on the beacon
func (c *Client) Timestomp(ctx context.Context, bid string, source, destination string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/timestomp", bid)
	req := TimeStompDto{Source: source, Destination: destination}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute timestomp: %w", err)
	}
	return &resp, nil
}

// --- Process Management Operations ---

// Ps lists processes on the beacon
func (c *Client) Ps(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/ps", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute ps: %w", err)
	}
	return &resp, nil
}

// Kill terminates a process by PID on the beacon
func (c *Client) Kill(ctx context.Context, bid string, pid int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/killProcess", bid)
	req := KillDto{PID: pid}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute kill: %w", err)
	}
	return &resp, nil
}

// GetPrivs enables all available privileges on the beacon
func (c *Client) GetPrivs(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/getPrivs", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute getprivs: %w", err)
	}
	return &resp, nil
}

// SetEnv sets an environment variable on the beacon
func (c *Client) SetEnv(ctx context.Context, bid string, key, value string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/setenv", bid)
	req := SetEnvDto{Key: key, Value: value}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute setenv: %w", err)
	}
	return &resp, nil
}

// Exit tells the beacon to exit
func (c *Client) Exit(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/exit", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute exit: %w", err)
	}
	return &resp, nil
}

// JobStop stops an active job on the beacon
func (c *Client) JobStop(ctx context.Context, bid string, jid int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/jobStop", bid)
	req := JobKillDto{JID: jid}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute job stop: %w", err)
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
