package repos

import "farukh.go/profile/models"

type BankRepository interface {
	Transfer(from int, to int, value float32) <-chan []models.ValueResponse
	GetValue(cardNumber int) <-chan models.ValueResponse
	NewCard() <-chan models.ValueResponse
}