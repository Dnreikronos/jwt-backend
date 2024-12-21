package main

import (
	"fmt"

	"github.com/Dnreikronos/jwt-backend/configs"
	"github.com/Dnreikronos/jwt-backend/db"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
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

	db, err := db.OpenConnection()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
}
