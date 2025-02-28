package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Ask the user for the migration name
	fmt.Print("Enter the migration name: ")
	var migrationName string
	_, err = fmt.Scanln(&migrationName) // Get input from the user
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Validate migration name
	if migrationName == "" {
		log.Fatal("Migration name cannot be empty!")
	}

	// Run the migration create command with the specified name
	cmd := exec.Command("go", "run", "github.com/golang-migrate/migrate/v4/cmd/migrate", "create", "-ext", "sql", "-dir", "migrations", "-seq", migrationName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error creating migration: %v", err)
	}

	// Print success message
	fmt.Printf("Migration file %s created successfully.\n", migrationName)
}
