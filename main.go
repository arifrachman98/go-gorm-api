package main

import (
	"log"
	"net/http"

	"github.com/arifrachman98/go-gorm-api/initializers"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Can not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Can not load environment variables", err)
	}

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"Status": "success", "message": message})
	})

	log.Fatal(server.Run(":" + config.ServerPort))
}
