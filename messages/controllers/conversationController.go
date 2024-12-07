package controllers

import (
	"fmt"
	"messages/messages/initializers"
	"messages/messages/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetChat(c *gin.Context) {
	id, _ := c.Get("userId")
	fmt.Printf("UserId in GetConversations: %x\n", id.(int))
	var conv models.Chat
	err := initializers.DB.First(&conv).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"userId": id.(int),
	})
}

func GetUserChats(c *gin.Context) {
	userId, _ := c.Get("userId")
	var chats []models.Chat
	result := initializers.DB.Find(&chats, userId.(int)).Order("updated_at")
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.Bind(&chats)
	c.JSON(http.StatusOK, gin.H{
		"chats": chats,
	})
}
