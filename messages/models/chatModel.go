package models

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey" json:"id" uri:"chatid" binding:"required"`
	Name     string    `json:"name"`
	UsersID  uint64    `gorm:"index" json:"userId"`
	Users    Users     `json:"users"`
	Messages []Message `json:"messages"`
}

type Users struct {
	gorm.Model
	ID    uint `gorm:"primaryKey" json:"id"`
	User1 uint `json:"user1"`
	User2 uint `json:"user2"`
}
