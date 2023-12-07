package handlers

import (
	"farukh.go/profile/di"
	"github.com/gin-gonic/gin"
)

var container = di.GetContainer()
var bank = container.Bank
var userRepo = container.UserRepository

func GetCredentialsHandler(c *gin.Context) {
	// TODO: Описать Handler'ы
}

func SendMoneyHandler(c *gin.Context) {
	// TODO: Отправить запрос в bank service
	// получить ответ в канале с
}

func CreateUserHandler(ctx *gin.Context) {
	bankCardChannel := bank.NewCard()
	ctx.BindJSON((<-userRepo.CreateUser(ctx.Copy().Params.ByName("name"), (<-bankCardChannel).CardNumber)))
}
