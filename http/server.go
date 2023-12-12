package http

import (
	"net/http"

	hr "farukh.go/profile/http/handlers"
	kk "farukh.go/profile/keycloak"
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

	router.GET("/test/login", func(ctx *gin.Context) {
		kk.LoginAdmin()
		ctx.IndentedJSON(http.StatusOK, 1)
	})

	router.GET("/test/create/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		name = kk.CreateUser(name)
		println(name)
		ctx.IndentedJSON(http.StatusOK, 1)
	})

	router.GET("/test/get/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		user := kk.GetUser(name)
		println(*user.Username)
		ctx.IndentedJSON(http.StatusOK, 1)
	})

	router.POST("/test/register", func(ctx *gin.Context) {
		var body models.RegisterRequest
		ctx.BindJSON(&body)

		userId := kk.Register(body)

		ctx.IndentedJSON(http.StatusOK, userId)
	})

	router.POST("/test/auth", func(ctx *gin.Context) {
		var body models.RegisterRequest
		ctx.BindJSON(&body)

		token := kk.Auth(body)

		ctx.IndentedJSON(http.StatusOK, token)
	})

	router.GET("/test/clients", func(ctx *gin.Context) {
		kk.GetClients()
		ctx.IndentedJSON(http.StatusOK, 1)
	})

	router.POST("/test/check", func(ctx *gin.Context) {
		token := token{}
		err := ctx.BindJSON(&token)
		if err != nil {
			panic(err.Error())
		}
		kk.Decode(token.Token)
		ctx.IndentedJSON(http.StatusOK, nil)
	})
	router.POST("/test/rolecheck", func(ctx *gin.Context) {
		var body models.RegisterRequest
		ctx.BindJSON(&body)

		kk.RoleCheck(body)

		ctx.IndentedJSON(http.StatusOK, "")
	})

	router.GET("/test/cliet/:username", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, kk.GetUserByUsername(ctx.Param("username")))
	})

	router.Run()
}

type token struct {
	Token string `json:"token"`
}
