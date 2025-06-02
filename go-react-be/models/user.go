package models

import "gorm.io/gorm"

// User represents the user model in the database
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"` // In a real app, hash this!
}
