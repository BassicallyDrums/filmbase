package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("ERROR loading .env: %s", err)
	}

	var (
		APIKey   string = os.Getenv("APIKey")
		APIToken string = os.Getenv("APIToken")
	)

}
