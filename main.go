package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/pmwals09/go-harvest/go-harvest"
)

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
	projectAssignmentsResponse, err := client.GetMyProjectAssignments(goharvest.GetProjectAssignmentParameters{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem getting current user project assignments: %s", err.Error())
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
	timeEntriesResponse, err := client.GetTimeEntries(goharvest.GetTimeEntryParameters{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem getting time entries: %s", err.Error())
	}

	fmt.Println("TIME ENTRY")
	timeEntry, err := client.GetTimeEntry(timeEntriesResponse.TimeEntries[0].ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem getting time entry: %s", err.Error())
	}
	fmt.Printf("%+v\n\n", timeEntry)
}
