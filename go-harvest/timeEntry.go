package goharvest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"
)

// A response object from requesting time entries
type TimeEntryResponse struct {
	TimeEntries []TimeEntry `json:"time_entries"`
	Pagination
}

// A time entry
type TimeEntry struct {
	// Unique ID for the time entry. Listed as 'bigint' in documentation
	ID uint64 `json:"id"`

	// Date of the time entry.
	SpentDate Date `json:"spent_date"`

	// An object containing the id and name of the associated user.
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"user"`

	// A user assignment object of the associated user.
	UserAssignment UserAssignment `json:"user_assignment"`

	// An object containing the id and name of the associated client.
	Client struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"client"`

	// An object containing the id and name of the associated project.
	Project struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"project"`

	// An object containing the id and name of the associated task.
	Task struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"task"`

	// A task assignment object of the associated task.
	TaskAssignment TaskAssignment `json:"task_assignment"`

	// An object containing the id, group_id, account_id, permalink,
	// service, and service_icon_url of the associated external reference.
	ExternalReference ExternalReference `json:"external_reference"`

	// Once the time entry has been invoiced, this field will include the
	// associated invoice’s id and number.
	Invoice struct {
		ID     int    `json:"id"`
		Number string `json:"number"`
	} `json:"invoice"`

	// Number of (decimal time) hours tracked in this time entry.
	Hours float64 `json:"hours"`

	// Number of (decimal time) hours already tracked in this time entry,
	// before the timer was last started.
	HoursWithoutTimer float64 `json:"hours_without_timer"`

	// Number of (decimal time) hours tracked in this time entry used in
	// summary reports and invoices. This value is rounded according to the
	// Time Rounding setting in your Preferences.
	RoundedHours float64 `json:"rounded_hours"`

	// Notes attached to the time entry.
	Notes string `json:"notes"`

	// Whether or not the time entry has been locked.
	IsLocked bool `json:"is_locked"`

	// Why the time entry has been locked.
	LockedReason string `json:"locked_reason"`

	// Whether or not the time entry has been approved via
	// Timesheet Approval.
	IsClosed bool `json:"is_closed"`

	// Whether or not the time entry has been marked as invoiced.
	IsBilled bool `json:"is_billed"`

	// Date and time the running timer was started (if tracking by duration).
	// Use the ISO 8601 Format. Returns null for stopped timers.
	TimerStartedAt time.Time `json:"timer_started_at"`

	// Time the time entry was started (if tracking by start/end times).
	StartedTime *KitchenTime `json:"started_time"`

	// Time the time entry was ended (if tracking by start/end times).
	EndedTime *KitchenTime `json:"ended_time"`

	// Whether or not the time entry is currently running.
	IsRunning bool `json:"is_running"`

	// Whether or not the time entry is billable.
	Billable bool `json:"billable"`

	// Whether or not the time entry counts towards the project budget.
	Budgeted bool `json:"budgeted"`

	// The billable rate for the time entry.
	BillableRate float64 `json:"billable_rate"`

	// The cost rate for the time entry.
	CostRate float64 `json:"cost_rate"`

	// Date and time the time entry was created. Use the ISO 8601 Format.
	CreatedAt time.Time `json:"created_at"`

	// Date and time the time entry was last updated. Use the ISO
	// 8601 Format.
	UpdatedAt time.Time `json:"updated_at"`
}

type GetTimeEntryParameters struct {
	// Only return time entries belonging to the user with the given ID.
	UserID int `json:"user_id" url:"user_id,omitempty"`

	// Only return time entries belonging to the client with the given ID.
	ClientID int `json:"client_id" url:"client_id,omitempty"`

	// Only return time entries belonging to the project with the given ID.
	ProjectID int `json:"project_id" url:"project_id,omitempty"`

	// Only return time entries belonging to the task with the given ID.
	TaskID int `json:"task_id" url:"task_id,omitempty"`

	// Only return time entries with the given external_reference ID.
	ExternalReferenceID string `json:"external_reference_id" url:"external_reference_id,omitempty"`

	// Pass true to only return time entries that have been invoiced and
	// false to return time entries that have not been invoiced.
	IsBilled bool `json:"is_billed" url:"is_billed,omitempty"`

	// Pass true to only return running time entries and false to return non-
	// running time entries.
	IsRunning bool `json:"is_running" url:"is_running,omitempty"`

	// Only return time entries that have been updated since the given date
	// and time. Use the ISO 8601 Format.
	UpdatedSince time.Time `json:"updated_since" url:"updated_since,omitempty"`

	// Only return time entries with a spent_date on or after the given date.
	From Date `json:"from" url:"from,omitempty"`

	// Only return time entries with a spent_date on or before the
	// given date.
	To Date `json:"to" url:"to,omitempty"`

	// The page number to use in pagination. For instance, if you make a list
	// request and receive 2000 records, your subsequent call can include
	// page=2 to retrieve the next page of the list. (Default: 1)
	Page int `json:"page" url:"page,omitempty"`

	// The number of records to return per page. Can range between 1 and
	// 2000. (Default: 2000)
	PerPage int `json:"per_page" url:"per_page,omitempty"`
}

