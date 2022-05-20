package entities

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID   uint   `json:"userid"`
	ProjekID uint   `json:"projekid"`
	Name     string `json:"name"`
	Duedate  string `json:"duedate"`
	Status   string `json:"status"`
}
