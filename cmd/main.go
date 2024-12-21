package main

import (
	"fmt"

	"github.com/Dnreikronos/jwt-backend/configs"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	err = configs.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}
}
