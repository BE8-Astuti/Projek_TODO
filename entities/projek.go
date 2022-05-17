package entities

import (
	"time"

	"gorm.io/gorm"
)

type Projek struct {
	gorm.Model
	UserID      uint      `json:"userid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Contributor string    `json:"contributor"`
}
