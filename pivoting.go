package csclient

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
)

// --- DTO Types for Pivoting and Lateral Movement ---

// LinkDto is used for linking to SMB beacons
type LinkDto struct {
	Target string `json:"target"`
	Pipe   string `json:"pipe,omitempty"`
}

// ConnectDto is used for linking to TCP beacons
type ConnectDto struct {
	Target string `json:"target"`
	Port   int    `json:"port,omitempty"`
}

// UnlinkDto is used for disconnecting from pivot beacons
type UnlinkDto struct {
	Host string `json:"host"`
	Pid  int    `json:"pid,omitempty"`
}

// SshSpawnDto is used for SSH connection with username/password
type SshSpawnDto struct {
	Target   string `json:"target"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// SshKeySpawnDto is used for SSH connection with key file
type SshKeySpawnDto struct {
	Target   string            `json:"target"`
	Port     int               `json:"port,omitempty"`
	Username string            `json:"username"`
	Key      string            `json:"key"`
	Files    map[string]string `json:"files,omitempty"`
}

// RemoteExecDto is used for remote command execution
type RemoteExecDto struct {
	Method  string `json:"method"`
	Target  string `json:"target"`
	Command string `json:"command"`
}

// JumpDto is used for remote beacon execution (jump)
type JumpDto struct {
	Exploit  string `json:"exploit"`
	Target   string `json:"target"`
	Listener string `json:"listener"`
}

// RemoteExecInfoDto describes a remote execution method
type RemoteExecInfoDto struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// RemoteExploitInfoDto describes a remote exploit/jump method
type RemoteExploitInfoDto struct {
	Name        string `json:"name"`
	Arch        string `json:"arch"`
	Description string `json:"description,omitempty"`
}

// --- DTO Types for Tunneling ---

// Socks4StartDto is used for starting SOCKS4a server
type Socks4StartDto struct {
	Port int `json:"port"`
}

// SocksAuthDto is used for SOCKS5 authentication
type SocksAuthDto struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// Socks5StartDto is used for starting SOCKS5 server
type Socks5StartDto struct {
	Port          int           `json:"port"`
	Auth          *SocksAuthDto `json:"auth,omitempty"`
	EnableLogging bool          `json:"enableLogging,omitempty"`
}

// RportForwardBindDto is used for starting reverse port forwarding
type RportForwardBindDto struct {
	BindPort    int    `json:"bindPort"`
	ForwardHost string `json:"forwardHost"`
	ForwardPort int    `json:"forwardPort"`
}

// RportFwdStopDto is used for stopping reverse port forwarding
type RportFwdStopDto struct {
	BindPort int `json:"bindPort"`
}

// BrowserPivotSetupDto is used for starting browser pivot
type BrowserPivotSetupDto struct {
	Pid  int    `json:"pid"`
	Arch string `json:"arch"`
}

// --- Client Methods for Pivoting and Lateral Movement ---

// LinkSmb connects to an SMB beacon and re-establishes control
func (c *Client) LinkSmb(ctx context.Context, bid string, target string, pipe string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/link/smb", bid)
	req := LinkDto{
		Target: target,
		Pipe:   pipe,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute link smb: %w", err)
	}
	return &resp, nil
}

// LinkTcp connects to a TCP beacon and re-establishes control
func (c *Client) LinkTcp(ctx context.Context, bid string, target string, port int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/link/tcp", bid)
	req := ConnectDto{
		Target: target,
		Port:   port,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute link tcp: %w", err)
	}
	return &resp, nil
}

// Unlink disconnects from a named pipe or TCP beacon
func (c *Client) Unlink(ctx context.Context, bid string, host string, pid int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/unlink", bid)
	req := UnlinkDto{
		Host: host,
		Pid:  pid,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute unlink: %w", err)
	}
	return &resp, nil
}

// SpawnSsh spawns a temporary process to run an SSH client and attempt to login using username/password
func (c *Client) SpawnSsh(ctx context.Context, bid string, target string, port int, username string, password string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/ssh", bid)
	req := SshSpawnDto{
		Target:   target,
		Port:     port,
		Username: username,
		Password: password,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to spawn ssh: %w", err)
	}
	return &resp, nil
}

// SpawnSshKey spawns a temporary process to run an SSH client and attempt to login using an SSH key
// The keyPath should be a local file path to the SSH private key in PEM format
func (c *Client) SpawnSshKey(ctx context.Context, bid string, target string, port int, username string, keyPath string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/sshKey", bid)

	// Read and encode the SSH key file
	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SSH key file: %w", err)
	}
	encodedKey := base64.StdEncoding.EncodeToString(keyData)

	// Use @files/ prefix to indicate file upload
	keyFileKey := "key.pem"
	req := SshKeySpawnDto{
		Target:   target,
		Port:     port,
		Username: username,
		Key:      fmt.Sprintf("@files/%s", keyFileKey),
		Files: map[string]string{
			keyFileKey: encodedKey,
		},
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to spawn ssh with key: %w", err)
	}
	return &resp, nil
}

// ListRemoteExecMethods lists available remote command execution methods
func (c *Client) ListRemoteExecMethods(ctx context.Context, bid string) ([]RemoteExecInfoDto, error) {
	var methods []RemoteExecInfoDto
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/remoteExec/command", bid)
	if err := c.doRequest(ctx, "GET", endpoint, nil, &methods, true); err != nil {
		return nil, fmt.Errorf("failed to list remote exec methods: %w", err)
	}
	return methods, nil
}

// RemoteExec executes a command on a target via specific remote execution method
func (c *Client) RemoteExec(ctx context.Context, bid string, method string, target string, command string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/remoteExec/command", bid)
	req := RemoteExecDto{
		Method:  method,
		Target:  target,
		Command: command,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute remote exec: %w", err)
	}
	return &resp, nil
}

// ListJumpMethods lists available remote beacon execution methods for jump operations
func (c *Client) ListJumpMethods(ctx context.Context, bid string) ([]RemoteExploitInfoDto, error) {
	var methods []RemoteExploitInfoDto
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/remoteExec/beacon", bid)
	if err := c.doRequest(ctx, "GET", endpoint, nil, &methods, true); err != nil {
		return nil, fmt.Errorf("failed to list jump methods: %w", err)
	}
	return methods, nil
}

// Jump executes a beacon session on a remote target with the specified remote execution method
func (c *Client) Jump(ctx context.Context, bid string, exploit string, target string, listener string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/remoteExec/beacon", bid)
	req := JumpDto{
		Exploit:  exploit,
		Target:   target,
		Listener: listener,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute jump: %w", err)
	}
	return &resp, nil
}

// --- Client Methods for Tunneling ---

// Socks4Start starts a SOCKS4a server on the specified port
func (c *Client) Socks4Start(ctx context.Context, bid string, port int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/socks4Start", bid)
	req := Socks4StartDto{
		Port: port,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to start socks4: %w", err)
	}
	return &resp, nil
}

// Socks5Start starts a SOCKS5 server on the specified port with optional authentication
func (c *Client) Socks5Start(ctx context.Context, bid string, port int, auth *SocksAuthDto, enableLogging bool) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/socks5Start", bid)
	req := Socks5StartDto{
		Port:          port,
		Auth:          auth,
		EnableLogging: enableLogging,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to start socks5: %w", err)
	}
	return &resp, nil
}

// SocksStopAll stops all SOCKS servers and terminates existing connections
func (c *Client) SocksStopAll(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/socksStop/all", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to stop all socks: %w", err)
	}
	return &resp, nil
}

// SocksStop stops the specific SOCKS server on the given port and terminates existing connections
func (c *Client) SocksStop(ctx context.Context, bid string, port int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/socksStop/%d", bid, port)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to stop socks: %w", err)
	}
	return &resp, nil
}

// RportfwdStart starts reverse port forwarding on the specified bind port
// It forwards connections to the forwardHost:forwardPort
func (c *Client) RportfwdStart(ctx context.Context, bid string, bindPort int, forwardHost string, forwardPort int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/rportfwdStart/onTeamserver", bid)
	req := RportForwardBindDto{
		BindPort:    bindPort,
		ForwardHost: forwardHost,
		ForwardPort: forwardPort,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to start reverse port forward: %w", err)
	}
	return &resp, nil
}

// RportfwdStop stops reverse port forwarding on the specific bind port
func (c *Client) RportfwdStop(ctx context.Context, bid string, bindPort int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/rportfwdStop/onTeamserver", bid)
	req := RportFwdStopDto{
		BindPort: bindPort,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to stop reverse port forward: %w", err)
	}
	return &resp, nil
}

// BrowserPivotStart starts a browser pivot into the specified process
// To hijack authenticated web sessions, make sure the process is an Internet Explorer tab
func (c *Client) BrowserPivotStart(ctx context.Context, bid string, pid int, arch string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/inject/browserpivotStart", bid)
	req := BrowserPivotSetupDto{
		Pid:  pid,
		Arch: arch,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to start browser pivot: %w", err)
	}
	return &resp, nil
}

// BrowserPivotStop tears down the browser pivoting sessions associated with this beacon
func (c *Client) BrowserPivotStop(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/browserpivotStop", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to stop browser pivot: %w", err)
	}
	return &resp, nil
}
