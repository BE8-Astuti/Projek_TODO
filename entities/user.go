package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string   `json:"username" gorm:"unique"`
	Name     string   `json:"name"`
	Email    string   `json:"email" gorm:"unique"`
	Password string   `json:"password" form:"password"`
	Phone    string   `json:"phone" gorm:"unique"`
	Gender   string   `json:"gender"`
	Projek   []Projek `gorm:"foreignKey:UserID"`
	Task     []Task   `gorm:"foreignKey:UserID"`
}
