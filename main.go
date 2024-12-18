package main

import (
	"Go-Gin/controllers"
	"Go-Gin/initializers"
	"Go-Gin/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.Migrate()
}

func main() {
	fmt.Println("Hello Gin...")
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	// User routes
	r.POST("/signup", controllers.SingUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate",middleware.RequireAuth ,controllers.Validate)

	r.Run()
}
