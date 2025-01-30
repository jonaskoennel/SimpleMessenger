package models

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	ID           uint `gorm:"primaryKey"`
	Name         string
	Participants []ChatParticipants
	Messages     []Message
}

/*
	type User struct {
		gorm.Model
		ID       uint    `gorm:"primaryKey" json:"id"`
		Username string  `json:"username"`
		Chats    []*Chat `gorm:"many2many:chat_participants;"`
	}
*/
type ChatParticipants struct {
	gorm.Model
	ChatID uint
	UserID uint
}

type ChatPreview struct {
	ID           uint   `gorm:"primaryKey" `
	Name         string `json:"name"`
	Participants []ParticipantPreview
}

type ParticipantPreview struct {
	ChatID uint
	UserID uint
}

/*
type ChatParticipants struct {
	gorm.Model
	ChatId uint
	Chat   Chat
	UserID uint
	User   UserProfile
}
*/
