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
	
	MYSQL_URI := os.Getenv("MYSQL_URI")

	env := &domain.Env{
		MYSQL_URI:           MYSQL_URI,
	}

	validate := validator.New()

	if err := validate.Struct(env); err != nil {
		panic(err)
	}
	return env
}

