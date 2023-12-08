package http

import (
	"github.com/gin-gonic/gin"
	hr "farukh.go/profile/http/handlers"
)

func Run() {
	router := gin.Default()
	router.GET("/credentials/:id", hr.CreateUserHandler)
	router.GET("/create/:name", hr.CreateUserHandler)
	router.POST("/send", hr.SendMoneyHandler)
	router.Run("localhost:8080")
}