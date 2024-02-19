package goharvest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"
)

type TimeEntryResponse struct {
	TimeEntries []TimeEntry `json:"time_entries"`
	Pagination
}

type TimeEntry struct {
	ID        uint64 `json:"id"`         // Unique ID for the time entry. Listed as 'bigint' in documentation
	SpentDate Date   `json:"spent_date"` // Date of the time entry.
	User      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"user"` // An object containing the id and name of the associated user.
	UserAssignment UserAssignment `json:"user_assignment"` // A user assignment object of the associated user.
	Client         struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"client"` // An object containing the id and name of the associated client.
	Project struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"project"` // An object containing the id and name of the associated project.
	Task struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"task"` // An object containing the id and name of the associated task.
	TaskAssignment    TaskAssignment    `json:"task_assignment"`    // A task assignment object of the associated task.
	ExternalReference ExternalReference `json:"external_reference"` // An object containing the id, group_id, account_id, permalink, service, and service_icon_url of the associated external reference.
	Invoice           struct {
		ID     int    `json:"id"`
		Number string `json:"number"`
	} `json:"invoice"` // Once the time entry has been invoiced, this field will include the associated invoice’s id and number.
	Hours             float64      `json:"hours"`               // Number of (decimal time) hours tracked in this time entry.
	HoursWithoutTimer float64      `json:"hours_without_timer"` // Number of (decimal time) hours already tracked in this time entry, before the timer was last started.
	RoundedHours      float64      `json:"rounded_hours"`       // Number of (decimal time) hours tracked in this time entry used in summary reports and invoices. This value is rounded according to the Time Rounding setting in your Preferences.
	Notes             string       `json:"notes"`               // Notes attached to the time entry.
	IsLocked          bool         `json:"is_locked"`           // Whether or not the time entry has been locked.
	LockedReason      string       `json:"locked_reason"`       // Why the time entry has been locked.
	IsClosed          bool         `json:"is_closed"`           // Whether or not the time entry has been approved via Timesheet Approval.
	IsBilled          bool         `json:"is_billed"`           // Whether or not the time entry has been marked as invoiced.
	TimerStartedAt    time.Time    `json:"timer_started_at"`    // Date and time the running timer was started (if tracking by duration). Use the ISO 8601 Format. Returns null for stopped timers.
	StartedTime       *KitchenTime `json:"started_time"`        // Time the time entry was started (if tracking by start/end times).
	EndedTime         *KitchenTime `json:"ended_time"`          // Time the time entry was ended (if tracking by start/end times).
	IsRunning         bool         `json:"is_running"`          // Whether or not the time entry is currently running.
	Billable          bool         `json:"billable"`            // Whether or not the time entry is billable.
	Budgeted          bool         `json:"budgeted"`            // Whether or not the time entry counts towards the project budget.
	BillableRate      float64      `json:"billable_rate"`       // The billable rate for the time entry.
	CostRate          float64      `json:"cost_rate"`           // The cost rate for the time entry.
	CreatedAt         time.Time    `json:"created_at"`          // Date and time the time entry was created. Use the ISO 8601 Format.
	UpdatedAt         time.Time    `json:"updated_at"`          // Date and time the time entry was last updated. Use the ISO 8601 Format.
}

type GetTimeEntryParameters struct {
	UserID              int       `json:"user_id" url:"user_id,omitempty"`                             // Only return time entries belonging to the user with the given ID.
	ClientID            int       `json:"client_id" url:"client_id,omitempty"`                         // Only return time entries belonging to the client with the given ID.
	ProjectID           int       `json:"project_id" url:"project_id,omitempty"`                       // Only return time entries belonging to the project with the given ID.
	TaskID              int       `json:"task_id" url:"task_id,omitempty"`                             // Only return time entries belonging to the task with the given ID.
	ExternalReferenceID string    `json:"external_reference_id" url:"external_reference_id,omitempty"` // Only return time entries with the given external_reference ID.
	IsBilled            bool      `json:"is_billed" url:"is_billed,omitempty"`                         // Pass true to only return time entries that have been invoiced and false to return time entries that have not been invoiced.
	IsRunning           bool      `json:"is_running" url:"is_running,omitempty"`                       // Pass true to only return running time entries and false to return non-running time entries.
	UpdatedSince        time.Time `json:"updated_since" url:"updated_since,omitempty"`                 // Only return time entries that have been updated since the given date and time. Use the ISO 8601 Format.
	From                Date      `json:"from" url:"from,omitempty"`                                   // Only return time entries with a spent_date on or after the given date.
	To                  Date      `json:"to" url:"to,omitempty"`                                       // Only return time entries with a spent_date on or before the given date.
	Page                int       `json:"page" url:"page,omitempty"`                                   // The page number to use in pagination. For instance, if you make a list request and receive 2000 records, your subsequent call can include page=2 to retrieve the next page of the list. (Default: 1)
	PerPage             int       `json:"per_page" url:"per_page,omitempty"`                           // The number of records to return per page. Can range between 1 and 2000. (Default: 2000)
}

