package main

import (
	"github.com/gin-gonic/gin"
	"zrs.user.api/controllers"
	"zrs.user.api/firebase"
	"zrs.user.api/middlewares"
)

// Connection mongoDB with helper class
// var collection = helper.ConnectDB()

// var client *mongo.Client
func main() {
	//initialize firbase
	firebase.InitFirebase()
	middlewares.FirebaseApp = firebase.App
	// Create a Gin router
	router := gin.Default()
	// helper.ConnectDB()

	// Define a route and a handler
	// public route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
	//protected route
	auth := router.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	// auth.GET("/items")
	// Route for getting a list of items
	auth.GET("/items", controllers.GetItems)

	// Route for creating a new item
	auth.POST("/items", controllers.CreateItem)
	// Start the server
	router.Run(":8080")
}
