package models

import (
	"gorm.io/gorm"
)

type Conversation struct {
	gorm.Model
	Name   string
	userID int
}
