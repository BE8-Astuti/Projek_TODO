package entities

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID   uint      `json:"userid"`
	ProjekID uint      `json:"projekid"`
	Name     string    `json:"name"`
	Duedate  time.Time `json:"duedate"`
	Status   string    `json:"status"`
}
