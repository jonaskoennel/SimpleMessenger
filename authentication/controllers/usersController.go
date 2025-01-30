package controllers

import (
	"authentication/auth/initializers"
	"authentication/auth/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{Email: body.Email, Username: body.Username, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	fmt.Println(c.Cookie("Authorization"))
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	var ret []byte
	c.Request.Body.Read(ret)
	fmt.Printf("Request Body: %s\n", string(ret))

	fmt.Printf("Login: {Username: %s, Password: %s}\n", body.Email, body.Password)
	var user models.User
	err := initializers.DB.First(&user, "email = ?", body.Email).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("AUTH_JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"Cookie": tokenString,
	})

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	fmt.Println(user.(models.User).Email)
	c.JSON(http.StatusOK, gin.H{
		"sub": user.(models.User).ID,
	})
}

func GetUserByUsername(c *gin.Context) {
	var body struct {
		Username string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	fmt.Printf("Username: %s\n", body.Username)
	var ret struct {
		ID uint
	}
	err := initializers.DB.Model(&models.User{}).Where("username LIKE = ?", body.Username).First(&ret).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find user in db",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userId": ret.ID,
	})

}
