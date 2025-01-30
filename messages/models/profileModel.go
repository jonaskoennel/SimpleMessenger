package models

import (
	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	UserID    uint
	Firstname string
	Lastname  string
}

type ProfileAPI struct {
	gorm.Model
	ID        uint   `uri:"id" binding:"required,uuid" json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
