package csclient

import (
	"context"
	"fmt"
)

// ============================================================
// Listener DTOs
// ============================================================

// ListenerBaseDto represents basic listener information
type ListenerBaseDto struct {
	Name    string `json:"name"`
	Color   string `json:"color"`   // DEFAULT, GREEN, PINK, YELLOW, GRAY, BLUE
	Error   string `json:"error"`   // Error message if listener failed
	Payload string `json:"payload"` // Payload type
}

// ============================================================
// Listener Management Methods
// ============================================================

// ListListeners retrieves all configured listeners
func (c *Client) ListListeners(ctx context.Context) ([]ListenerBaseDto, error) {
	var listeners []ListenerBaseDto
	if err := c.doRequest(ctx, "GET", "/api/v1/listeners", nil, &listeners, true); err != nil {
		return nil, fmt.Errorf("failed to list listeners: %w", err)
	}
	return listeners, nil
}

// GetListener retrieves a specific listener by name
func (c *Client) GetListener(ctx context.Context, name string) (map[string]interface{}, error) {
	var listener map[string]interface{}
	endpoint := fmt.Sprintf("/api/v1/listeners/%s", name)
	if err := c.doRequest(ctx, "GET", endpoint, nil, &listener, true); err != nil {
		return nil, fmt.Errorf("failed to get listener: %w", err)
	}
	return listener, nil
}

// DeleteListener removes a listener by name
func (c *Client) DeleteListener(ctx context.Context, name string) error {
	endpoint := fmt.Sprintf("/api/v1/listeners/%s", name)
	if err := c.doRequest(ctx, "DELETE", endpoint, nil, nil, true); err != nil {
		return fmt.Errorf("failed to delete listener: %w", err)
	}
	return nil
}

// ============================================================
// HTTP Listener Methods
// ============================================================

// AddHttpListener creates a new HTTP listener
// config should contain the HttpListenerDto fields as per the OpenAPI spec:
// - name (required): listener name
// - color (required): DEFAULT, GREEN, PINK, YELLOW, GRAY, BLUE
// - hosts (required): array of HTTP hosts
// - host (required): HTTP host for stager
// - ignoreProxySettings (required): boolean
// - httpPort, httpBindPort, httpHostHeader, hostRotationStrategy, maxRetryStrategy, profile, httpProxy, guardRails (optional)
func (c *Client) AddHttpListener(ctx context.Context, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	if err := c.doRequest(ctx, "POST", "/api/v1/listeners/http", config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to add http listener: %w", err)
	}
	return resp, nil
}

// UpdateHttpListener updates an existing HTTP listener
func (c *Client) UpdateHttpListener(ctx context.Context, name string, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	endpoint := fmt.Sprintf("/api/v1/listeners/http/%s", name)
	if err := c.doRequest(ctx, "PUT", endpoint, config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to update http listener: %w", err)
	}
	return resp, nil
}

// ============================================================
// HTTPS Listener Methods
// ============================================================

// AddHttpsListener creates a new HTTPS listener
// config should contain the HttpsListenerDto fields as per the OpenAPI spec:
// - name (required): listener name
// - color (required): DEFAULT, GREEN, PINK, YELLOW, GRAY, BLUE
// - hosts (required): array of HTTPS hosts
// - host (required): HTTPS host for stager
// - ignoreProxySettings (required): boolean
// - httpPort, httpBindPort, httpHostHeader, hostRotationStrategy, maxRetryStrategy, profile, httpProxy, guardRails (optional)
func (c *Client) AddHttpsListener(ctx context.Context, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	if err := c.doRequest(ctx, "POST", "/api/v1/listeners/https", config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to add https listener: %w", err)
	}
	return resp, nil
}

// UpdateHttpsListener updates an existing HTTPS listener
func (c *Client) UpdateHttpsListener(ctx context.Context, name string, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	endpoint := fmt.Sprintf("/api/v1/listeners/https/%s", name)
	if err := c.doRequest(ctx, "PUT", endpoint, config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to update https listener: %w", err)
	}
	return resp, nil
}

// ============================================================
// DNS Listener Methods
// ============================================================

// AddDnsListener creates a new DNS listener
// config should contain the DnsListenerDto fields as per the OpenAPI spec:
// - name (required): listener name
// - color (required): DEFAULT, GREEN, PINK, YELLOW, GRAY, BLUE
// - hosts (required): array of DNS hosts
// - host (required): DNS host for stager
// - dnsBindPort, hostRotationStrategy, maxRetryStrategy, profile, guardRails (optional)
func (c *Client) AddDnsListener(ctx context.Context, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	if err := c.doRequest(ctx, "POST", "/api/v1/listeners/dns", config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to add dns listener: %w", err)
	}
	return resp, nil
}

// UpdateDnsListener updates an existing DNS listener
func (c *Client) UpdateDnsListener(ctx context.Context, name string, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	endpoint := fmt.Sprintf("/api/v1/listeners/dns/%s", name)
	if err := c.doRequest(ctx, "PUT", endpoint, config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to update dns listener: %w", err)
	}
	return resp, nil
}

// ============================================================
// TCP Listener Methods
// ============================================================

// AddTcpListener creates a new TCP listener
// config should contain the TcpListenerDto fields as per the OpenAPI spec:
// - name (required): listener name
// - color (required): DEFAULT, GREEN, PINK, YELLOW, GRAY, BLUE
// - port (required): TCP port (1-65535)
// - localHostOnly (required): boolean
// - guardRails (optional)
func (c *Client) AddTcpListener(ctx context.Context, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	if err := c.doRequest(ctx, "POST", "/api/v1/listeners/tcp", config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to add tcp listener: %w", err)
	}
	return resp, nil
}

