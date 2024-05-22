package db

import "gorm.io/gorm"

func CreateTasks(db *gorm.DB, data *Tasks) error {
	result := db.Create(&data)
	return result.Error
}

func GetTasks(db *gorm.DB) ([]Tasks, error) {
	var tasks []Tasks
	result := db.Find(&tasks)
	return tasks, result.Error
}
