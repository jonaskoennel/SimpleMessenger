package controllers

import (
	"messages/messages/initializers"
	"messages/messages/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOwnUserProfile(c *gin.Context) {
	userid, _ := c.Get("userId")
	userProfile := models.ProfileAPI{}
	result := initializers.DB.Model(&models.UserProfile{}).First(&userProfile, userid)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"profile": userProfile,
	})

}

func CreateUserProfile(c *gin.Context) {
	userid, _ := c.Get("userId")

	profile := models.UserProfile{UserID: userid.(uint)}
	if err := c.ShouldBindUri(&profile); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	result := initializers.DB.Create(&profile)
	if result.Error != nil {
		c.JSON(400, gin.H{"msg": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"firsname": profile.Firstname, "lastname": profile.Lastname})
}
