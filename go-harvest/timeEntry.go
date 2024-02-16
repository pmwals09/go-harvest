package goharvest

import (
	"time"
)

type TimeEntryResponse struct {
  TimeEntries []TimeEntry `json:"time_entries"`
  Pagination
}

type ExternalReference struct {
	ID             int    `json:"id"`
	GroupID        int    `json:"group_id"`
	AccountID      int    `json:"account_id"`
	Permalink      string `json:"permalink"`
	Service        string `json:"service"`
	ServiceIconURL string `json:"service_icon_url"`
}

type SpentDate struct {
  time.Time
}

func (s *SpentDate) UnmarshalJSON(input []byte) error {
  newTime, err := time.Parse(time.DateOnly, string(input[1:len(input)-1]))
  if err != nil {
    s.Time = time.Time{}
    return err
  }

  s.Time = newTime
  return nil
}

type TimeEntry struct {
	ID        uint64    `json:"id"`         // Unique ID for the time entry. Listed as 'bigint' in documentation
	SpentDate SpentDate `json:"spent_date"` // Date of the time entry.
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
		ID     int `json:"id"`
		Number string `json:"number"`
	} `json:"invoice"` // Once the time entry has been invoiced, this field will include the associated invoice’s id and number.
	Hours             float32   `json:"hours"`               // Number of (decimal time) hours tracked in this time entry.
	HoursWithoutTimer float32   `json:"hours_without_timer"` // Number of (decimal time) hours already tracked in this time entry, before the timer was last started.
	RoundedHours      float32   `json:"rounded_hours"`       // Number of (decimal time) hours tracked in this time entry used in summary reports and invoices. This value is rounded according to the Time Rounding setting in your Preferences.
	Notes             string    `json:"notes"`               // Notes attached to the time entry.
	IsLocked          bool      `json:"is_locked"`           // Whether or not the time entry has been locked.
	LockedReason      string    `json:"locked_reason"`       // Why the time entry has been locked.
	IsClosed          bool      `json:"is_closed"`           // Whether or not the time entry has been approved via Timesheet Approval.
	IsBilled          bool      `json:"is_billed"`           // Whether or not the time entry has been marked as invoiced.
	TimerStartedAt    time.Time `json:"timer_started_at"`    // Date and time the running timer was started (if tracking by duration). Use the ISO 8601 Format. Returns null for stopped timers.
	StartedTime       time.Time `json:"started_time"`        // Time the time entry was started (if tracking by start/end times).
	EndedTime         time.Time `json:"ended_time"`          // Time the time entry was ended (if tracking by start/end times).
	IsRunning         bool      `json:"is_running"`          // Whether or not the time entry is currently running.
	Billable          bool      `json:"billable"`            // Whether or not the time entry is billable.
	Budgeted          bool      `json:"budgeted"`            // Whether or not the time entry counts towards the project budget.
	BillableRate      float32   `json:"billable_rate"`       // The billable rate for the time entry.
	CostRate          float32   `json:"cost_rate"`           // The cost rate for the time entry.
	CreatedAt         time.Time `json:"created_at"`          // Date and time the time entry was created. Use the ISO 8601 Format.
	UpdatedAt         time.Time `json:"updated_at"`          // Date and time the time entry was last updated. Use the ISO 8601 Format.
}
