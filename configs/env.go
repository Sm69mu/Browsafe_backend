package configs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var Port string

func LoadEnvVariables() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory")
	}
	envPaths := []string{
		filepath.Join(cwd, ".env"),
		filepath.Join(cwd, "../.env"),
	}
	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			break
		}
	}
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080" // Default port
	}
	fmt.Printf("Server running on port: %s\n", Port)
}