type CreateTimeEntryBody interface {
	GetTimeEntryBodyParams() string
	IsValid() bool
}

type CreateTimeEntryBodyStartEnd struct {
	UserID            *int               `json:"user_id,omitempty" url:"user_id,omitempty"`                       // The ID of the user to associate with the time entry. Defaults to the currently authenticated user’s ID. - optional
	ProjectID         int                `json:"project_id" url:"project_id,omitempty"`                           // The ID of the project to associate with the time entry. - required
	TaskID            int                `json:"task_id" url:"task_id,omitempty"`                                 // The ID of the task to associate with the time entry. - required
	SpentDate         Date               `json:"spent_date" url:"spent_date,omitempty"`                           // The ISO 8601 formatted date the time entry was spent. - required
	StartedTime       *KitchenTime       `json:"started_time,omitempty" url:"started_time,omitempty"`             // The time the entry started. Defaults to the current time. Example: “8:00am”. - optional
	EndedTime         *KitchenTime       `json:"ended_time,omitempty" url:"ended_time,omitempty"`                 // The time the entry ended. If provided, is_running will be set to false. If not provided, is_running will be set to true. - optional
	Notes             string             `json:"notes,omitempty" url:"notes,omitempty"`                           // Any notes to be associated with the time entry. - optional
	ExternalReference *ExternalReference `json:"external_reference,omitempty" url:"external_reference,omitempty"` // An object containing the id, group_id, account_id, and permalink of the external reference. - optional
}

func (b CreateTimeEntryBodyStartEnd) GetTimeEntryBodyParams() string {
	return fmt.Sprintf("%+v", b)
}
func (b CreateTimeEntryBodyStartEnd) IsValid() bool {
	return b.ProjectID != 0 && b.TaskID != 0 && !b.SpentDate.IsZero()
}

type CreateTimeEntryBodyDuration struct {
	UserID            *int               `json:"user_id" url:"user_id,omitempty"`                       // The ID of the user to associate with the time entry. Defaults to the currently authenticated user’s ID. - optional
	ProjectID         int                `json:"project_id" url:"project_id,omitempty"`                 // The ID of the project to associate with the time entry. - required
	TaskID            int                `json:"task_id" url:"task_id,omitempty"`                       // The ID of the task to associate with the time entry. - required
	SpentDate         Date               `json:"spent_date" url:"spent_date,omitempty"`                 // The ISO 8601 formatted date the time entry was spent. - required
	Hours             *float64           `json:"hours" url:"hours,omitempty"`                           // The current amount of time tracked. If provided, the time entry will be created with the specified hours and is_running will be set to false. If not provided, hours will be set to 0.0 and is_running will be set to true. - optional
	Notes             string             `json:"notes" url:"notes,omitempty"`                           // Any notes to be associated with the time entry. - optional
	ExternalReference *ExternalReference `json:"external_reference" url:"external_reference,omitempty"` // An object containing the id, group_id, account_id, and permalink of the external reference. - optional
}

func (b CreateTimeEntryBodyDuration) GetTimeEntryBodyParams() string {
	return fmt.Sprintf("%+v", b)
}
func (b CreateTimeEntryBodyDuration) IsValid() bool {
	return b.ProjectID != 0 && b.TaskID != 0 && !b.SpentDate.IsZero()
}

func (c *Client) GetTimeEntries(params GetTimeEntryParameters) (TimeEntryResponse, error) {
	tr := TimeEntryResponse{}
	urlTail, err := buildPathWithParams[GetTimeEntryParameters]("/v2/time_entries", params)
	if err != nil {
		return tr, err
	}
	res, err := c.Get(urlTail)
	if err != nil {
		return tr, err
	}
	err = json.NewDecoder(res.Body).Decode(&tr)
	if err != nil {
		return tr, err
	}
	return tr, nil
}

func (c *Client) GetTimeEntry(id uint64) (TimeEntry, error) {
	te := TimeEntry{}
	urlTail := fmt.Sprintf("/v2/time_entries/%d", id)
	res, err := c.Get(urlTail)
	if err != nil {
		return te, err
	}
	err = json.NewDecoder(res.Body).Decode(&te)
	if err != nil {
		return te, err
	}
	return te, nil
}

func (c *Client) CreateTimeEntry(body CreateTimeEntryBody) (TimeEntry, error) {
	te := TimeEntry{}
	if body.IsValid() {
		urlTail := "/v2/time_entries"
		fmt.Printf("%+v\n", body)
		res, err := c.Post(urlTail, body)
		ba, _ := io.ReadAll(res.Body)
		newResBa := bytes.NewBuffer(ba)
		newBody := io.NopCloser(newResBa)
		res.Body = newBody
		fmt.Println(string(ba))
		if err != nil {
			return te, err
		}
		err = json.NewDecoder(res.Body).Decode(&te)
		if err != nil {
			return te, err
		}
		return te, nil
	} else {
		return te, errors.New("Invalid body")
	}
}
