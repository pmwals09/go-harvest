package goharvest

import "time"

type ProjectAssignmentProject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type ProjectAssignmentClient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProjectAssignment struct {
	ID               int                      `json:"id"`                 // Unique ID for the project assignment.
	IsActive         bool                     `json:"is_active"`          // Whether the project assignment is active or archived.
	IsProjectManager bool                     `json:"is_project_manager"` // Determines if the user has Project Manager permissions for the project.
	UseDefaultRates  bool                     `json:"use_default_rates"`  // Determines which billable rate(s) will be used on the project for this user when `bill_by` is `People`. When `true`, the project will use the user’s default billable rates. When `false`, the project will use the custom rate defined on this user assignment.
	HourlyRate       float32                  `json:"hourly_rate"`        // Custom rate used when the project’s bill_by is People and use_default_rates is false.
	Budget           float32                  `json:"budget"`             // Budget used when the project’s budget_by is person.
	CreatedAt        time.Time                `json:"created_at"`         // Date and time the project assignment was created.
	UpdatedAt        time.Time                `json:"updated_at"`         // Date and time the project assignment was last updated.
	Project          ProjectAssignmentProject `json:"project"`            // An object containing the assigned project id, name, and code.
	Client           ProjectAssignmentClient  `json:"client"`             // An object containing the project’s client id and name.
	TaskAssignments  []TaskAssignment         `json:"task_assignments"`   // Array of task assignment objects associated with the project.
}

type ProjectAssignmentResponse struct {
	ProjectAssignments []ProjectAssignment `json:"project_assignments"`
	Pagination
}
