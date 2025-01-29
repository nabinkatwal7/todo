package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title string `gorm:"not null"`
	Completed bool `gorm:"default:false"`
	DueDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}