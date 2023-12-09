package di

import (
	"farukh.go/profile/dao/db"
	"farukh.go/profile/dao/services"
	"farukh.go/profile/repos"
)

type BaseContainer struct {
	Bank           repos.BankRepository
	UserRepository repos.UserRepository
}

func (c *BaseContainer) init() {
	var bank = services.BankCommunicator{}
	var user = db.UserRepositoryImpl{}

	c.Bank = bank.New()
	c.UserRepository = user.New()
}

func GetContainer() BaseContainer {
	if (container == BaseContainer{}) {
		(&container).init()
	}
	return container
}

var container = BaseContainer{}
