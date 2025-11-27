package csclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is the Cobalt Strike API client
type Client struct {
	baseURL    string
	httpClient *http.Client
	token      string
	maxRetries int
	retryDelay time.Duration
}

// NewClient creates a new Cobalt Strike API client
func NewClient(host string, port int) *Client {
	return &Client{
		baseURL: fmt.Sprintf("https://%s:%d", host, port),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		maxRetries: 3,
		retryDelay: 2 * time.Second,
	}
}

// SetHTTPClient allows setting a custom HTTP client (e.g., for custom TLS config)
func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

// SetRetryPolicy sets the retry policy for failed requests
func (c *Client) SetRetryPolicy(maxRetries int, retryDelay time.Duration) {
	c.maxRetries = maxRetries
	c.retryDelay = retryDelay
}

// Login authenticates with the Cobalt Strike server
func (c *Client) Login(ctx context.Context, username, password string, durationMs int) (*AuthDto, error) {
	req := LoginRequest{
		Username:   username,
		Password:   password,
		DurationMs: durationMs,
	}

	var auth AuthDto
	if err := c.doRequest(ctx, "POST", "/api/auth/login", req, &auth, false); err != nil {
		return nil, fmt.Errorf("login failed: %w", err)
	}

	c.token = auth.AccessToken
	return &auth, nil
}

// doRequest performs an HTTP request with retry logic
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}, result interface{}, requireAuth bool) error {
	var lastErr error

	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(c.retryDelay):
			}
		}

		err := c.doRequestOnce(ctx, method, path, body, result, requireAuth)
		if err == nil {
			return nil
		}

		lastErr = err

		// Don't retry on certain errors
		if isNonRetryableError(err) {
			return lastErr
		}

		// Don't retry if context is cancelled
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	return fmt.Errorf("request failed after %d attempts: %w", c.maxRetries+1, lastErr)
}

// doRequestOnce performs a single HTTP request
func (c *Client) doRequestOnce(ctx context.Context, method, path string, body interface{}, result interface{}, requireAuth bool) error {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return &APIError{
				StatusCode: 0,
				Message:    fmt.Sprintf("failed to marshal request: %v", err),
				Retryable:  false,
			}
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, reqBody)
	if err != nil {
		return &APIError{
			StatusCode: 0,
			Message:    fmt.Sprintf("failed to create request: %v", err),
			Retryable:  false,
		}
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if requireAuth && c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &APIError{
			StatusCode: 0,
			Message:    fmt.Sprintf("request failed: %v", err),
			Retryable:  true,
		}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("failed to read response: %v", err),
			Retryable:  true,
		}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		retryable := resp.StatusCode >= 500 || resp.StatusCode == 429 // Retry on server errors and rate limits
		msg := string(respBody)
		if msg == "" {
			msg = fmt.Sprintf("HTTP %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		}
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    msg,
			Retryable:  retryable,
		}
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return &APIError{
				StatusCode: resp.StatusCode,
				Message:    fmt.Sprintf("failed to unmarshal response: %v", err),
				Retryable:  false,
			}
		}
	}

	return nil
}

// isNonRetryableError checks if an error should not be retried
func isNonRetryableError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return !apiErr.Retryable
	}
	return false
}
