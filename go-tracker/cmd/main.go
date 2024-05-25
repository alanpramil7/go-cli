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
	deleteId := flag.Int("delete", 0, "Name of id to be deleted")
	download := flag.String("download", "", "File name to download")
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
	} else if *deleteId > 0 {
		// Delete the task
		err := db.DeleteTask(database, *deleteId)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task deleted sucessfully")
	} else if *download != "" {
		//Download file containing data in CSV format
		fullpath := utils.CreateFile(*download)
		utils.WriteFile(database, fullpath)
    fmt.Println("Data downloaded sucessfully")
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
