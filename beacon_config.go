package csclient

import (
	"context"
	"fmt"
)

// ============================================================
// Beacon Configuration DTOs
// ============================================================

// NoteDto represents a note to assign to a beacon
type NoteDto struct {
	Note string `json:"note"`
}

// SpawnToDto represents spawn-to process configuration
type SpawnToDto struct {
	Arch string `json:"arch"` // "x86" or "x64"
	Path string `json:"path"` // e.g., "%windir%\\sysnative\\rundll32.exe"
}

// PpidDto represents parent process ID configuration
type PpidDto struct {
	PID int `json:"pid"`
}

// DnsModeDto represents DNS beacon mode configuration
type DnsModeDto struct {
	Mode string `json:"mode"` // "dns", "dns6", or "dnsTxt"
}

// SyscallMethodDto represents syscall method configuration
type SyscallMethodDto struct {
	Method string `json:"method"` // syscall method name
}

// ConsoleCommandDto represents a console command to execute
type ConsoleCommandDto struct {
	Command   string            `json:"command"`
	Arguments string            `json:"arguments,omitempty"`
	Files     map[string]string `json:"files,omitempty"`
}

// SpoofedArgumentsAddDto represents spoofed arguments to add
type SpoofedArgumentsAddDto struct {
	Command       string `json:"command"`
	FakeArguments string `json:"fakeArguments"`
}

// SpoofedArgumentsRemoveDto represents spoofed arguments to remove
type SpoofedArgumentsRemoveDto struct {
	Command string `json:"command"`
}

// HostCallbackInfoDto represents a C2 host callback entry
type HostCallbackInfoDto struct {
	Hostname string `json:"hostname"`
	URI      string `json:"uri"`
}

// HostCallbackAddDto represents C2 hosts to add
type HostCallbackAddDto struct {
	Infos []HostCallbackInfoDto `json:"infos"`
}

// HostCallbackUpdateInfoDto represents a C2 host update entry
type HostCallbackUpdateInfoDto struct {
	FromHost string `json:"fromHost"`
	ToHost   string `json:"toHost"`
	ToURI    string `json:"toUri,omitempty"`
}

// HostCallbackUpdateDto represents C2 hosts to update
type HostCallbackUpdateDto struct {
	Infos []HostCallbackUpdateInfoDto `json:"infos"`
}

// HostCallbackRemoveDto represents C2 hosts to remove
type HostCallbackRemoveDto struct {
	Hostnames []string `json:"hostnames"`
}

// PublishOutputDto represents a message to publish to the beacon transcript
type PublishOutputDto struct {
	Message string `json:"message"`
}

// ============================================================
// Beacon Configuration Methods
// ============================================================

// BeaconInfo retrieves beacon metadata and configuration
func (c *Client) BeaconInfo(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/beaconInfo", bid)
	if err := c.doRequest(ctx, "POST", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to get beacon info: %w", err)
	}
	return &resp, nil
}

// DisableBeaconGate disables beacon gate for the specified beacon
func (c *Client) DisableBeaconGate(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/beaconGate/disable", bid)
	if err := c.doRequest(ctx, "POST", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to disable beacon gate: %w", err)
	}
	return &resp, nil
}

// EnableBeaconGate enables beacon gate for the specified beacon
func (c *Client) EnableBeaconGate(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/beaconGate/enable", bid)
	if err := c.doRequest(ctx, "POST", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to enable beacon gate: %w", err)
	}
	return &resp, nil
}

// DisableBlockDlls disables block DLLs for the specified beacon
func (c *Client) DisableBlockDlls(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/blockdlls/disable", bid)
	if err := c.doRequest(ctx, "POST", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to disable block dlls: %w", err)
	}
	return &resp, nil
}

// EnableBlockDlls enables block DLLs for the specified beacon
func (c *Client) EnableBlockDlls(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/blockdlls/enable", bid)
	if err := c.doRequest(ctx, "POST", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to enable block dlls: %w", err)
	}
	return &resp, nil
}

// SetSleepTime sets the beacon's sleep time and jitter
func (c *Client) SetSleepTime(ctx context.Context, bid string, sleep, jitter int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/sleepTime", bid)
	req := SleepDto{Sleep: sleep, Jitter: jitter}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to set sleep time: %w", err)
	}
	return &resp, nil
}

