package auth

import (
	"log"
	"os"

	"github.com/BassicallyDrums/filmbase/api/models"
	"github.com/joho/godotenv"
)

func Authenticate() models.Auth {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("ERROR loading .env: %s", err)
	}

	var (
		APIKey   string = os.Getenv("APIKey")
		APIToken string = os.Getenv("APIToken")
	)

	auth := models.Auth{
		APIKey:   APIKey,
		APIToken: APIToken,
	}

	return auth
}
