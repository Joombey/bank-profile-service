package handlers

import (
	"strconv"

	"farukh.go/profile/dao/db"
	"farukh.go/profile/di"
	"farukh.go/profile/models"
)

func GetCredentialsHandler(pathVar string) db.UserTable {
	userRepo := di.GetContainer().UserRepository

	id, _ := strconv.Atoi(pathVar)
	return <-userRepo.GetUserById(id)
}

func SendMoneyHandler(sendbody models.TransferDTO) []models.ValueResponse {
	bank := di.GetContainer().Bank

	return <-bank.Transfer(sendbody.From, sendbody.To, sendbody.Value)
}

func CreateUserHandler(pathVar string) db.UserTable{
	bank := di.GetContainer().Bank
	userRepo := di.GetContainer().UserRepository
	cardNumber := (<-bank.NewCard()).CardNumber
	return <-userRepo.CreateUser(pathVar, cardNumber)
}

func Delete(id int) db.UserTable {
	user := <-di.GetContainer().UserRepository.GetUserById(id)
	go di.GetContainer().UserRepository.Delete(id)
	go di.GetContainer().Bank.Delete(user.CardNumber)
	return user
}