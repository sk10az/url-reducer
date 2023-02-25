package internal

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("reducer.db"), &gorm.Config{})
	if err != nil {
		return db, err
	}
	db.AutoMigrate(&URL{})
	return db, err
}
