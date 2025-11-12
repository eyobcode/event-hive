package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type DatabaseConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

type JWTConfig struct {
	Secret string
	Expiry int // minutes
}

type Config struct {
	Database DatabaseConfig
	Port     int
	JWT      JWTConfig
}

func LoadConfig() Config {
	_ = godotenv.Load() // load .env if exists

	return Config{
		Database: DatabaseConfig{
			User: getEnv("DB_USER", "postgres"),
			Pass: getEnv("DB_PASS", ""),
			Host: getEnv("DB_HOST", "localhost"),
			Port: getEnv("DB_PORT", "5432"),
			Name: getEnv("DB_NAME", "event_hive"),
		},
		Port: getEnvInt("SERVER_PORT", 8080),
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "supersecret"),
			Expiry: getEnvInt("JWT_EXPIRY_MINUTES", 60),
		},
	}
}

// internal helpers
func getEnv(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value, exist := os.LookupEnv(key); exist {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return defaultValue
}
