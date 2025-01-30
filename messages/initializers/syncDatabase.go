package initializers

import (
	"fmt"
	"log"
	"messages/messages/models"
)

func SyncDatabase() {
	/*
		DB.Migrator().CreateTable(&models.Conversation{})
		DB.Migrator().CreateTable(&models.Message{})
	*/
	DB.AutoMigrate(&models.ChatParticipants{})
	DB.AutoMigrate(&models.Chat{})
	DB.AutoMigrate(&models.Message{})
	err := insertTestData()
	if err != nil {
		log.Println(err)
	}
}

func insertTestData() error {

	messages := []*models.Message{
		{Text: "Das ist eine Test-Nachricht!", SenderId: 1, ChatId: 1},
		{Text: "Wow cool", SenderId: 2, ChatId: 1},
		{Text: "Lalalalala!", SenderId: 1, ChatId: 1},
	}
	participants := []models.ChatParticipants{
		{UserID: 1},
		{UserID: 2},
	}

	conv := models.Chat{Name: "Test", Participants: participants}
	err := DB.Create(&conv).Error
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	err = DB.Create(messages).Error
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
