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

	"farukh.go/profile/models"
)

const (
	valueRoute string = "localhost:8081/value"
	baseApi    string = "http://localhost:8081"
)

type BankCommunicator struct {
	Client *http.Client
}

func (bank BankCommunicator) New() *BankCommunicator {
	client := http.Client{Timeout: time.Duration(10) * time.Second}
	return &BankCommunicator{
		Client: &client,
	}
}

func (bank BankCommunicator) Transfer(from int, to int, value float32) <-chan []models.ValueResponse {
	route := fmt.Sprintf("%s/transfer", baseApi)
	channel := make(chan []models.ValueResponse)
	go func() {
		model := models.TransferDTO{
			From:  from,
			To:    to,
			Value: value,
		}
		jsonData, _ := json.MarshalIndent(&model, "", " ")
		response, _ := bank.Client.Post(route, "application/json", bytes.NewBuffer(jsonData))
		responseModels := make([]models.ValueResponse, 0, 2)
		json.NewDecoder(response.Body).Decode(&responseModels)
		channel <- responseModels
		close(channel)
	}()
	return channel
}

func (bank BankCommunicator) GetValue(cardNumber int) <-chan models.ValueResponse {
	route := fmt.Sprintf("%s/%d", baseApi, cardNumber)
	channel := make(chan models.ValueResponse)
	go func() {
		response, err := bank.Client.Get(route)
		if err != nil {
			log.Panicf("error gettings value from bank", err)
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
	route := fmt.Sprintf("%s/new-card", baseApi)
	channel := make(chan models.ValueResponse)
	go func() {
		var model models.ValueResponse
		response, _ := bank.Client.Get(route)
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
