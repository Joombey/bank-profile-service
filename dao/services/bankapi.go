package services

import (
	"bytes"
	bin "encoding/binary"
	json "encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	cts "farukh.go/profile/constants"
	"farukh.go/profile/models"
)

type BankCommunicator struct {
	client http.Client
}

func (bank BankCommunicator) New() *BankCommunicator {
	client := http.Client{Timeout: time.Duration(10) * time.Second}
	return &BankCommunicator{
		client: client,
	}
}

func (bank BankCommunicator) Transfer(from int, to int, value float32) <-chan []models.ValueResponse {
	channel := make(chan []models.ValueResponse)
	go func() {
		model := models.TransferDTO{
			From:  from,
			To:    to,
			Value: value,
		}
		jsonData, _ := json.MarshalIndent(&model, "", " ")
		response, _ := bank.client.Post(cts.TransferMoney, "application/json", bytes.NewBuffer(jsonData))
		responseModels := make([]models.ValueResponse, 0, 2)
		json.NewDecoder(response.Body).Decode(&responseModels)
		channel <- responseModels
		close(channel)
	}()
	return channel
}

func (bank BankCommunicator) GetValue(cardNumber int) <-chan models.ValueResponse {
	route := fmt.Sprintf("%s/%d", cts.GetValue, cardNumber)
	channel := make(chan models.ValueResponse)
	go func() {
		response, err := bank.client.Get(route)
		if err != nil {
			log.Panicf("error gettings value from bank %s", err.Error())
		}
		defer response.Body.Close()

		var model models.ValueResponse
		json.NewDecoder(response.Body).Decode(&model)

		channel <- model

		close(channel)
	}()
	return channel
}

func (bank *BankCommunicator) NewCard() <-chan models.ValueResponse {
	channel := make(chan models.ValueResponse)
	go func() {
		fmt.Print(bank.client)
		var model models.ValueResponse
		response, _ := bank.client.Get(cts.CreateCard)
		json.NewDecoder(response.Body).Decode(&model)
		channel <- model
		close(channel)
	}()
	return channel
}

func Float64frombytes(bytes []byte) float64 {
	bits := bin.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float32frombytes(bytes []byte) float32 {
	bits := bin.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}
