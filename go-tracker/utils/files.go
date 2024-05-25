package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/alanpramil7/go-time-tracker/db"
	"gorm.io/gorm"
)

const (
	basePath = "/home/alan/Downloads/"
)

func CreateFile(filename string) (filepath string) {
	fullPath := path.Join(basePath, filename+".csv")
	if _, err := os.Stat(fullPath); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(fullPath)
		if err != nil {
			log.Fatalf("Error on creating file: %v", err)
		}
		file.Close()
	} else {
		log.Fatal("File name is already present")
	}

	return fullPath
}

func WriteFile(dbatabase *gorm.DB, fullPath string) {
	tasks, err := db.GetTasks(dbatabase)
	if err != nil {
		log.Panic(err)
	}

	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	header := []byte("\"ID\",\"Date\",\"Time\",\"Task\"\n")
	_, err = file.Write(header)
	if err != nil {
		log.Panic(err)
	}

	for _, task := range tasks {
		fmt.Println(task)
		data := fmt.Sprintf("\"%d\",\"%s\",\"%s\",\"%s\"\n", task.ID, task.Date, task.Time, task.Task)
		_, err = file.Write([]byte(data))
		if err != nil {
			log.Panic(err)
		}
	}
}
