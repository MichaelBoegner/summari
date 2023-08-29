package main

import (
	"emails/grab"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	grab.GetEmails()
	dotenv := goDotEnvVariable("OPENAI_API_KEY")
	fmt.Printf("env value %s", dotenv)
}
