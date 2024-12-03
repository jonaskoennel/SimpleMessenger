package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	TextMessage string
	UserID      string
	//Conversation Conversation
}
