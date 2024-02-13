package handler

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jpalvesfernandes/rinha-backend-2024-q1/models"
	"gorm.io/gorm"
)

func GetStatement(c *gin.Context) {
	id := c.Param("id")

	var wg sync.WaitGroup

	resultChan := make(chan BankStatement, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()

		var account models.Account
		if err := db.Preload("Transactions", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC").Limit(10)
		}).First(&account, id).Error; err != nil {
			logger.Warn("Account not found:", "id", id)
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("invalid data")})
			return
		}

		bankStatement := BankStatement{
			Balance: BalanceDetails{
				Balance:      account.Balance,
				Date:         time.Now(),
				AccountLimit: account.AccountLimit,
			},
			LastTransactions: make([]Transactions, len(account.Transactions)),
		}

		for i, transaction := range account.Transactions {
			bankStatement.LastTransactions[i] = Transactions{
				Amount:      transaction.Amount,
				Type:        transaction.Type,
				Description: transaction.Description,
				CreatedAt:   transaction.CreatedAt,
			}
		}

		resultChan <- bankStatement
	}()

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	bankStatement, ok := <-resultChan
	if !ok {
		logger.Error("Failed to fetch account details: ", "id", id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch account details"})
		return
	}

	logger.Info("Account details fetched successfully: ", "id", id)
	c.JSON(http.StatusOK, bankStatement)
}
