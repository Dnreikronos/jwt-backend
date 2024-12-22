package main

import (
	"fmt"
	"net/http"

	"github.com/Dnreikronos/jwt-backend/configs"
	"github.com/Dnreikronos/jwt-backend/db/connection"
	"github.com/gin-gonic/gin"
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

	db, err := connection.OpenConnection()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
