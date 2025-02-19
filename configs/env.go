package configs

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

var Port string

func LoadEnvVariables() {
	err := godotenv.Load("../.env")
	Port = os.Getenv("PORT")
	fmt.Print(Port)
	if err!=nil{
		log.Fatal("Error loading .env file")
	}
}