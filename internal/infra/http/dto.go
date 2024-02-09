package http

type TransactionInput struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

type TransactionResponse struct {
	Balance int `json:"saldo"`
	Limit   int `json:"limite"`
}
