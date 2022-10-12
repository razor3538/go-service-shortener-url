package config

import (
	"github.com/joho/godotenv"
	"os"
)

type env struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
	Host       string
}

// Env сущность хранящая переменные для подключения к бд
var Env env
var TemplateEnv env

func init() {
	_ = godotenv.Load(".env")
	Env = env{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		Port:       os.Getenv("PORT"),
		Host:       os.Getenv("HOST"),
	}
}
