package goharvest

import "time"

type TaskAssignmentProject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
type TaskAssignmentTask struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type TaskAssignment struct {
	ID         int                   `json:"id"`          // Unique ID for the task assignment.
	Project    TaskAssignmentProject `json:"project"`     // An object containing the id, name, and code of the associated project.
	Task       TaskAssignmentTask    `json:"task"`        // An object containing the id and name of the associated task.
	IsActive   bool                  `json:"is_active"`   // Whether the task assignment is active or archived.
	Billable   bool                  `json:"billable"`    // Whether the task assignment is billable or not. For example: if set to true, all time tracked on this project for the associated task will be marked as billable.
	HourlyRate float32               `json:"hourly_rate"` // Rate used when the project’s bill_by is Tasks.
	Budget     float32               `json:"budget"`      // Budget used when the project’s budget_by is task or task_fees.
	CreatedAt  time.Time             `json:"created_at"`  // Date and time the task assignment was created.
	UpdatedAt  time.Time             `json:"updated_at"`  // Date and time the task assignment was last updated.
}
