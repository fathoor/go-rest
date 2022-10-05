package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload" // Auto-load .env file
)

type config struct {
	// Database
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASS     string
	DB_NAME     string
	DB_SSL_MODE string

	// Server
	SERVER_HOST string
	SERVER_PORT string
}

var Config = config{
	DB_HOST:     os.Getenv("DB_HOST"),
	DB_PORT:     os.Getenv("DB_PORT"),
	DB_USER:     os.Getenv("DB_USER"),
	DB_PASS:     os.Getenv("DB_PASSWORD"),
	DB_NAME:     os.Getenv("DB_NAME"),
	DB_SSL_MODE: os.Getenv("DB_SSL_MODE"),
	SERVER_HOST: os.Getenv("SERVER_HOST"),
	SERVER_PORT: os.Getenv("SERVER_PORT"),
}
