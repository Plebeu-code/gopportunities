package main

import (
	"github.com/Plebeu-code/gopportunities/config"
	"github.com/Plebeu-code/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Config
	err := config.Init()
	if err != nil {
		// panic(err)
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	//Initialize Router
	router.Initialize()
}
