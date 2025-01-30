package utils

import (
	"fmt"
	"messages/messages/initializers"
	"messages/messages/models"
)

func CreateNewMessage(message *models.Message) error {
	if message.ChatId == 0 {
		return fmt.Errorf("ChatID cant be 0!")
	}
	if message.Text == "" {
		return fmt.Errorf("Message text must not be empty!")
	}
	if message.SenderId == 0 {
		return fmt.Errorf("UserID cant be 0!")
	}
	fmt.Printf("Message: {ChatID: %d, UserID: %d, Text: %s}\n", message.ChatId, message.SenderId, message.Text)
	/*
		var chat models.Chat
		err := initializers.DB.First(&chat, message.ChatID).Error
		if err != nil {
			return err
		}
	*/
	//err := initializers.DB.Model(&models.Message{}).Create(message).Error
	err := initializers.DB.Create(message).Error
	if err != nil {
		return err
	}
	return nil
}

func IsChatParticipant(chatId uint, userId uint) (bool, error) {
	participants := &models.ChatParticipants{UserID: userId, ChatID: chatId}
	result := initializers.DB.Find(participants)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func GetChatParticipants(chatId uint) ([]models.ChatParticipants, error) {
	participants := []models.ChatParticipants{}
	err := initializers.DB.Where("chat_id=?", chatId).Find(&participants).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(participants)
	return participants, nil

}
