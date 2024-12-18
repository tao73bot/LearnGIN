package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	Email    string `gorm:"unique"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Posts    []Post `json:"posts"`
}