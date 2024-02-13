package models

type Account struct {
	ID           uint
	AccountLimit int
	Balance      int
	Transactions []Transaction
	Version      int `gorm:"column:version"`
}
