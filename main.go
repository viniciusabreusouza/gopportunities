package main

import (
	"github.com/viniciusabreusouza/gopportunities/config"
	"github.com/viniciusabreusouza/gopportunities/router"
)

func main() {
	logger := config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		panic(err)
	}

	router.Initialize()
}
