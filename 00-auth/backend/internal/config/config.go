package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
 DtatabaseURL string
 JWTSecret string
}

func Load() *Config{
	if err := godotenv.Load(); err != nil {
		log.Fatalf("can't find .env file and load it to os system: %v",err)
	}

	cfg := &Config{
		DtatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if cfg.DtatabaseURL == ""{
		log.Fatal("DATABASE_URL is empty")
	}

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is empty")
	}

	return cfg
}
