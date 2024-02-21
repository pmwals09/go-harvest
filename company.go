package goharvest

import "encoding/json"

type Company struct {
	// The Harvest URL for the company.
	BaseUri string `json:"base_uri" url:"base_uri,omitempty"`

	// The Harvest domain for the company.
	FullDomain string `json:"full_domain" url:"full_domain,omitempty"`

	// The name of the company.
	Name string `json:"name" url:"name,omitempty"`

	// Whether the company is active or archived.
	IsActive bool `json:"is_active" url:"is_active,omitempty"`

	// The weekday used as the start of the week. Returns one of: Saturday,
	// Sunday, or Monday.
	WeekStartDay string `json:"week_start_day" url:"week_start_day,omitempty"`

	// Whether time is tracked via duration or start and end times.
	WantsTimestampTimers bool `json:"wants_timestamp_timers" url:"wants_timestamp_timers,omitempty"`

	// The format used to display time in Harvest. Returns either decimal or
	// hours_minutes.
	TimeFormat string `json:"time_format" url:"time_format,omitempty"`

	// The format used to display date in Harvest. Returns one of: %m/%d/%Y,
	// %d/%m/%Y, %Y-%m-%d, %d.%m.%Y,.%Y.%m.%d or %Y/%m/%d.
	DateFormat string `json:"date_format" url:"date_format,omitempty"`

	// The type of plan the company is on. Examples: trial, free, or simple-v4
	PlanType string `json:"plan_type" url:"plan_type,omitempty"`

	// Used to represent whether the company is using a 12-hour or 24-hour clock.
	// Returns either 12h or 24h.
	Clock string `json:"clock" url:"clock,omitempty"`

	// How to display the currency code when formatting currency. Returns one of:
	// iso_code_none, iso_code_before, or iso_code_after.
	CurrencyCodeDisplay string `json:"currency_code_display" url:"currency_code_display,omitempty"`

	// How to display the currency symbol when formatting currency. Returns one
	// of: symbol_none, symbol_before, or symbol_after.
	CurrencySymbolDisplay string `json:"currency_symbol_display" url:"currency_symbol_display,omitempty"`

	// Symbol used when formatting decimals.
	DecimalSymbol string `json:"decimal_symbol" url:"decimal_symbol,omitempty"`

	// Separator used when formatting numbers.
	ThousandsSeparator string `json:"thousands_separator" url:"thousands_separator,omitempty"`

	// The color scheme being used in the Harvest web client.
	ColorScheme string `json:"color_scheme" url:"color_scheme,omitempty"`

	// The weekly capacity in seconds.
	WeeklyCapacity int `json:"weekly_capacity" url:"weekly_capacity,omitempty"`

	// Whether the expense module is enabled.
	ExpenseFeature bool `json:"expense_feature" url:"expense_feature,omitempty"`

	// Whether the invoice module is enabled.
	InvoiceFeature bool `json:"invoice_feature" url:"invoice_feature,omitempty"`

	// Whether the estimate module is enabled.
	EstimateFeature bool `json:"estimate_feature" url:"estimate_feature,omitempty"`

	// Whether the approval module is enabled.
	ApprovalFeature bool `json:"approval_feature" url:"approval_feature,omitempty"`
}

// Retrieves the company for the currently authenticated user. Returns a
// company object and a 200 OK response code.
func (c *Client) GetCompany() (Company, error) {
	company := Company{}
	res, err := c.Get("/v2/company")
	if err != nil {
		return company, err
	}

	err = json.NewDecoder(res.Body).Decode(&company)
	if err != nil {
		return company, err
	}
	return company, nil
}

type CompanyUpdateParameters struct {
	// Whether time is tracked via duration or start and end times.
	WantsTimestampTimers *bool `json:"wants_timestamp_timers" url:"wants_timestamp_timers,omitempty"`

	// The weekly capacity in seconds.
	WeeklyCapacity *int `json:"weekly_capacity" url:"weekly_capacity,omitempty"`
}

// Updates the company setting the values of the parameters passed. Any
// parameters not provided will be left unchanged. Returns a company object
// and a 200 OK response code if the call succeeded.
func (c *Client) UpdateCompany(params CompanyUpdateParameters) (Company, error) {
	company := Company{}
	res, err := c.Patch("/v2/company", params)
	if err != nil {
		return company, err
	}
	err = json.NewDecoder(res.Body).Decode(&company)
	if err != nil {
		return company, err
	}
	return company, nil
}
