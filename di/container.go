package di

import (
	"farukh.go/profile/dao/db"
	"farukh.go/profile/dao/services"
	"farukh.go/profile/models"
)

func init() {
	container.init()
}

type BankRepository interface {
	Transfer(from int, to int, value float32) <-chan []models.ValueResponse
	GetValue(cardNumber int) <-chan models.ValueResponse
	NewCard() <-chan models.ValueResponse
}

type UserRepository interface {
	CreateUser(name string, cardNumber int) <-chan db.UserTable
}

type BaseContainer struct {
	Bank           BankRepository
	UserRepository UserRepository
}

func (c *BaseContainer) init() {
	var bank = services.BankCommunicator{}
	var user = db.UserRepositoryImpl{}
	c = &BaseContainer{
		Bank:           bank.New(),
		UserRepository: user.New(),
	}
}

func GetContainer() *BaseContainer {
	if (container == &BaseContainer{}) {
		container.init()
	}
	return container
}

var container *BaseContainer
