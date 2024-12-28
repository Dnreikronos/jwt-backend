package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Dnreikronos/jwt-backend/configs"
	"github.com/Dnreikronos/jwt-backend/db/connection"
	"github.com/Dnreikronos/jwt-backend/db/migration"
	"github.com/Dnreikronos/jwt-backend/handlers"
	"github.com/gin-contrib/cors"
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

	migration.RunMigrations(db)

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	corsOrigin := os.Getenv("CORS")

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{corsOrigin},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	r.POST("/register", handlers.CreateUserHandler)
	r.POST("/login", handlers.LoginHandler)
	authorized := r.Group("/", handlers.AuthMiddleware())
	{
		authorized.GET("/profile", handlers.ProfileHandler)
	}

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
