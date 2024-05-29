package main

import (
	"maestrore/core"
	"maestrore/domain"

	"github.com/joho/godotenv"
)

func main() {
	error := godotenv.Load()
	if error != nil {
		panic("Error loading .env file")
	}

	config := core.NewConfig()

	apiService := domain.NewAPIService(config)
	apiService.Init()
	apiService.Run()
}
