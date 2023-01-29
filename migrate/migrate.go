package main

import (
	"fmt"
	"log"

	"github.com/arifrachman98/go-gorm-api/initializers"
	"github.com/arifrachman98/go-gorm-api/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? can not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration Complete")
}
