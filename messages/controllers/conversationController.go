package controllers

import (
	"fmt"
	"messages/messages/initializers"
	"messages/messages/models"
	"messages/messages/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	userId, _ := c.MustGet("userId").(uint)

	var chats []models.Chat
	result := initializers.DB.Preload("Messages").Find(&chats, userId).Order("created_at")
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

func LoadUserChats(c *gin.Context) {
	var chat []models.Chat
	var part []models.ChatParticipants
	userid := c.MustGet("userId").(uint)
	res := initializers.DB.Where("user_id=?", userid).Find(&part)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": res.Error.Error(),
		})
		return
	}

	var chatIds []uint
	for _, partis := range part {
		chatIds = append(chatIds, partis.ChatID)
	}

	result := initializers.DB.Model(&models.Chat{}).Preload("Participants").Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return db.Order("messages.id DESC")
	}).Find(&chat, "chats.id IN ?", chatIds)

	//result := initializers.DB.Joins("Participants").Find(&chat)

	//result := initializers.DB.Model(&models.Chat{}).Joins("Participants").Find(&chat)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.Bind(&chat)
	c.JSON(http.StatusOK, gin.H{
		"chats": chat,
	})

}

func GetChatPreview(c *gin.Context) {
	//userId, _ := c.Get("userId")

	var chats []models.ChatPreview
	result := initializers.DB.Model(&models.Chat{}).Preload("Messages").Find(&chats, 1).Order("created_at")
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

func CreateChat(c *gin.Context) {
	userid := c.MustGet("userId").(uint)
	type Body struct {
		Name     string
		Username string
	}
	body := Body{}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body does not match",
		})
		return
	}

	chat, derr := utils.CreateNewChat(userid, body.Name, body.Username)
	if derr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": derr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"chat": chat,
	})
}
