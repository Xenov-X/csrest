package csclient

import (
	"context"
	"fmt"
	"time"
)

// GetTask retrieves detailed information about a specific task
func (c *Client) GetTask(ctx context.Context, taskID string) (*TaskDetailDto, error) {
	var task TaskDetailDto
	path := fmt.Sprintf("/api/v1/tasks/%s", taskID)
	if err := c.doRequest(ctx, "GET", path, nil, &task, true); err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}
	return &task, nil
}

// ListTasks retrieves all tasks
func (c *Client) ListTasks(ctx context.Context) ([]TaskSummaryDto, error) {
	var tasks []TaskSummaryDto
	if err := c.doRequest(ctx, "GET", "/api/v1/tasks", nil, &tasks, true); err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}
	return tasks, nil
}

// GetBeaconTasksSummary retrieves task summaries for a specific beacon
func (c *Client) GetBeaconTasksSummary(ctx context.Context, bid string) ([]TaskSummaryDto, error) {
	var tasks []TaskSummaryDto
	path := fmt.Sprintf("/api/v1/beacons/%s/tasks/summary", bid)
	if err := c.doRequest(ctx, "GET", path, nil, &tasks, true); err != nil {
		return nil, fmt.Errorf("failed to get beacon tasks: %w", err)
	}
	return tasks, nil
}

// GetBeaconTasksDetail retrieves detailed tasks for a specific beacon
func (c *Client) GetBeaconTasksDetail(ctx context.Context, bid string) ([]TaskDetailDto, error) {
	var tasks []TaskDetailDto
	path := fmt.Sprintf("/api/v1/beacons/%s/tasks/detail", bid)
	if err := c.doRequest(ctx, "GET", path, nil, &tasks, true); err != nil {
		return nil, fmt.Errorf("failed to get beacon task details: %w", err)
	}
	return tasks, nil
}

// WaitForTaskCompletion polls a task until it completes or times out
func (c *Client) WaitForTaskCompletion(ctx context.Context, taskID string, timeout time.Duration) (*TaskDetailDto, error) {
	deadline := time.Now().Add(timeout)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			if time.Now().After(deadline) {
				return nil, fmt.Errorf("timeout waiting for task completion")
			}

			task, err := c.GetTask(ctx, taskID)
			if err != nil {
				return nil, err
			}

			// Log current task status for debugging
			fmt.Printf("[CSREST DEBUG] Task %s status: %s, command: %s\n", taskID, task.TaskStatus, task.TaskCommand)

			if task.TaskStatus == TaskStatusCompleted ||
			   task.TaskStatus == TaskStatusOutputReceived ||
			   task.TaskStatus == TaskStatusFailed {
				return task, nil
			}
		}
	}
}
