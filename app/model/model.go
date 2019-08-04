package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserObj struct {
	Name    string
	Email   string
	Created time.Time
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&UserObj{})
	return db
}
