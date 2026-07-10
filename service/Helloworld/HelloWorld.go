package service

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HelloWorldRouder(api *gin.RouterGroup) {
	api.POST("/hello", HandleHelloworldRequest)
	api.GET("/", HandleHelloworldRequest)
}

func HandleHelloworldRequest(ctx *gin.Context) {
	log.Println("Printing Hello World")
	ctx.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
