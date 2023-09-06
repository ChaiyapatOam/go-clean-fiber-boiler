package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Print("Error loading .env file")
	}

	value := os.Getenv(key)

	if value == "" {
		fmt.Println("Error Get env Key:", key)
		panic(err)
	}

	return value

}
