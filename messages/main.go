package main

import (
	"flag"
	"log"
	"messages/messages/controllers"
	"messages/messages/initializers"
	"messages/messages/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

var addr = flag.String("addr", ":8090", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/conversations", middleware.Validate, controllers.GetConversations)
	r.Run("localhost:8090")
	/*
		r := gin.Default()
		r.GET("/loadAllMessages", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		r.Run()
		flag.Parse()
		hub := newHub()
		go hub.run()
		http.HandleFunc("/", serveHome)
		http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
			serveWs(hub, w, r)
		})
		err := http.ListenAndServe(*addr, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	*/
}
