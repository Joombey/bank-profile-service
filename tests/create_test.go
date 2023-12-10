package main

import (
	"fmt"
	"testing"

	hr "farukh.go/profile/http/handlers"
)

func TestCreateUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		user := hr.CreateUserHandler(fmt.Sprintf("%d", i))
		if user.Id != i || user.Name != fmt.Sprintf("%d", i) {
			t.Errorf("Wrong Creation expected %d, %s got %d, %s", i, fmt.Sprintf("%d", i), user.Id, user.Name)
		}
	}
}
