package handler

type NewTransaction struct {
	Amount      int    `json:"valor"`
	Description string `json:"descricao"`
	Type        string `json:"tipo"`
	AccountID   uint
}
