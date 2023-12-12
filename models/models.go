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

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type InsertRequest struct {
	CardNumber int     `json:"card_number"`
	Value      float32 `json:"value"`
}