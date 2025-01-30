package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
func Validate(c *gin.Context) {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	type Token struct {
		Sub int `json:"sub"`
	}
	var token Token
	url := "http://localhost:8080/validate"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: cookie})
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.Set("userId", token.Sub)
	c.Next()

}
*/

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		type Token struct {
			Sub uint `json:"sub"`
		}
		var token Token
		url := "http://localhost:8080/validate"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		req.AddCookie(&http.Cookie{Name: "Authorization", Value: cookie})
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		err = json.NewDecoder(resp.Body).Decode(&token)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		fmt.Printf("Validation Middleware got UserId:%d", token.Sub)
		c.Set("userId", token.Sub)
		c.Next()
	}
}
