package controllers

import (
	"fmt"
	"messages/messages/initializers"
	"messages/messages/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetConversations(c *gin.Context) {
	id, _ := c.Get("userId")
	fmt.Printf("UserId in GetConversations: %x\n", id.(int))
	var conv models.Conversation
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
