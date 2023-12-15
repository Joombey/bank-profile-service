package di

import (
	"farukh.go/profile/dao/db"
	"farukh.go/profile/dao/services"

	// "farukh.go/profile/internal"
	"farukh.go/profile/repos"
	// testImpls "farukh.go/profile/testrepoimpls"
)

type Uploader interface {
	Upload(num int, value float32) float32
}

type BaseContainer struct {
	Bank           repos.BankRepository
	UserRepository repos.UserRepository
}

func (c *BaseContainer) init() {
	c.Bank = services.BankCommunicator{}.New()
	c.UserRepository = db.UserRepositoryImpl{}.New()
}

func GetUploader() Uploader {
	return uploader
}

func GetContainer() BaseContainer {
	if (container == BaseContainer{}) {
		(&container).init()
	}
	return container
}

var container = BaseContainer{}
var uploader Uploader
