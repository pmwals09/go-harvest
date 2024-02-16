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
  user, err := client.GetMe()
  if err != nil {
    fmt.Fprintf(os.Stderr, "A problem occurred getting the current user: %w", err)
  }
  fmt.Printf("%+v\n", user)
}
