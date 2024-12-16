package controllers

import (
	"fmt"
	"messages/messages/initializers"
	"messages/messages/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllMessages(c *gin.Context) {
	var message models.Message
	if err := c.ShouldBindUri(&message); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	fmt.Println(message.ID)
	var messages []models.MessageAPI
	result := initializers.DB.Model(&message).Where("chat_id = ?", 1).Find(&messages).Order("updated_at")
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
