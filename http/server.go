package http

import (
	"github.com/gin-gonic/gin"
	hr "farukh.go/profile/http/handlers"
)

func run() {
	router := gin.Default()
	router.GET("/credentials/:id", hr.CreateUserHandler)
	router.Run("localhost:8080")
}