// UpdateTcpListener updates an existing TCP listener
func (c *Client) UpdateTcpListener(ctx context.Context, name string, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	endpoint := fmt.Sprintf("/api/v1/listeners/tcp/%s", name)
	if err := c.doRequest(ctx, "PUT", endpoint, config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to update tcp listener: %w", err)
	}
	return resp, nil
}

// ============================================================
// SMB Listener Methods
// ============================================================

// AddSmbListener creates a new SMB listener
// config should contain the SmbListenerDto fields as per the OpenAPI spec:
// - name (required): listener name
// - color (required): DEFAULT, GREEN, PINK, YELLOW, GRAY, BLUE
// - pipename (required): named pipe name (1-118 chars)
// - guardRails (optional)
func (c *Client) AddSmbListener(ctx context.Context, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	if err := c.doRequest(ctx, "POST", "/api/v1/listeners/smb", config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to add smb listener: %w", err)
	}
	return resp, nil
}

// UpdateSmbListener updates an existing SMB listener
func (c *Client) UpdateSmbListener(ctx context.Context, name string, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	endpoint := fmt.Sprintf("/api/v1/listeners/smb/%s", name)
	if err := c.doRequest(ctx, "PUT", endpoint, config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to update smb listener: %w", err)
	}
	return resp, nil
}

// ============================================================
// External C2 Listener Methods
// ============================================================

// AddExternalC2Listener creates a new External C2 listener
// config should contain the ExternalC2ListenerDto fields as per the OpenAPI spec:
// - name (required): listener name
// - color (required): DEFAULT, GREEN, PINK, YELLOW, GRAY, BLUE
// - port (required): TCP port (1-65535)
// - localHostOnly (required): boolean
func (c *Client) AddExternalC2Listener(ctx context.Context, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	if err := c.doRequest(ctx, "POST", "/api/v1/listeners/externalC2", config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to add external c2 listener: %w", err)
	}
	return resp, nil
}

// UpdateExternalC2Listener updates an existing External C2 listener
func (c *Client) UpdateExternalC2Listener(ctx context.Context, name string, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	endpoint := fmt.Sprintf("/api/v1/listeners/externalC2/%s", name)
	if err := c.doRequest(ctx, "PUT", endpoint, config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to update external c2 listener: %w", err)
	}
	return resp, nil
}

// ============================================================
// User Defined C2 Listener Methods
// ============================================================

// AddUserDefinedC2Listener creates a new User Defined C2 listener
// config should contain the UserDefinedC2ListenerDto fields as per the OpenAPI spec:
// - name (required): listener name
// - color (required): DEFAULT, GREEN, PINK, YELLOW, GRAY, BLUE
// - port (required): TCP port (1-65535)
// - localHostOnly (required): boolean
// - debugOnly (required): boolean
// - udc2Bof (required): BOF file reference (@files/... or @artifacts/...)
// - files (optional): map of filename -> base64 content (required if using @files/)
// - guardRails (optional)
func (c *Client) AddUserDefinedC2Listener(ctx context.Context, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	if err := c.doRequest(ctx, "POST", "/api/v1/listeners/userDefinedC2", config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to add user defined c2 listener: %w", err)
	}
	return resp, nil
}

// UpdateUserDefinedC2Listener updates an existing User Defined C2 listener
func (c *Client) UpdateUserDefinedC2Listener(ctx context.Context, name string, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	endpoint := fmt.Sprintf("/api/v1/listeners/userDefinedC2/%s", name)
	if err := c.doRequest(ctx, "PUT", endpoint, config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to update user defined c2 listener: %w", err)
	}
	return resp, nil
}

// ============================================================
// Foreign HTTP Listener Methods
// ============================================================

// AddForeignHttpListener creates a new Foreign HTTP listener
// config should contain the ForeignHttpListenerDto fields as per the OpenAPI spec:
// - name (required): listener name
// - color (required): DEFAULT, GREEN, PINK, YELLOW, GRAY, BLUE
// - host (required): HTTP host
// - port (required): HTTP port (1-65535)
func (c *Client) AddForeignHttpListener(ctx context.Context, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	if err := c.doRequest(ctx, "POST", "/api/v1/listeners/foreignHttp", config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to add foreign http listener: %w", err)
	}
	return resp, nil
}

// UpdateForeignHttpListener updates an existing Foreign HTTP listener
func (c *Client) UpdateForeignHttpListener(ctx context.Context, name string, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	endpoint := fmt.Sprintf("/api/v1/listeners/foreignHttp/%s", name)
	if err := c.doRequest(ctx, "PUT", endpoint, config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to update foreign http listener: %w", err)
	}
	return resp, nil
}

// ============================================================
// Foreign HTTPS Listener Methods
// ============================================================

// AddForeignHttpsListener creates a new Foreign HTTPS listener
// config should contain the ForeignHttpsListenerDto fields as per the OpenAPI spec:
// - name (required): listener name
// - color (required): DEFAULT, GREEN, PINK, YELLOW, GRAY, BLUE
// - host (required): HTTPS host
// - port (required): HTTPS port (1-65535)
func (c *Client) AddForeignHttpsListener(ctx context.Context, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	if err := c.doRequest(ctx, "POST", "/api/v1/listeners/foreignHttps", config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to add foreign https listener: %w", err)
	}
	return resp, nil
}

// UpdateForeignHttpsListener updates an existing Foreign HTTPS listener
func (c *Client) UpdateForeignHttpsListener(ctx context.Context, name string, config map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	endpoint := fmt.Sprintf("/api/v1/listeners/foreignHttps/%s", name)
	if err := c.doRequest(ctx, "PUT", endpoint, config, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to update foreign https listener: %w", err)
	}
	return resp, nil
}
