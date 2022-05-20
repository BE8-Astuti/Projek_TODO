package entities

import (
	"gorm.io/gorm"
)

type Projek struct {
	gorm.Model
	UserID      uint   `json:"userid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Contributor string `json:"contributor"`
}
