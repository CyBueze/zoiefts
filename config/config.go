package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds every runtime setting the app needs, loaded once at
// startup from the environment (and .env in development).
type Config struct {
	Port    string
	AppEnv  string
	BaseURL string

	DBPath string

	PaystackSecretKey string
	PaystackPublicKey string

	BankName          string
	BankAccountName   string
	BankAccountNumber string

	BankName2          string
	BankAccountName2   string
	BankAccountNumber2 string

	ImageBaseURL string
}

func Load() *Config {
	// Ignore the error: in production environment variables are set directly
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, reading from environment")
	}

	cfg := &Config{
		Port:    getEnv("PORT", "8080"),
		AppEnv:  getEnv("APP_ENV", "development"),
		BaseURL: getEnv("BASE_URL", "http://localhost:8080"),

		DBPath: getEnv("DB_PATH", "./data/zoie.db"),

		PaystackSecretKey: getEnv("PAYSTACK_SECRET_KEY", ""),
		PaystackPublicKey: getEnv("PAYSTACK_PUBLIC_KEY", ""),

		BankName:          getEnv("BANK_NAME", ""),
		BankAccountName:   getEnv("BANK_ACCOUNT_NAME", ""),
		BankAccountNumber: getEnv("BANK_ACCOUNT_NUMBER", ""),

		BankName2:          getEnv("BANK_NAME_2", ""),
		BankAccountName2:   getEnv("BANK_ACCOUNT_NAME_2", ""),
		BankAccountNumber2: getEnv("BANK_ACCOUNT_NUMBER_2", ""),

		ImageBaseURL: getEnv("IMAGE_BASE_URL", ""),
	}

	if cfg.PaystackSecretKey == "" {
		log.Println("WARNING: PAYSTACK_SECRET_KEY is not set — donations will fail")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
