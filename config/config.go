package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	AppEnv string
	BaseURL string

	DatabaseURL string

	PaystackSecretKey string
	PaystackPublicKey string

	BankName          string
	BankAccountName   string
	BankAccountNumber string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, reading from environment")
	}

	cfg := &Config{
		Port:        getEnv("PORT", "8080"),
		AppEnv:      getEnv("APP_ENV", "development"),
		BaseURL:     getEnv("BASE_URL", "http://localhost:8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),

		PaystackSecretKey: getEnv("PAYSTACK_SECRET_KEY", ""),
		PaystackPublicKey: getEnv("PAYSTACK_PUBLIC_KEY", ""),

		BankName:          getEnv("BANK_NAME", ""),
		BankAccountName:   getEnv("BANK_ACCOUNT_NAME", ""),
		BankAccountNumber: getEnv("BANK_ACCOUNT_NUMBER", ""),
	}

	if cfg.PaystackSecretKey == "" {
		log.Println("WARNING: PAYSTACK_SECRET_KEY is not set — donations will fail")
	}

	if cfg.DatabaseURL == "" {
		log.Println("WARNING: DATABASE_URL is not set — database connection will fail")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}