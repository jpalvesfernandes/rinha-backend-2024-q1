package main

import (
	"github.com/jpalvesfernandes/rinha-backend-2024-q1/config"
	"github.com/jpalvesfernandes/rinha-backend-2024-q1/router"
)

var logger *config.Logger

func main() {
	err := config.Init()
	logger = config.GetLogger("main")

	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	logger.Info("Config initialized successfully")
	router.Start()
}
