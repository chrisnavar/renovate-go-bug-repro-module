package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func greet() string {
	return "Hello, Renovate!"
}

func main() {
	// Initialize Viper for configuration (example usage)
	viper.SetDefault("app.name", "RenovateGoApp")
	viper.SetDefault("app.version", "1.0.0")

	fmt.Printf("%s version %s is starting...\n", viper.GetString("app.name"), viper.GetString("app.version"))

	// Initialize Gin for a simple web server (example usage)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": greet(),
		})
	})

	log.Println("Server is running on :8080")
	// Start the server (will block until terminated)
	// For a real application, you'd want proper error handling here
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

