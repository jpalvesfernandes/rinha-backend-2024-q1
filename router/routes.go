package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jpalvesfernandes/rinha-backend-2024-q1/handler"
)

func initRoutes(r *gin.Engine) {
	handler.InitHandler()
	r.POST("/clientes/:id/transacoes", handler.CreateTransaction)
	r.GET("/clientes/:id/extrato", handler.GetStatement)
}
