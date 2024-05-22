package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error){
  dsn := "host=localhost user=postgres password=postgres dbname=go_tracker port=5432 sslmode=disable"
  db ,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, err
  }

  db.AutoMigrate(&Tasks{})

  return db, nil
}
