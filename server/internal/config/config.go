package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

type Config struct {
	Database DatabaseConfig
	//Server   any
	//Auth     any
}

func LoadConfig() Config {
	_ = godotenv.Load()

	return Config{
		Database: DatabaseConfig{
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
		},
		//Server: ServerConfig{
		//	Port: os.Getenv("SERVER_PORT"),
		//},
		//Auth: AuthConfig{
		//	JWTSecret: os.Getenv("JWT_SECRET"),
		//},
	}
}
