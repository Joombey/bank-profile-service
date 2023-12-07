package repos

import "farukh.go/profile/dao/db"

type UserRepository interface {
	CreateUser(name string, cardNumber int) (<-chan db.UserTable)
	GetUserById(id int) (<-chan db.UserTable)
}