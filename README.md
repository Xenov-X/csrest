# csrest

Go client library for the Cobalt Strike REST API (Cobalt Strike v4.12+).

> [!CAUTION]
> This project is in early stage active development - Expect significant changes. The API is also in BETA, so may also be subject to change.

## Overview

This package provides a type-safe Go interface for interacting with the Cobalt Strike REST API. It handles authentication, beacon management, BOF execution, and task retrieval.

## Features

- JWT-based authentication with configurable token expiration
- Beacon listing and details retrieval
- BOF (Beacon Object File) execution with multiple argument formats:
  - String arguments
  - Packed binary arguments
  - Typed arguments with auto-packing
- Task management with automatic polling
- Timeout handling and context support
- Custom HTTP client support (for TLS configuration)

## Installation

```bash
go get github.com/xenov-x/csrest
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    csclient "github.com/xenov-x/csrest"
)

func main() {
    ctx := context.Background()
    
    // Create client
    client := csclient.NewClient("10.0.0.1", 50443)
    
    // Authenticate
    auth, err := client.Login(ctx, "operator", "password", 3600000)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Authenticated: %s\n", auth.AccessToken)
    
    // List beacons
    beacons, err := client.ListBeacons(ctx)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, beacon := range beacons {
        fmt.Printf("Beacon %s: %s@%s\n", beacon.BID, beacon.User, beacon.Computer)
    }
}
```

## Usage Examples

### Authentication

```go
client := csclient.NewClient("10.0.0.1", 50433)

// Login with 1 hour token expiration
auth, err := client.Login(ctx, "username", "password", 3600000)
if err != nil {
    log.Fatal(err)
}
```

### Beacon Management

```go
// List all beacons
beacons, err := client.ListBeacons(ctx)

// Get specific beacon
beacon, err := client.GetBeacon(ctx, "beacon-id-123")

// Access beacon information
fmt.Printf("User: %s\n", beacon.User)
fmt.Printf("Computer: %s\n", beacon.Computer)
fmt.Printf("Last Check-in: %s\n", beacon.LastCheckinFormatted)
fmt.Printf("Sleep: %ds (Jitter: %d%%)\n", beacon.Sleep.Sleep, beacon.Sleep.Jitter)
fmt.Printf("Admin: %v\n", beacon.IsAdmin)
```

### BOF Execution - String Arguments

```go
import "encoding/base64"

// Read BOF file
bofData, err := os.ReadFile("/path/to/bof.o")
if err != nil {
    log.Fatal(err)
}

// Execute BOF with string arguments
resp, err := client.ExecuteBOFString(ctx, beaconID, csclient.InlineExecuteStringDto{
    BOF: "@files/bof.o",
    Entrypoint: "go",
    Arguments: "arg1 arg2 arg3",
    Files: map[string]string{
        "bof.o": base64.StdEncoding.EncodeToString(bofData),
    },
})

// Get task ID
taskID := resp.TaskID
```

### BOF Execution - Typed Arguments

```go
// Execute BOF with typed arguments (auto-packed)
resp, err := client.ExecuteBOFPack(ctx, beaconID, csclient.InlineExecutePackDto{
    BOF: "@files/bof.o",
    Entrypoint: "go",
    Arguments: []csclient.BOFArgument{
        csclient.StringArg{Type: "string", Value: "target.exe"},
        csclient.IntArg{Type: "int", Value: 1234},
        csclient.WStringArg{Type: "wstring", Value: "Wide string"},
        csclient.ShortArg{Type: "short", Value: 100},
    },
    Files: map[string]string{
        "bof.o": base64.StdEncoding.EncodeToString(bofData),
    },
})
```

### Task Management

```go
// Get task details
task, err := client.GetTask(ctx, taskID)

// Wait for task completion (with timeout)
task, err := client.WaitForTaskCompletion(ctx, taskID, 5*time.Minute)
if err != nil {
    log.Fatal(err)
}

// Check task status
switch task.TaskStatus {
case csclient.TaskStatusCompleted:
    fmt.Println("Task completed successfully")
case csclient.TaskStatusFailed:
    fmt.Println("Task failed")
case csclient.TaskStatusOutputReceived:
    fmt.Println("Output received")
}

// Parse output
for _, result := range task.Result {
    if output, ok := result["output"].(string); ok {
        fmt.Println(output)
    }
}
```

