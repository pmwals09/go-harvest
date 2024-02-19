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
	// Unique ID for the task assignment.
	ID int `json:"id"`

	// An object containing the id, name, and code of the associated project.
	Project TaskAssignmentProject `json:"project"`

	// An object containing the id and name of the associated task.
	Task TaskAssignmentTask `json:"task"`

	// Whether the task assignment is active or archived.
	IsActive bool `json:"is_active"`

	// Whether the task assignment is billable or not. For example: if set to
	// true, all time tracked on this project for the associated task will be
	// marked as billable.
	Billable bool `json:"billable"`

	// Rate used when the project’s bill_by is Tasks.
	HourlyRate float64 `json:"hourly_rate"`

	// Budget used when the project’s budget_by is task or task_fees.
	Budget float64 `json:"budget"`

	// Date and time the task assignment was created.
	CreatedAt time.Time `json:"created_at"`

	// Date and time the task assignment was last updated.
	UpdatedAt time.Time `json:"updated_at"`
}
