package db

type Tasks struct {
	Date string `gorm:"column:date;type:text"`
	Time string `gorm:"column:time;type:text"`
	Task string `gorm:"column:task;type:text"`
}
