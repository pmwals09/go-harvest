package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pmwals09/go-harvest/go-harvest"
)

type KT struct {
	time.Time
}

func (k KT) MarshalJSON() ([]byte, error) {
	str := k.Format(time.Kitchen)
	return []byte(`"` + str + `"`), nil
}

type Obj struct {
	Key KT `json:"key"`
}

func main() {
	// TODO: Should be in .env or by flag, flag takes precedence
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env")
	}
	PAT := os.Getenv("HARVEST_PAT")
	accountID := os.Getenv("HARVEST_ACCOUNT_ID")

	client := goharvest.NewClient(PAT, accountID, "pmwals09@gmail.com")

	fmt.Println("CURRENT USER")
	user, err := client.GetMe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem getting current user: %s", err.Error())
	}
	fmt.Printf("%+v\n\n", user)

	fmt.Println("PROJECT ASSIGNMENTS")
	projectAssignmentsResponse, err := client.GetMyProjectAssignments(
		goharvest.GetProjectAssignmentParameters{})
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Problem getting current user project assignments: %s", err.Error())
	}
	for _, pa := range projectAssignmentsResponse.ProjectAssignments {
		code, name, clientName := pa.Project.Code, pa.Project.Name, pa.Client.Name
		fmt.Println(code, "-", clientName, "-", name)
		for _, task := range pa.TaskAssignments {
			fmt.Println(task.Task.Name)
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("TIME ENTRIES")
	timeEntriesResponse, err := client.GetTimeEntries(
		goharvest.GetTimeEntryParameters{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem getting time entries: %s", err.Error())
	}

	fmt.Println("TIME ENTRY")
	timeEntry, err := client.GetTimeEntry(timeEntriesResponse.TimeEntries[0].ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem getting time entry: %s", err.Error())
	}
	fmt.Printf("%+v\n\n", timeEntry)

	fmt.Println("POST TIME ENTRY")
	project := projectAssignmentsResponse.ProjectAssignments[0]
	startTime := time.Date(2024, time.February, 17, 11, 0, 0, 0, time.Now().Location())
	duration := 1.0
	timeEntryPost, err := client.CreateTimeEntry(
		goharvest.CreateTimeEntryBodyDuration{
			UserID:            &user.ID,
			ProjectID:         project.Project.ID,
			TaskID:            project.TaskAssignments[0].Task.ID,
			SpentDate:         goharvest.Date{Time: startTime},
			Hours:             &duration,
			Notes:             "This is a test",
			ExternalReference: nil,
		})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem creating time entry: %s", err.Error())
	}
	fmt.Printf("%+v\n\n", timeEntryPost)
}
