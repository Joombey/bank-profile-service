package http

import (
	"net/http"
	"strconv"
	"github.com/Depado/ginprom"
	hr "farukh.go/profile/http/handlers"
	"farukh.go/profile/models"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	p := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	router.Use(p.Instrument())
	
	router.GET("/credentials/:id", func(ctx *gin.Context) {
		user := hr.GetCredentialsHandler(ctx.Param("id"))
		ctx.IndentedJSON(http.StatusOK, user)
	})
	router.GET("/create/:name", func(ctx *gin.Context) {
		response := hr.CreateUserHandler(ctx.Param("name"))
		ctx.IndentedJSON(http.StatusOK, response)
	})
	router.POST("/send", func(ctx *gin.Context) {
		var sendbody models.TransferDTO
		ctx.BindJSON(&sendbody)
		response := hr.SendMoneyHandler(sendbody)
		ctx.IndentedJSON(http.StatusOK, response)
	})
	router.GET("/block/:id", func(ctx *gin.Context) {
		idTpDelete, _ := strconv.Atoi(ctx.Param("id"))
		deletedUser := hr.Delete(idTpDelete)
		ctx.IndentedJSON(200, deletedUser)
	})

	router.Run()
}
