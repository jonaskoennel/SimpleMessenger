package main

import (
	"authentication/auth/controllers"
	"authentication/auth/initializers"
	"authentication/auth/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}
func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/user/username", middleware.RequireAuth, controllers.GetUserByUsername)
	r.Run()

}
