package http

import (
	hr "farukh.go/profile/http/handlers"
	"farukh.go/profile/models"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.GET("/credentials/:id", func(ctx *gin.Context) {
		hr.GetCredentialsHandler(ctx.Param(":id"))
	})
	router.GET("/create/:name", func(ctx *gin.Context) {
		hr.CreateUserHandler(ctx.Param(":name"))
	})
	router.POST("/send", func(ctx *gin.Context) {		
		var sendbody models.TransferDTO 
		ctx.BindJSON(&sendbody)
		hr.SendMoneyHandler(sendbody)
	})
	router.Run()
}