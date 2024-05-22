package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alanpramil7/go-time-tracker/db"
	"github.com/alanpramil7/go-time-tracker/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Parse command line arguments
	dateStr := flag.String("date", "", "Date of task in DD-MM-YYYY format")
	timeStr := flag.String("time", "", "Time of task in HH:MM format")
	task := flag.String("task", "", "Name of the task")
	flag.Parse()

	// Initialize db connections
	database, err := db.InitializeDB()
	if err != nil {
		log.Fatalf("Error connecting database: %v", err)
	}

	data := &db.Tasks{
		Date: *dateStr,
		Time: *timeStr,
		Task: *task,
	}

	if *dateStr != "" && *timeStr != "" && *task != "" {
		// Store task in db
		err := db.CreateTasks(database, data)
		if err != nil {
			log.Fatalf("Error creating database entry: %v", err)
		}
	} else {
		// Get all tasks
		tasks, err := db.GetTasks(database)
		if err != nil {
			log.Fatalf("Error getting data: %v", err)
		}

    // Display all tasks in table format
    m := utils.CreateTable(tasks)
    if _, err := tea.NewProgram(m).Run(); err != nil {
      fmt.Println("Error displaying table", err)
      os.Exit(1)
    }
	}
}
