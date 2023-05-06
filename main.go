package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	log.Printf("%v", &json)
	println(&json)
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"body": json,
	})
}

func main() {
	r := gin.Default()
	r.POST("/webhook", Login)
	r.Run(":8080")
}
