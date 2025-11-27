package csclient

import "time"

// APIError represents an API error with retry information
type APIError struct {
	StatusCode int
	Message    string
	Retryable  bool
}

func (e *APIError) Error() string {
	return e.Message
}

// LoginRequest represents the login request payload
type LoginRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	DurationMs int    `json:"duration_ms,omitempty"`
}

// AuthDto represents the authentication response
type AuthDto struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// AsyncCommandResponse represents the response from async commands
type AsyncCommandResponse struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	StatusURL string `json:"statusUrl,omitempty"`
	TaskID    string `json:"taskId,omitempty"`
}

// TaskStatus represents the status of a task
type TaskStatus string

const (
	TaskStatusNotFound       TaskStatus = "NOT_FOUND"
	TaskStatusInProgress     TaskStatus = "IN_PROGRESS"
	TaskStatusCompleted      TaskStatus = "COMPLETED"
	TaskStatusFailed         TaskStatus = "FAILED"
	TaskStatusOutputReceived TaskStatus = "OUTPUT_RECEIVED"
)

// TaskSummaryDto represents a task summary
type TaskSummaryDto struct {
	TaskID      string     `json:"taskId"`
	BID         string     `json:"bid"`
	JID         int        `json:"jid,omitempty"`
	TaskCommand string     `json:"taskCommand"`
	User        string     `json:"user"`
	Created     time.Time  `json:"created"`
	Updated     *time.Time `json:"updated,omitempty"`
	TaskStatus  TaskStatus `json:"taskStatus"`
}

// TaskDetailDto represents detailed task information with outputs
type TaskDetailDto struct {
	TaskSummaryDto
	Result  []map[string]interface{} `json:"result,omitempty"`
	Error   []ErrorMessageDto        `json:"error,omitempty"`
	Tactics []string                 `json:"tactics,omitempty"`
}

// ErrorMessageDto represents an error message
type ErrorMessageDto struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

// SleepDto represents beacon sleep configuration
type SleepDto struct {
	Sleep  int `json:"sleep"`  // Sleep time in seconds
	Jitter int `json:"jitter"` // Jitter percentage (0-99)
}

// BeaconDto represents beacon information
type BeaconDto struct {
	BID                  string    `json:"bid"`
	PBID                 string    `json:"pbid,omitempty"`
	Computer             string    `json:"computer"`
	User                 string    `json:"user"`
	Impersonated         string    `json:"impersonated,omitempty"`
	IsAdmin              bool      `json:"isAdmin,omitempty"`
	Process              string    `json:"process"`
	PID                  int       `json:"pid"`
	Host                 string    `json:"host,omitempty"`
	Internal             string    `json:"internal"`
	External             string    `json:"external"`
	OS                   string    `json:"os,omitempty"`
	Version              string    `json:"version,omitempty"`
	Build                int       `json:"build,omitempty"`
	Charset              string    `json:"charset,omitempty"`
	SystemArch           string    `json:"systemArch,omitempty"`
	BeaconArch           string    `json:"beaconArch,omitempty"`
	Session              string    `json:"session"`
	Listener             string    `json:"listener"`
	PivotHint            string    `json:"pivotHint,omitempty"`
	Port                 int       `json:"port,omitempty"`
	Note                 string    `json:"note,omitempty"`
	Color                string    `json:"color,omitempty"`
	Alive                bool      `json:"alive"`
	LinkState            string    `json:"linkState,omitempty"`
	LastCheckinTime      time.Time `json:"lastCheckinTime"`
	LastCheckinMs        int       `json:"lastCheckinMs"`
	LastCheckinFormatted string    `json:"lastCheckinFormatted"`
	Sleep                SleepDto  `json:"sleep"`
	SupportsSleep        bool      `json:"supportsSleep"`
}

// InlineExecuteStringDto represents BOF execution with string arguments
type InlineExecuteStringDto struct {
	BOF        string            `json:"bof"`
	Entrypoint string            `json:"entrypoint,omitempty"`
	Arguments  string            `json:"arguments,omitempty"`
	Files      map[string]string `json:"files,omitempty"`
}

// InlineExecutePackedDto represents BOF execution with packed arguments
type InlineExecutePackedDto struct {
	BOF        string            `json:"bof"`
	Entrypoint string            `json:"entrypoint,omitempty"`
	Arguments  string            `json:"arguments,omitempty"` // base64 encoded packed args
	Files      map[string]string `json:"files,omitempty"`
}

// InlineExecutePackDto represents BOF execution with typed arguments
type InlineExecutePackDto struct {
	BOF        string            `json:"bof"`
	Entrypoint string            `json:"entrypoint,omitempty"`
	Arguments  []BOFArgument     `json:"arguments,omitempty"`
	Files      map[string]string `json:"files,omitempty"`
}

// BOFArgument is an interface for different BOF argument types
type BOFArgument interface {
	bofArgument()
}

// BinaryArg represents a binary argument
type BinaryArg struct {
	Type  string `json:"type"`
	Value string `json:"value"` // base64 encoded
}

func (BinaryArg) bofArgument() {}

// IntArg represents an integer argument
type IntArg struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

func (IntArg) bofArgument() {}

// ShortArg represents a short integer argument
type ShortArg struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

func (ShortArg) bofArgument() {}

// StringArg represents a string argument
type StringArg struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (StringArg) bofArgument() {}

// WStringArg represents a wide string argument
type WStringArg struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (WStringArg) bofArgument() {}

// PowerShellDto represents PowerShell command execution request
type PowerShellDto struct {
	Commandlet string `json:"commandlet"`
	Arguments  string `json:"arguments,omitempty"`
}

// UploadDto represents file upload request
type UploadDto struct {
	File  string            `json:"file"`            // @files/filename reference to files map
	Files map[string]string `json:"files,omitempty"` // Map of filename -> base64 content
}

// EmptyDto represents an empty request body
type EmptyDto struct{}