// SetSpawnTo sets the spawn-to process for the specified beacon
func (c *Client) SetSpawnTo(ctx context.Context, bid string, arch, path string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/spawnto", bid)
	req := SpawnToDto{Arch: arch, Path: path}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to set spawnto: %w", err)
	}
	return &resp, nil
}

// UnsetSpawnTo unsets the spawn-to process for the specified beacon
func (c *Client) UnsetSpawnTo(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/spawnto", bid)
	if err := c.doRequest(ctx, "DELETE", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to unset spawnto: %w", err)
	}
	return &resp, nil
}

// SetPpid sets the parent process ID for the specified beacon
func (c *Client) SetPpid(ctx context.Context, bid string, ppid int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/ppid", bid)
	req := PpidDto{PID: ppid}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to set ppid: %w", err)
	}
	return &resp, nil
}

// UnsetPpid unsets the parent process ID for the specified beacon
func (c *Client) UnsetPpid(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/ppid", bid)
	if err := c.doRequest(ctx, "DELETE", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to unset ppid: %w", err)
	}
	return &resp, nil
}

// SetDnsMode sets the DNS beacon mode (dns, dns6, or dnsTxt)
func (c *Client) SetDnsMode(ctx context.Context, bid string, mode string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/dnsMode", bid)
	req := DnsModeDto{Mode: mode}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to set dns mode: %w", err)
	}
	return &resp, nil
}

// GetSyscallMethod retrieves the current syscall method for the beacon
func (c *Client) GetSyscallMethod(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/syscallMethod", bid)
	if err := c.doRequest(ctx, "GET", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to get syscall method: %w", err)
	}
	return &resp, nil
}

// SetSyscallMethod sets the syscall method for the beacon
func (c *Client) SetSyscallMethod(ctx context.Context, bid string, method string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/syscallMethod", bid)
	req := SyscallMethodDto{Method: method}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to set syscall method: %w", err)
	}
	return &resp, nil
}

// SetNote assigns a note to the beacon
func (c *Client) SetNote(ctx context.Context, bid string, note string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/note", bid)
	req := NoteDto{Note: note}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to set note: %w", err)
	}
	return &resp, nil
}

// ListSpoofedArguments lists all spoofed argument configurations
func (c *Client) ListSpoofedArguments(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/spoofedArguments", bid)
	if err := c.doRequest(ctx, "GET", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to list spoofed arguments: %w", err)
	}
	return &resp, nil
}

// AddSpoofedArguments adds a spoofed argument configuration
func (c *Client) AddSpoofedArguments(ctx context.Context, bid string, command, fakeArguments string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/spoofedArguments", bid)
	req := SpoofedArgumentsAddDto{Command: command, FakeArguments: fakeArguments}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to add spoofed arguments: %w", err)
	}
	return &resp, nil
}

// RemoveSpoofedArguments removes a spoofed argument configuration
func (c *Client) RemoveSpoofedArguments(ctx context.Context, bid string, command string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/spoofedArguments", bid)
	req := SpoofedArgumentsRemoveDto{Command: command}
	if err := c.doRequest(ctx, "DELETE", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to remove spoofed arguments: %w", err)
	}
	return &resp, nil
}

// ============================================================
// C2 Host Management Methods
// ============================================================

// GetC2Host retrieves the C2 host callback information for the beacon
func (c *Client) GetC2Host(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/c2/host", bid)
	if err := c.doRequest(ctx, "GET", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to get c2 host: %w", err)
	}
	return &resp, nil
}

// AddC2Host adds C2 hosts to the beacon's callback list
func (c *Client) AddC2Host(ctx context.Context, bid string, hosts []HostCallbackInfoDto) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/c2/host", bid)
	req := HostCallbackAddDto{Infos: hosts}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to add c2 host: %w", err)
	}
	return &resp, nil
}

// UpdateC2Host updates C2 hosts in the beacon's callback list
func (c *Client) UpdateC2Host(ctx context.Context, bid string, updates []HostCallbackUpdateInfoDto) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/c2/host", bid)
	req := HostCallbackUpdateDto{Infos: updates}
	if err := c.doRequest(ctx, "PUT", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to update c2 host: %w", err)
	}
	return &resp, nil
}

