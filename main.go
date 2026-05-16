package main

import (
	"backend-auth/handlers"
	"backend-auth/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")

	protected.Use(middleware.AuthMiddleware())

	protected.GET("/profile", handlers.GetProfile)
	protected.GET("/users", handlers.GetUsers)

	r.Run(":8080")
}