package goharvest

import "time"

type UserAssignment struct {
	ID      int `json:"id"` // Unique ID for the user assignment.
	Project struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"project"` // An object containing the id, name, and code of the associated project.
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"user"` // An object containing the id and name of the associated user.
	IsActive         bool      `json:"is_active"`          // Whether the user assignment is active or archived.
	IsProjectManager bool      `json:"is_project_manager"` // Determines if the user has Project Manager permissions for the project.
	UseDefaultRates  bool      `json:"use_default_rates"`  // Determines which billable rate(s) will be used on the project for this user when bill_by is People. When true, the project will use the user’s default billable rates. When false, the project will use the custom rate defined on this user assignment.
	HourlyRate       float32   `json:"hourly_rate"`        // Custom rate used when the project’s bill_by is People and use_default_rates is false.
	Budget           float32   `json:"budget"`             // Budget used when the project’s budget_by is person.
	CreatedAt        time.Time `json:"created_at"`         // Date and time the user assignment was created.
	UpdatedAt        time.Time `json:"updated_at"`         // Date and time the user assignment was last updated.
}
