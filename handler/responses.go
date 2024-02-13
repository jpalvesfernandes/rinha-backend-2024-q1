package handler

import "time"

type Transaction struct {
	Amount      int    `json:"valor"`
	Description string `json:"descricao"`
	Type        string `json:"tipo"`
}

type TransactionResponse struct {
	AccountLimit int `json:"limite"`
	Balance      int `json:"saldo"`
}

type BankStatement struct {
	Balance          BalanceDetails `json:"saldo"`
	LastTransactions []Transactions `json:"ultimas_transacoes"`
}

type BalanceDetails struct {
	Balance      int       `json:"total"`
	Date         time.Time `json:"data_extrato"`
	AccountLimit int       `json:"limite"`
}

type Transactions struct {
	Amount      int       `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizada_em"`
}
