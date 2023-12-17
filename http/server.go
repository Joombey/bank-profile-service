package http

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	hr "farukh.go/profile/http/handlers"
	"farukh.go/profile/models"
	"github.com/Depado/ginprom"
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

func RunYan() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err.Error())
	}

	router := gin.Default()
	router.GET("/home", func(ctx *gin.Context) {
		ctx.IndentedJSON(200, "ok")
	})
	router.Run(fmt.Sprintf(":%d", port))
}
