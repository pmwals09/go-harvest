# Go Harvest

This package provides a Go library that can interact with the Harvest API.
The progress toward 100% compatibility can be found below.

## Usage

The intended usage of this client is as follows:

```go
package main

import (
    harvest "github.com/pmwals09/go-harvest"
)

func main() {
    client := harvest.NewClient()
    timeEntries := client.ReadTimeEntries()
    newTimeEntry := harvest.TimeEntry{}
    client.CreateTimeEntry(newTimeEntry)
}
```

Most of the endpoints in question require authentication.
Harvest allows for two types of authentication - the only one supported here is thePersonal Access Token (PAT).
You must have a valid `HARVEST_PAT` environment variable set in order to use this library.

## FAQ

### Do I really need to provide an email or contact link?

Yes, you really do, according to [Harvest's documentation](https://help.getharvest.com/api-v2/introduction/overview/general/#api-requests). You can't use mine, sorry.


## API Coverage

- [Authentication](#Authentication)
- [Clients API](#Clients_API)
- [Company Settings](#Company_Settings)
- [Invoices API](#Invoices_API)
- [Estimates API](#Estimates_API)
- [Expenses API](#Expenses_API)
- [Tasks API](#Tasks_API)
- [Timesheets API](#Timesheets_API)
- [Projects API](#Projects_API)
- [Roles API](#Roles_API)
- [Users API](#Users_API)
- [Reports API](#Reports_API)

### Authentication

Documentation: [Authentication](https://help.getharvest.com/api-v2/authentication-api/authentication/authentication/)

- [ ] Personal Access Tokens
- [ ] OAuth2

### Clients API

Documentation:
- [Client Contacts](https://help.getharvest.com/api-v2/clients-api/clients/contacts/)
- [Clients](https://help.getharvest.com/api-v2/clients-api/clients/clients/)

- [ ] GET /v2/contacts
- [ ] GET /v2/contacts/{CONTACT_ID}
- [ ] POST /v2/contacts
- [ ] PATCH /v2/contacts/{CONTACT_ID}
- [ ] DELETE /v2/contacts/{CONTACT_ID}
- [ ] GET /v2/clients
- [ ] GET /v2/clients/{CLIENT_ID}
- [ ] POST /v2/clients
- [ ] PATCH /v2/clients/{CLIENT_ID}
- [ ] DELETE /v2/clients/{CLIENT_ID}

### Company Settings

Documentation: [Company](https://help.getharvest.com/api-v2/company-api/company/company/)

- [ ] GET /v2/company
- [ ] PATCH /v2/company

### Invoices API

Documentation:
- [Invoice Messages](https://help.getharvest.com/api-v2/invoices-api/invoices/invoice-messages/)
- [Invoice Payments](https://help.getharvest.com/api-v2/invoices-api/invoices/invoice-payments/)
- [Invoices](https://help.getharvest.com/api-v2/invoices-api/invoices/invoices/)
- [Invoice Item Categories](https://help.getharvest.com/api-v2/invoices-api/invoices/invoice-item-categories/)

- [ ] GET /v2/invoices/{INVOICE_ID}/messages
- [ ] POST /v2/invoices/{INVOICE_ID}/messages
- [ ] GET /v2/invoices/{INVOICE_ID}/messages/new
- [ ] DELETE /v2/invoices/{INVOICE_ID}/messages/{message_ID}
- [ ] GET /v2/invoices/{INVOICE_ID}/payments
- [ ] POST /v2/invoices/{INVOICE_ID}/payments
- [ ] DELETE /v2/invoices/{INVOICE_ID}/payments/{PAYMENT_ID}
- [ ] GET /v2/invoices
- [ ] GET /v2/invoices/{INVOICE_ID}
- [ ] POST /v2/invoices
- [ ] PATCH /v2/invoices/{INVOICE_ID}
- [ ] DELETE /v2/invoices/{INVOICE_ID}
- [ ] GET /v2/invoice_item_categories
- [ ] GET /v2/invoice_item_categories/{INVOICE_ITEM_CATEGORY_ID}
- [ ] POST /v2/invoice_item_categories
- [ ] PATCH /v2/invoice_item_categories/{INVOICE_ITEM_CATEGORY_ID}
- [ ] DELETE /v2/invoice_item_categories/{INVOICE_ITEM_CATEGORY_ID}

### Estimates API

Documentation:
- [Estimate Messages](https://help.getharvest.com/api-v2/estimates-api/estimates/estimate-messages/)
- [Estimates](https://help.getharvest.com/api-v2/estimates-api/estimates/estimates/)
- [Estimate Item Categories](https://help.getharvest.com/api-v2/estimates-api/estimates/estimate-item-categories/)

- [ ] GET /v2/estimates/{estimate_ID}/messages
- [ ] POST /v2/estimates/{estimate_ID}/messages
- [ ] DELETE /v2/estimates/{estimate_ID}/messages/{message_ID}
- [ ] GET /v2/estimates
- [ ] GET /v2/estimates/{ESTIMATE_ID}
- [ ] POST /v2/estimates
- [ ] PATCH /v2/estimates/{ESTIMATE_ID}
- [ ] DELETE /v2/estimates/{ESTIMATE_ID}
- [ ] GET /v2/estimate_item_categories
- [ ] GET /v2/estimate_item_categories/{ESTIMATE_ITEM_CATEGORY_ID}
- [ ] POST /v2/estimate_item_categories
- [ ] PATCH /v2/estimate_item_categories/{ESTIMATE_ITEM_CATEGORY_ID}
- [ ] DELETE /v2/estimate_item_categories/{ESTIMATE_ITEM_CATEGORY_ID}

### Expenses API

Documentation:
- [Expenses](https://help.getharvest.com/api-v2/expenses-api/expenses/expenses/)
- [Expense Categories](https://help.getharvest.com/api-v2/expenses-api/expenses/expense-categories/)

- [ ] GET /v2/expenses
- [ ] GET /v2/expenses/{EXPENSE_ID}
- [ ] POST /v2/expenses
- [ ] PATCH /v2/expenses/{EXPENSE_ID}
- [ ] DELETE /v2/expenses/{EXPENSE_ID}
- [ ] GET /v2/expense_categories
- [ ] GET /v2/expense_categories/{EXPENSE_CATEGORY_ID}
- [ ] POST /v2/expense_categories
- [ ] PATCH /v2/expense_categories/{EXPENSE_CATEGORY_ID}
- [ ] DELETE /v2/expense_categories/{EXPENSE_CATEGORY_ID}

### Tasks API

Documentation: [Tasks](https://help.getharvest.com/api-v2/tasks-api/tasks/tasks/)

- [ ] GET /v2/tasks
- [ ] GET /v2/tasks/{TASK_ID}
- [ ] POST /v2/tasks
- [ ] PATCH /v2/tasks/{TASK_ID}
- [ ] DELETE /v2/tasks/{TASK_ID}

### Timesheets API

Documentation: [Time Entries](https://help.getharvest.com/api-v2/timesheets-api/timesheets/time-entries/)

- [x] GET /v2/time_entries
- [x] GET /v2/time_entries/{TIME_ENTRY_ID}
- [x] POST /v2/time_entries
- [ ] PATCH /v2/time_entries/{TIME_ENTRY_ID}
- [ ] DELETE /v2/time_entries/{TIME_ENTRY_ID}/external_reference
- [ ] DELETE /v2/time_entries/{TIME_ENTRY_ID}
- [ ] PATCH /v2/time_entries/{TIME_ENTRY_ID}/restart
- [ ] PATCH /v2/time_entries/{TIME_ENTRY_ID}/stop

### Projects API

Documentation:
- [Project User Assignments](https://help.getharvest.com/api-v2/projects-api/projects/user-assignments/)
- [Project Task Assignments](https://help.getharvest.com/api-v2/projects-api/projects/task-assignments/)
- [Projects](https://help.getharvest.com/api-v2/projects-api/projects/projects/)

- [x] GET /v2/user_assignments
- [ ] GET /v2/projects/{PROJECT_ID}/user_assignments
- [ ] GET /v2/projects/{PROJECT_ID}/user_assignments/{USER_ASSIGNMENT_ID}
- [ ] POST /v2/projects/{PROJECT_ID}/user_assignments
- [ ] PATCH /v2/projects/{PROJECT_ID}/user_assignments/{USER_ASSIGNMENT_ID}
- [ ] DELETE /v2/projects/{PROJECT_ID}/user_assignments/{USER_ASSIGNMENT_ID}
- [ ] GET /v2/task_assignments
- [ ] GET /v2/projects/{PROJECT_ID}/task_assignments
- [ ] GET /v2/projects/{PROJECT_ID}/task_assignments/{TASK_ASSIGNMENT_ID}
- [ ] POST /v2/projects/{PROJECT_ID}/task_assignments
- [ ] PATCH /v2/projects/{PROJECT_ID}/task_assignments/{TASK_ASSIGNMENT_ID}
- [ ] DELETE /v2/projects/{PROJECT_ID}/task_assignments/{TASK_ASSIGNMENT_ID}
- [ ] GET /v2/projects
- [ ] GET /v2/projects/{PROJECT_ID}
- [ ] POST /v2/projects
- [ ] PATCH /v2/projects/{PROJECT_ID}
- [ ] DELETE /v2/projects/{PROJECT_ID}

### Roles API

Documentation: [Roles](https://help.getharvest.com/api-v2/roles-api/roles/roles/)

- [ ] GET /v2/roles
- [ ] GET /v2/roles/{ROLE_ID}
- [ ] POST /v2/roles
- [ ] PATCH /v2/roles/{ROLE_ID}
- [ ] DELETE /v2/roles/{ROLE_ID}

### Users API

Documentation:
- [User Teammates](https://help.getharvest.com/api-v2/users-api/users/teammates/)
- [User Billable Rates](https://help.getharvest.com/api-v2/users-api/users/billable-rates/)
- [User Cost Rates](https://help.getharvest.com/api-v2/users-api/users/cost-rates/)
- [User Project Assignments](https://help.getharvest.com/api-v2/users-api/users/project-assignments/)
- [Users](https://help.getharvest.com/api-v2/users-api/users/users/)

- [-] GET /v2/users/{USER_ID}/teammates
- [-] PATCH /v2/users/{USER_ID}/teammates
- [-] GET /v2/users/{USER_ID}/billable_rates
- [-] GET /v2/users/{USER_ID}/billable_rates/{billable_RATE_ID}
- [-] POST /v2/users/{USER_ID}/billable_rates
- [-] GET /v2/users/{USER_ID}/cost_rates
- [-] GET /v2/users/{USER_ID}/cost_rates/{COST_RATE_ID}
- [-] POST /v2/users/{USER_ID}/cost_rates
- [-] GET /v2/users/{USER_ID}/project_assignments
- [x] GET /v2/users/me/project_assignments
- [-] GET /v2/users
- [x] GET /v2/users/me
- [-] GET /v2/users/{USER_ID}
- [-] POST /v2/users
- [-] PATCH /v2/users/{USER_ID}
- [-] PATCH /v2/users/{USER_ID}
- [-] DELETE /v2/users/{USER_ID}

### Reports API

Documentation:
- [Expense Reports](https://help.getharvest.com/api-v2/reports-api/reports/expense-reports/)
- [Uninvoiced Report](https://help.getharvest.com/api-v2/reports-api/reports/uninvoiced-report/)
- [Time Reports](https://help.getharvest.com/api-v2/reports-api/reports/time-reports/)
- [Project Budget Report](https://help.getharvest.com/api-v2/reports-api/reports/project-budget-report/)
 
- [ ] GET /v2/reports/expenses/clients
- [ ] GET /v2/reports/expenses/projects
- [ ] GET /v2/reports/expenses/categories
- [ ] GET /v2/reports/expenses/team
- [ ] GET /v2/reports/uninvoiced
- [ ] GET /v2/reports/time/clients
- [ ] GET /v2/reports/time/projects
- [ ] GET /v2/reports/time/tasks
- [ ] GET /v2/reports/time/team
- [ ] GET /v2/reports/project_budget
