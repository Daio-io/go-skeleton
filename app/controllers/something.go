package controllers

import (
	"github.com/gin-gonic/gin"
  "go-skeleton/app/models"
)

// GetSomething handler
func GetSomething(c *gin.Context) {
	result := models.SomethingModel{Some:"something", Thing:2}
		c.JSON(200, gin.H{"status": "success", "response": &result})
}
