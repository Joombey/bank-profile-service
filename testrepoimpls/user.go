package testrepoimpls

import (
	"farukh.go/profile/dao/db"
)

type TestUserRepositoryImpl struct {
	dbMock []pair
}

type pair struct {
	name       string
	cardNumber int
}

func (repo TestUserRepositoryImpl) New() *TestUserRepositoryImpl {
	list := make([]pair, 0)
	return &TestUserRepositoryImpl{dbMock: list}
}

func (repo *TestUserRepositoryImpl) CreateUser(name string, cardNumber int) <-chan db.UserTable {
	userChan := make(chan db.UserTable)
	repo.dbMock = append(repo.dbMock, pair{name: name, cardNumber: cardNumber})
	go func() {
		userChan <- db.UserTable{
			Id:         len(repo.dbMock) - 1,
			Name:       name,
			CardNumber: cardNumber,
		}
	}()
	return userChan
}

func (repo *TestUserRepositoryImpl) GetUserById(id int) <-chan db.UserTable {
	userChan := make(chan db.UserTable)

	go func() {
		defer close(userChan)
		userChan <- db.UserTable{
			Id:         id,
			Name:       repo.dbMock[id].name,
			CardNumber: repo.dbMock[id].cardNumber,
		}
	}()

	return userChan
}