// RemoveC2Host removes C2 hosts from the beacon's callback list
func (c *Client) RemoveC2Host(ctx context.Context, bid string, hostnames []string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/c2/host", bid)
	req := HostCallbackRemoveDto{Hostnames: hostnames}
	if err := c.doRequest(ctx, "DELETE", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to remove c2 host: %w", err)
	}
	return &resp, nil
}

// ============================================================
// Failover Notification Methods
// ============================================================

// GetFailoverNotification retrieves the failover notification status
func (c *Client) GetFailoverNotification(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/c2/failoverNotification", bid)
	if err := c.doRequest(ctx, "GET", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to get failover notification: %w", err)
	}
	return &resp, nil
}

// DisableFailoverNotification disables failover notification for the beacon
func (c *Client) DisableFailoverNotification(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/c2/failoverNotification/disable", bid)
	if err := c.doRequest(ctx, "POST", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to disable failover notification: %w", err)
	}
	return &resp, nil
}

// EnableFailoverNotification enables failover notification for the beacon
func (c *Client) EnableFailoverNotification(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/c2/failoverNotification/enable", bid)
	if err := c.doRequest(ctx, "POST", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to enable failover notification: %w", err)
	}
	return &resp, nil
}

// ============================================================
// Beacon Management Methods
// ============================================================

// DeleteBeacon removes a beacon from the team server
func (c *Client) DeleteBeacon(ctx context.Context, bid string) error {
	endpoint := fmt.Sprintf("/api/v1/beacons/%s", bid)
	if err := c.doRequest(ctx, "DELETE", endpoint, nil, nil, true); err != nil {
		return fmt.Errorf("failed to delete beacon: %w", err)
	}
	return nil
}

// ClearCommandQueue clears the command queue for the specified beacon
func (c *Client) ClearCommandQueue(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/clearCommandQueue", bid)
	if err := c.doRequest(ctx, "POST", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to clear command queue: %w", err)
	}
	return &resp, nil
}

// CheckIn forces a DNS beacon to check in immediately
func (c *Client) CheckIn(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/checkIn", bid)
	if err := c.doRequest(ctx, "POST", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to force check in: %w", err)
	}
	return &resp, nil
}

// ConsoleCommand executes a console command on the beacon
func (c *Client) ConsoleCommand(ctx context.Context, bid string, command, arguments string, files map[string]string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/consoleCommand", bid)
	req := ConsoleCommandDto{
		Command:   command,
		Arguments: arguments,
		Files:     files,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute console command: %w", err)
	}
	return &resp, nil
}

// ListCommandHelp retrieves help for all beacon commands
func (c *Client) ListCommandHelp(ctx context.Context, bid string) ([]map[string]interface{}, error) {
	var resp []map[string]interface{}
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/help", bid)
	if err := c.doRequest(ctx, "GET", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to list command help: %w", err)
	}
	return resp, nil
}

// GetCommandHelp retrieves help for a specific beacon command
func (c *Client) GetCommandHelp(ctx context.Context, bid string, command string) (map[string]interface{}, error) {
	var resp map[string]interface{}
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/help/%s", bid, command)
	if err := c.doRequest(ctx, "GET", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to get command help: %w", err)
	}
	return resp, nil
}

// ============================================================
// Jobs and Tasks Methods
// ============================================================

// ListJobs lists all active jobs for the specified beacon
func (c *Client) ListJobs(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/jobs", bid)
	if err := c.doRequest(ctx, "POST", endpoint, nil, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to list jobs: %w", err)
	}
	return &resp, nil
}

// TaskLog publishes a log message to the beacon/task transcript
func (c *Client) TaskLog(ctx context.Context, taskID string, message string) error {
	endpoint := fmt.Sprintf("/api/v1/tasks/%s/log", taskID)
	req := PublishOutputDto{Message: message}
	if err := c.doRequest(ctx, "POST", endpoint, req, nil, true); err != nil {
		return fmt.Errorf("failed to publish task log: %w", err)
	}
	return nil
}

// TaskError publishes an error message to the beacon/task transcript
func (c *Client) TaskError(ctx context.Context, taskID string, message string) error {
	endpoint := fmt.Sprintf("/api/v1/tasks/%s/error", taskID)
	req := PublishOutputDto{Message: message}
	if err := c.doRequest(ctx, "POST", endpoint, req, nil, true); err != nil {
		return fmt.Errorf("failed to publish task error: %w", err)
	}
	return nil
}
