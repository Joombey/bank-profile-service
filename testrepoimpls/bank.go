package testrepoimpls

import (
	"farukh.go/profile/models"
)

type BankTestImple struct {
	valueList []float32
}

func (b BankTestImple) New() *BankTestImple {
	list := make([]float32, 0)
	return &BankTestImple{valueList: list}
}

func (b *BankTestImple) Transfer(from int, to int, value float32) <-chan []models.ValueResponse {
	valueChannel := make(chan []models.ValueResponse)
	go func() {
		defer close(valueChannel)
		fromValue := b.valueList[from]
		toValue := b.valueList[to]
		b.valueList[from] = fromValue - value
		b.valueList[to] = toValue + value
		valueChannel <- []models.ValueResponse{
			{
				CardNumber: from,
				Value:      fromValue - value,
			},
			{
				CardNumber: to,
				Value:      toValue + value,
			},
		}
	}()
	return valueChannel
}
func (b *BankTestImple) GetValue(cardNumber int) <-chan models.ValueResponse {
	valueChannel := make(chan models.ValueResponse)
	go func() {
		defer close(valueChannel)
		valueChannel <- models.ValueResponse{
			CardNumber: cardNumber,
			Value:      b.valueList[cardNumber],
		}
	}()
	return valueChannel
}
func (b *BankTestImple) NewCard() <-chan models.ValueResponse {
	valueChannel := make(chan models.ValueResponse)
	go func() {
		defer close(valueChannel)
		b.valueList = append(b.valueList, 0.0)
		valueChannel <- models.ValueResponse{
			CardNumber: len(b.valueList) - 1,
			Value:      0.0,
		}
	}()
	return valueChannel
}

func (b *BankTestImple) Upload(num int, value float32) float32 {
	b.valueList[num] = b.valueList[num] + value
	return b.valueList[num]
}
