package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	items := []string{"Item 1", "Item 2", "Item 3"}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func CreateItem(c *gin.Context) {
	var newItem struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Item created",
		"item":    newItem,
	})
}
