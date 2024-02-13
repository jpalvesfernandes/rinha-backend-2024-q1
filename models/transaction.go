package models

import (
	"time"
)

type Transaction struct {
	ID          uint
	Amount      int
	Type        string
	Description string
	CreatedAt   time.Time
	AccountID   uint `gorm:"index"`
}
