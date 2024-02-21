package goharvest

import (
	"encoding/json"
	"time"
)

type User struct {
	// Unique ID for the user.
	ID int `json:"id"`

	// The first name of the user.
	FirstName string `json:"first_name"`

	// The last name of the user.
	LastName string `json:"last_name"`

	// The email address of the user.
	Email string `json:"email"`

	// The user’s telephone number.
	Telephone string `json:"telephone"`

	// The user’s timezone.
	Timezone string `json:"timezone"`

	// Whether the user should be automatically added to future projects.
	HasAccessToAllFutureProjects bool `json:"has_access_to_all_future_projects"`

	// Whether the user is a contractor or an employee.
	IsContractor bool `json:"is_contractor"`

	// Whether the user is active or archived.
	IsActive bool `json:"is_active"`

	// The number of hours per week this person is available to work in
	// seconds, in half hour increments. For example, if a person’s capacity
	// is 35 hours, the API will return 126000 seconds.
	WeeklyCapacity int `json:"weekly_capacity"`

	// The billable rate to use for this user when they are added to
	// a project.
	DefaultHourlyRate float64 `json:"default_hourly_rate"`

	// The cost rate to use for this user when calculating a project’s costs
	// vs billable amount.
	CostRate float64 `json:"cost_rate"`

	// Descriptive names of the business roles assigned to this person. They
	// can be used for filtering reports, and have no effect in their
	// permissions in Harvest.
	Roles []string `json:"roles"`

	//help.getharvest.com/api-v2/users-api/users/users/#access-roles) that
	// determine the user’s permissions in Harvest. Possible values:
	// `administrator`, `manager` or `member`. Users with the manager role
	// can additionally be granted one or more of these roles:
	// `project_creator`, `billable_rates_manager`,
	// `managed_projects_invoice_drafter`,
	// `managed_projects_invoice_manager`, `client_and_task_manager`,
	// `time_and_expenses_manager`, `estimates_manager`.
	AccessRoles []string `json:"access_roles"` // [Access role(s)](https:

	// The URL to the user’s avatar image.
	AvatarURL string `json:"avatar_url"`

	// Date and time the user was created.
	CreatedAt time.Time `json:"created_at"`

	// Date and time the user was last updated.
	UpdatedAt time.Time `json:"updated_at"`
}

type GetProjectAssignmentParameters struct {
	// Only return project assignments that have been updated since the given
	// date and time.
	UpdatedSince time.Time `json:"updated_since" url:"updated_since,omitempty"`

	// DEPRECATED The page number to use in pagination. For instance, if
	// you make a list request and receive 2000 records, your subsequent call
	// can include page=2 to retrieve the next page of the list. (Default: 1)
	Page int `json:"page" url:"page,omitempty"`

	// The number of records to return per page. Can range between 1 and
	// 2000.  (Default: 2000)
	PerPage int `json:"per_page" url:"per_page,omitempty"`
}

// Returns a list of your active project assignments for the currently
// authenticated user. The project assignments are returned sorted by
// creation date, with the most recently created project assignments
// appearing first.
func (c *Client) GetMyProjectAssignments(params GetProjectAssignmentParameters) (ProjectAssignmentResponse, error) {
	pa := ProjectAssignmentResponse{}
	urlTail, err := buildPathWithParams[GetProjectAssignmentParameters]("/v2/users/me/project_assignments", params)
	if err != nil {
		return pa, err
	}
	res, err := c.Get(urlTail)
	if err != nil {
		return pa, err
	}
	err = json.NewDecoder(res.Body).Decode(&pa)
	if err != nil {
		return pa, err
	}
	return pa, nil
}

// Retrieves the currently authenticated user. Returns a user object and a
// 200 OK response code.
func (c *Client) GetMe() (User, error) {
	u := User{}
	urlTail := "/v2/users/me"
	res, err := c.Get(urlTail)
	if err != nil {
		return u, err
	}
	err = json.NewDecoder(res.Body).Decode(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}
