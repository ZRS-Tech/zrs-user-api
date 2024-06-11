package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"zrs.user.api/firebase"
)

// Connection mongoDB with helper class
// var collection = helper.ConnectDB()

// var client *mongo.Client
func main() {
	//initialize firbase
	firebase.InitFirebase()
	// Create a Gin router
	router := gin.Default()
	// helper.ConnectDB()

	// Define a route and a handler
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	// Route for getting a list of items
	router.GET("/items", func(c *gin.Context) {
		items := []string{"Item 1", "Item 2", "Item 3"}
		c.JSON(http.StatusOK, gin.H{
			"items": items,
		})
	})

	// Route for creating a new item
	router.POST("/items", func(c *gin.Context) {
		var newItem struct {
			Name string `json:"name" binding:"required"`
		}

		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// In a real application, you would save the item to a database
		c.JSON(http.StatusCreated, gin.H{
			"message": "Item created",
			"item":    newItem,
		})
	})
	// Start the server
	router.Run(":8080")
}
