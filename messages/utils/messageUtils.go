package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"messages/messages/initializers"
	"messages/messages/models"
	"net/http"
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

func GetUserFromUsername(username string) (uint, error) {
	type Body struct {
		Username string `json:"Username"`
	}

	type Response struct {
		UserId uint `json:"userId"`
	}

	body := &Body{Username: username}
	body1, err := json.Marshal(&body)
	url := "http://localhost:8080/user/username"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body1))
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	defer resp.Body.Close()
	response := &Response{}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Error in response from auth")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(bodyBytes, &response)
	//derr := json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		//fmt.Println(err)
		return 0, err
	}
	fmt.Println(response.UserId)
	return response.UserId, nil
}

func CreateNewChat(creatorId uint, chatName string, invUserName string) (models.Chat, error) {
	invUserId, err := GetUserFromUsername(invUserName)
	if creatorId == invUserId {
		return models.Chat{}, fmt.Errorf("You cannot invite yourself to chat")
	}
	if err != nil {
		return models.Chat{}, err
	}
	part := []models.ChatParticipants{
		{UserID: creatorId},
		{UserID: invUserId},
	}
	chat := models.Chat{Name: chatName, Participants: part, Messages: []models.Message{}}
	err = initializers.DB.Create(&chat).Error
	if err != nil {
		return models.Chat{}, err
	}
	return chat, nil
}