type CreateTimeEntryBody interface {
	GetTimeEntryBodyParams() string
	IsValid() bool
}

// The body required to create a time entry using the start and end time.
// Note that this same struct and endpoint is also used to start a timer.
// If an EndedTime is not provided, then the timer will be running. If a
// StartedTime is not provided, then the start time of the timer will be
// the current time.
//
// Also note that if your company settings are such that
// users are required to enter time as a duration rather than using start
// and end times, then this endpoint will only be used to start a timer,
// regardless of the values of StartedTime and EndedTime.
type CreateTimeEntryBodyStartEnd struct {
	// The ID of the user to associate with the time entry. Defaults to the
	// currently authenticated user’s ID. - optional
	UserID *int `json:"user_id,omitempty" url:"user_id,omitempty"`

	// The ID of the project to associate with the time entry. - required
	ProjectID int `json:"project_id" url:"project_id,omitempty"`

	// The ID of the task to associate with the time entry. - required
	TaskID int `json:"task_id" url:"task_id,omitempty"`

	// The ISO 8601 formatted date the time entry was spent. - required
	SpentDate Date `json:"spent_date" url:"spent_date,omitempty"`

	// The time the entry started. Defaults to the current time.
	// Example: “8:00am”. - optional
	StartedTime *KitchenTime `json:"started_time,omitempty" url:"started_time,omitempty"`

	// The time the entry ended. If provided, is_running will be set to
	// false. If not provided, is_running will be set to true. - optional
	EndedTime *KitchenTime `json:"ended_time,omitempty" url:"ended_time,omitempty"`

	// Any notes to be associated with the time entry. - optional
	Notes string `json:"notes,omitempty" url:"notes,omitempty"`

	// An object containing the id, group_id, account_id, and permalink of
	// the external reference. - optional
	ExternalReference *ExternalReference `json:"external_reference,omitempty" url:"external_reference,omitempty"`
}

func (b CreateTimeEntryBodyStartEnd) GetTimeEntryBodyParams() string {
	return fmt.Sprintf("%+v", b)
}
func (b CreateTimeEntryBodyStartEnd) IsValid() bool {
	return b.ProjectID != 0 && b.TaskID != 0 && !b.SpentDate.IsZero()
}

// The body required to create a time entry using the duration of an entry.
// If a duration is not provided and the company settings are such that it
// accepts time entries as durations, then an entry with a duration of
// "0.0" will be created.
//
// If, on the other hand, your company settings are such that users enter
// time with start and end times, then this will start a timer the begins
// at the current time.
type CreateTimeEntryBodyDuration struct {
	// The ID of the user to associate with the time entry. Defaults to the
	// currently authenticated user’s ID. - optional
	UserID *int `json:"user_id" url:"user_id,omitempty"`

	// The ID of the project to associate with the time entry. - required
	ProjectID int `json:"project_id" url:"project_id,omitempty"`

	// The ID of the task to associate with the time entry. - required
	TaskID int `json:"task_id" url:"task_id,omitempty"`

	// The ISO 8601 formatted date the time entry was spent. - required
	SpentDate Date `json:"spent_date" url:"spent_date,omitempty"`

	// The current amount of time tracked. If provided, the time entry will
	// be created with the specified hours and is_running will be set to
	// false. If not provided, hours will be set to 0.0 and is_running will
	// be set to true. - optional
	Hours *float64 `json:"hours" url:"hours,omitempty"`

	// Any notes to be associated with the time entry. - optional
	Notes string `json:"notes" url:"notes,omitempty"`

	// An object containing the id, group_id, account_id, and permalink of
	// the external reference. - optional
	ExternalReference *ExternalReference `json:"external_reference" url:"external_reference,omitempty"`
}

func (b CreateTimeEntryBodyDuration) GetTimeEntryBodyParams() string {
	return fmt.Sprintf("%+v", b)
}
func (b CreateTimeEntryBodyDuration) IsValid() bool {
	return b.ProjectID != 0 && b.TaskID != 0 && !b.SpentDate.IsZero()
}

// Returns a list of time entries. The time entries are returned sorted by
// spent_date date. At this time, the sort option can’t be customized.
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

// Retrieves the time entry with the given ID. Returns a time entry object
// and a 200 OK response code if a valid identifier was provided.
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

// Creates a new time entry object. Returns a time entry object and a
// 201 Created response code if the call succeeded.
//
// This same function is used to create time entries when an account is
// configured to track time via duration or via start and end time.
// Harvest will create the time entry accordingly depending on the account
// settings. You can verify this by visiting the Settings page in your
// Harvest account or by checking if wants_timestamp_timers is false in the
// Company API.
//
// If you provide the wrong type of body to this function - i.e., for the
// type opposite the type expected according to the settings - then this
// endpoint will create a running timer.
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
