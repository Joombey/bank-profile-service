package main

import (
	"fmt"
	"testing"

	hr "farukh.go/profile/http/handlers"
)

func TestGetCredentials(t *testing.T) {
	user := hr.CreateUserHandler("TEST")

	id := user.Id
	card := user.CardNumber
	result := hr.GetCredentialsHandler(fmt.Sprintf("%d", id))

	if result.Id != id || result.CardNumber != card || result.Name != "TEST" {
		t.Errorf("wrong vse expected %d %d %s got %d %d %s ", id, card, "TEST", result.Id, result.CardNumber, result.Name)
	}
}
