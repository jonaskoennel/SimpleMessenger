package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var addr = flag.String("addr", ":8050", "http service address")

func init() {
	//initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	flag.Parse()
	log.SetFlags(0)
	r.LoadHTMLGlob("pages/*")
	r.Static("/scripts", "./scripts")
	r.Static("/style", "./style")
	r.Static("/resources", "./resources")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	log.Fatal(r.Run(*addr))
}
