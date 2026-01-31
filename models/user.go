package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Phone    string `json:"phone" gorm:"unique"`
	Password string `json:"password"`
	Type     int    `json:"type"` // 0 = regular user, 1 = admin
}
