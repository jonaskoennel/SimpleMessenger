package controllers

import (
	"fmt"
	"messages/messages/initializers"
	"messages/messages/models"
	"messages/messages/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllMessages(c *gin.Context) {
	type Query struct {
		ChatId string `uri:"chat_id" binding:"required"`
	}
	var query Query
	if err := c.ShouldBindUri(&query); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	var messages []models.MessageAPI
	result := initializers.DB.Model(&models.Message{}).Where("chat_id=?", query.ChatId).Find(&messages).Order("id")
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}

func CreateMessage(c *gin.Context) {
	//var message models.MessageAPI
	var message models.Message
	err := c.Bind(&message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to map request data on MessageAPI-Model!",
		})
		return
	}
	fmt.Println(message.Text)
	userid := c.MustGet("userId").(uint)
	if userid == 0 {
		fmt.Println("Got no userid from middleware!")
	}

	fmt.Printf("UserId from cookie: %d\n", userid)
	var valid bool
	valid, err = utils.IsChatParticipant(message.ChatId, userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Chat does not exist or you are not part of this chat",
		})
		return
	}
	message.SenderId = userid

	err = utils.CreateNewMessage(&message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "Message is stored in database!",
	})
}
