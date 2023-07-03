package main

import (
	"github.com/LagLabs/backend-go.git/config"
	"github.com/LagLabs/backend-go.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config init error: %v", err)
		return
	}
	//Initialize router
	router.Initialize()
}