### Privilege Operations

```go
// Get current user ID (whoami)
resp, err := client.GetUID(ctx, beaconID)
task, err := client.WaitForTaskCompletion(ctx, resp.TaskID, 1*time.Minute)

// Attempt privilege escalation to SYSTEM
resp, err := client.GetSystem(ctx, beaconID)
task, err := client.WaitForTaskCompletion(ctx, resp.TaskID, 2*time.Minute)
```

### Custom HTTP Client

```go
import (
    "crypto/tls"
    "net/http"
)

// Create custom HTTP client with TLS config
httpClient := &http.Client{
    Timeout: 30 * time.Second,
    Transport: &http.Transport{
        TLSClientConfig: &tls.Config{
            InsecureSkipVerify: true,
            // Add custom CA certificates
            // RootCAs: certPool,
        },
    },
}

client := csclient.NewClient("10.0.0.1", 50433)
client.SetHTTPClient(httpClient)
```

## API Reference

### Client

- `NewClient(host string, port int) *Client` - Create new client
- `SetHTTPClient(client *http.Client)` - Set custom HTTP client
- `Login(ctx, username, password string, durationMs int) (*AuthDto, error)` - Authenticate

### Beacons

- `ListBeacons(ctx) ([]BeaconDto, error)` - List all beacons
- `GetBeacon(ctx, bid string) (*BeaconDto, error)` - Get beacon details
- `ExecuteBOFString(ctx, bid string, req InlineExecuteStringDto) (*AsyncCommandResponse, error)`
- `ExecuteBOFPacked(ctx, bid string, req InlineExecutePackedDto) (*AsyncCommandResponse, error)`
- `ExecuteBOFPack(ctx, bid string, req InlineExecutePackDto) (*AsyncCommandResponse, error)`
- `GetUID(ctx, bid string) (*AsyncCommandResponse, error)` - Get user ID
- `GetSystem(ctx, bid string) (*AsyncCommandResponse, error)` - Elevate to SYSTEM

### Tasks

- `GetTask(ctx, taskID string) (*TaskDetailDto, error)` - Get task details
- `ListTasks(ctx) ([]TaskSummaryDto, error)` - List all tasks
- `GetBeaconTasksSummary(ctx, bid string) ([]TaskSummaryDto, error)` - Get beacon tasks
- `WaitForTaskCompletion(ctx, taskID string, timeout time.Duration) (*TaskDetailDto, error)` - Poll until complete

## Types

### BeaconDto

Contains comprehensive beacon information:

- `BID` - Beacon ID
- `User` - User context
- `Computer` - Hostname
- `Internal` / `External` - IP addresses
- `Process` / `PID` - Process information
- `IsAdmin` - Admin status
- `OS` / `Version` / `Build` - OS information
- `Sleep` - Sleep configuration (SleepDto with Sleep and Jitter)
- `LastCheckinTime` / `LastCheckinFormatted` - Check-in information
- Many more fields...

### BOF Arguments

- `StringArg` - ASCII string
- `WStringArg` - Wide (Unicode) string
- `IntArg` - 32-bit integer
- `ShortArg` - 16-bit integer
- `BinaryArg` - Binary data (base64 encoded)

### Task Status

- `TaskStatusNotFound` - Task not found
- `TaskStatusInProgress` - Currently executing
- `TaskStatusCompleted` - Execution complete
- `TaskStatusFailed` - Execution failed
- `TaskStatusOutputReceived` - Output available

## Error Handling

All methods return errors that should be checked:

```go
task, err := client.GetTask(ctx, taskID)
if err != nil {
    log.Printf("Failed to get task: %v", err)
    return
}
```

## Context Support

All API methods accept a context for cancellation and timeout:

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

beacons, err := client.ListBeacons(ctx)
```

## Thread Safety

The client is safe for concurrent use from multiple goroutines.

## Serious Notes

- This tool is for authorized penetration testing only
- Always obtain proper authorization before use
- Review workflows before execution


