package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar .env")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
