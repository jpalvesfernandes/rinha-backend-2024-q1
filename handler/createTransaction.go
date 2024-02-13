package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jpalvesfernandes/rinha-backend-2024-q1/models"
	"gorm.io/gorm/clause"
)

func CreateTransaction(c *gin.Context) {
	id := c.Param("id")
	maxRetries := 5
	var retryCount int

	for retryCount < maxRetries {
		tx := db.Begin()

		a := &models.Account{}
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&a, id).Error; err != nil {
			tx.Rollback()
			logger.Warn("Account not found:", "id", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
			return
		}

		newTransaction := NewTransaction{}
		if err := c.ShouldBindJSON(&newTransaction); err != nil {
			tx.Rollback()
			logger.Error("Error binding json:", err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		if err := newTransaction.ValidateTransaction(); err != nil {
			tx.Rollback()
			logger.Error("Error validating transaction:", err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		if newTransaction.Type == "d" {
			a.Balance -= newTransaction.Amount
		} else {
			a.Balance += newTransaction.Amount
		}

		if a.Balance < a.AccountLimit*-1 {
			tx.Rollback()
			logger.Warn("Transaction rejected: Insufficient balance", "id:", a.ID)
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Insufficient balance"})
			return
		}

		addTransaction := models.Transaction{
			Amount:      newTransaction.Amount,
			Type:        newTransaction.Type,
			Description: newTransaction.Description,
			CreatedAt:   time.Now(),
			AccountID:   a.ID,
		}

		a.Transactions = append(a.Transactions, addTransaction)

		if err := tx.Save(&a).Error; err != nil {
			tx.Rollback()
			logger.Error("Error saving transaction:", err)
			retryCount++
			continue
		}

		tx.Commit()

		logger.Info("Transaction created:", "id:", a.ID, "amount:", newTransaction.Amount, "type:", newTransaction.Type, "description:", newTransaction.Description)

		response := TransactionResponse{AccountLimit: a.AccountLimit, Balance: a.Balance}
		c.JSON(http.StatusOK, response)
		return
	}

	if retryCount >= maxRetries {
		logger.Error("Transaction conflict detected: ", "id:", id)
		c.JSON(http.StatusConflict, gin.H{"error": "Transaction conflict, please retry"})
	}
}
