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
	DB.AutoMigrate(&models.Chat{})
	DB.AutoMigrate(&models.Message{})
	DB.AutoMigrate(&models.Chat{})
	err := insertTestData()
	if err != nil {
		log.Println(err)
	}
}

func insertTestData() error {
	messages := []models.Message{
		{Text: "Das ist eine Test-Nachricht!", UserID: 1},
		{Text: "Wow cool", UserID: 2},
		{Text: "Boah gar kein Bock", UserID: 1},
	}
	conv := models.Chat{Name: "Test", Users: models.Users{User1: 1, User2: 2}, Messages: messages}
	err := DB.Create(&conv).Error
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
