package initializers

import "messages/messages/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Conversation{})
	DB.AutoMigrate(&models.Message{})
}
