package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpalvesfernandes/rinha-backend-2024-q1/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetDB()
}

func (c *NewTransaction) ValidateTransaction() error {
	if c.Type != "d" && c.Type != "c" {
		return fmt.Errorf("tipo must be 'd' or 'c'")
	}
	if len(c.Description) == 0 || len(c.Description) > 10 {
		return fmt.Errorf("descricao must have between 1 to 10 characters")
	}
	return nil
}

func HandleJSONBindingError(c *gin.Context, err error) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
}
