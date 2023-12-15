package config

import (
	"fmt"
	"os"

	"github.com/chaiyapatoam/go-clean-fiber-boiler/domain"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func LoadEnv() *domain.Env {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Print("Error loading .env file")
		panic(err)
	}

	GOOGLE_CLIENTID := os.Getenv("GOOGLE_CLIENTID")
	GOOGLE_CLIENDSECRET := os.Getenv("GOOGLE_CLIENTSECRET")
	GOOGLE_REDIRECT := os.Getenv("GOOGLE_REDIRECT")
	MYSQL_URI := os.Getenv("MYSQL_URI")
	PORT := os.Getenv("PORT")
	SESSION_PREFIX := os.Getenv("SESSION_PREFIX")
	SESSION_SECRET := os.Getenv("SESSION_SECRET")
	HASH_SECRET := os.Getenv("HASH_SECRET")
	FRONTEND_URL := os.Getenv("FRONTEND_URL")

	env := &domain.Env{
		MYSQL_URI:           MYSQL_URI,
		GOOGLE_CLIENTID:     GOOGLE_CLIENTID,
		GOOGLE_CLIENTSECRET: GOOGLE_CLIENDSECRET,
		GOOGLE_REDIRECT:     GOOGLE_REDIRECT,
		PORT:                PORT,
		SESSION_PREFIX:      SESSION_PREFIX,
		SESSION_SECRET:      SESSION_SECRET,
		HASH_SECRET:         HASH_SECRET,
		FRONTEND_URL:        FRONTEND_URL,
	}

	validate := validator.New()

	if err := validate.Struct(env); err != nil {
		panic(err)
	}
	return env
}
