package main

import (
	service "order_management/service/Helloworld"
	"order_management/utility"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	log := utility.InitializeLogger()

	log.Info("Starting the server")
	api := router.Group("/api")

	//enabled the HelloWorld route
	service.HelloWorldRouder(api)

	router.Run(":8080")
}
