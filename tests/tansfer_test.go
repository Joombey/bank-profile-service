package main

import (
	"testing"

	"farukh.go/profile/di"
	hr "farukh.go/profile/http/handlers"
	"farukh.go/profile/models"
)

func TestTransferHandler(t *testing.T) {
	user1 := hr.CreateUserHandler("user1")
	user2 := hr.CreateUserHandler("user2")

	var uploader = di.GetUploader()
	uploader.Upload(user1.CardNumber, 500)

	result := hr.SendMoneyHandler(models.TransferDTO{
		From:  user1.CardNumber,
		To:    user2.CardNumber,
		Value: 250,
	})

	if result[0].Value != 250 || result[1].Value != 250 {
		t.Errorf("Wrong method calling expected %d %d got %f %f", 250, 250, result[0].Value, result[1].Value)
	}
}
