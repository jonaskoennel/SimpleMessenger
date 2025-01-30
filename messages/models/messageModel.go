package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	ChatId   uint   `uri:"chat_id" binding:"required"`
	Text     string `gorm:"type:text" json:"text"`
	SenderId uint
}

type MessageAPI struct {
	ID        uint      `json:"id"`
	ChatID    uint      `json:"chatId"`
	SenderID  uint      `json:"senderId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
