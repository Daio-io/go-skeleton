package controllers

import (
	"github.com/gin-gonic/gin"
)

type Something struct {
  // The struct keys are tagged with json
  // because go uses uppercase to determine if it is public
  // adding the json tag allows you to tag what the response key
  // will be - e.g. remove uppercase 
  Some string `json:"some"`
  Thing int `json:"thing"`
}

// GetSomething handler
func GetSomething(c *gin.Context) {
	result := Something{Some:"something", Thing:2}
		c.JSON(200, gin.H{"status": "success", "response": &result})
}
