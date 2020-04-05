package main

import (
	"github.com/gin-gonic/gin"

	"go-layered-architecture-practice/cmd/server"
)

func main() {
	router := gin.Default()

	router.POST("/users", func(c *gin.Context) { server.RegisterUser(c) })

	router.Run()
}
