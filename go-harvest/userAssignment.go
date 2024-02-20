package goharvest

import "time"

type UserAssignment struct {
	// Unique ID for the user assignment.
	ID int `json:"id"`

	// An object containing the id, name, and code of the associated project.
	Project struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"project"`

	// An object containing the id and name of the associated user.
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"user"`

	// Whether the user assignment is active or archived.
	IsActive bool `json:"is_active"`

	// Determines if the user has Project Manager permissions for the project.
	IsProjectManager bool `json:"is_project_manager"`

	// Determines which billable rate(s) will be used on the project for this
	// user when bill_by is People. When true, the project will use the
	// user’s default billable rates. When false, the project will use the
	// custom rate defined on this user assignment.
	UseDefaultRates bool `json:"use_default_rates"`

	// Custom rate used when the project’s bill_by is People and use_default
	// rates is false.
	HourlyRate float64 `json:"hourly_rate"`

	// Budget used when the project’s budget_by is person.
	Budget float64 `json:"budget"`

	// Date and time the user assignment was created.
	CreatedAt time.Time `json:"created_at"`

	// Date and time the user assignment was last updated.
	UpdatedAt time.Time `json:"updated_at"`
}
