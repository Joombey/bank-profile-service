package handlers

import (
	"net/http"
	"strconv"

	"farukh.go/profile/di"
	"farukh.go/profile/models"
	"github.com/gin-gonic/gin"
)

var container = di.GetContainer()
var bank = container.Bank
var userRepo = container.UserRepository

func GetCredentialsHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := <-userRepo.GetUserById(id)
	println(id == user.Id)
	c.IndentedJSON(http.StatusOK, user)
}

func SendMoneyHandler(c *gin.Context) {
	var sendbody models.TransferDTO
	c.BindJSON(&sendbody)
	response := bank.Transfer(sendbody.From, sendbody.To, sendbody.Value)
	c.IndentedJSON(http.StatusOK, <-response)
}

func CreateUserHandler(c *gin.Context) {
	cardNumber := (<-bank.NewCard()).CardNumber
	userChan := userRepo.CreateUser(c.Copy().Param("name"), cardNumber)
	c.IndentedJSON(http.StatusOK, <-userChan)
}
