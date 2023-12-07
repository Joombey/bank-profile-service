package models

type ValueResponse struct {
	CardNumber int     `json:"card_number"`
	Value      float32 `json:"value"`
}

type TransferDTO struct {
	From  int     `json:"from"`
	To    int     `json:"to"`
	Value float32 `json:"value"`
}
