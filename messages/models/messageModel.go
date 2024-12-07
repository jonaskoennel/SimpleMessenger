package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID     uint64 `gorm:"primaryKey" uri:"id" binding:"required"`
	Text   string `gorm:"type:text"`
	UserID uint
	ChatID uint64 `gorm:"index"`
	Chat   Chat
}

type MessageAPI struct {
	ID     uint64 `json:"id"`
	ChatID uint64 `json:"chatId"`
	UserID uint   `json:"userId"`
	Text   string `json:"text"`
	//CreatedAt time.Time `json:"created_at"`
}
