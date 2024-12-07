package main

import (
	"flag"
	"log"
	"messages/messages/controllers"
	"messages/messages/initializers"
	"messages/messages/websocket"
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
	flag.Parse()
	log.SetFlags(0)
	hub := websocket.NewHub()
	go hub.Run()
	r.LoadHTMLFiles("home.html")
	//r.Use(middleware.Validate())
	r.GET("/ws", func(c *gin.Context) {
		websocket.ServeWs(hub, c)
	})
	r.GET("/", websocket.ServeHome)
	r.GET("/conversations", controllers.GetChat)
	r.GET("/chats/messages/:id", controllers.GetAllMessages)
	r.GET("/chats", controllers.GetUserChats)

	log.Fatal(r.Run(*addr))
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
