
package model

import "time"
    
// Task Collection of documents with tasks details
type Task struct {
	// Completed Indicate task completion state
	Completed bool `json:"completed"`
	// CompletedAt Task completion date
	CompletedAt time.Time `json:"completed_at"`
	// Details Detail explanation of the task
	Details string `json:"details"`
	// DueAt Task due date
	DueAt time.Time `json:"due_at"`
	// Id A unique identifier for the task
	Id int64 `json:"id" tigris:"primaryKey:1,autoGenerate"`
	// Name Name of the task
	Name string `json:"name"`
	// Tags The list of task categories
	Tags []string `json:"tags"`
}
