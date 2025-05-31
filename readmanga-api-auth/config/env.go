package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Println("Erro ao carregar .env (modo dev)")
		}
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
